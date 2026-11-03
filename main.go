package main

import (
    "encoding/json"
    "net/http"
    "time"
    "a14industry.baa/smr-engine/engine"
    "a14industry.baa/smr-engine/models"
)

const htmlDashboard = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>A14INDUSTRY.BAA SMR Portal</title>
    <style>
        body { font-family: 'Segoe UI', Arial, sans-serif; background: #1e1e24; color: #f5f5f5; margin: 0; padding: 20px; }
        .container { max-width: 900px; margin: 0 auto; background: #2a2a35; padding: 30px; border-radius: 8px; box-shadow: 0 4px 15px rgba(0,0,0,0.5); }
        h1 { color: #00d2ff; border-bottom: 2px solid #3a3a4a; padding-bottom: 10px; margin-top: 0; }
        .form-grid { display: grid; grid-template-columns: 1fr 1fr; gap: 20px; margin-bottom: 20px; }
        label { display: block; margin-bottom: 5px; color: #aaa; font-weight: bold; }
        input, select { width: 100%; padding: 10px; background: #1e1e24; border: 1px solid #4a4a5a; border-radius: 4px; color: #fff; box-sizing: border-box; }
        button { background: #00d2ff; color: #1e1e24; border: none; padding: 12px 20px; font-size: 16px; font-weight: bold; border-radius: 4px; cursor: pointer; width: 100%; transition: background 0.2s; }
        button:hover { background: #00b4db; }
        .results { margin-top: 30px; background: #1e1e24; padding: 20px; border-radius: 6px; border-left: 5px solid #00d2ff; display: none; }
        .matrix-item { margin-bottom: 8px; padding-left: 10px; border-left: 2px solid #555; }
    </style>
</head>
<body>
    <div class="container">
        <h1>BWRX-300 SMR Adaptation App</h1>
        <form id="evaluationForm">
            <div class="form-grid">
                <div>
                    <label>Customer/JV ID</label>
                    <input type="text" id="customer_id" value="EuroPower-JV" required>
                </div>
                <div>
                    <label>Country Code</label>
                    <input type="text" id="country_code" value="UK" required>
                </div>
                <div>
                    <label>Seismic PGA (G threshold)</label>
                    <input type="number" step="0.01" id="seismic_pga" value="0.35" required>
                </div>
                <div>
                    <label>Ambient Peak Temp (°C)</label>
                    <input type="number" step="0.1" id="ambient_temp" value="33.5" required>
                </div>
                <div>
                    <label>Grid Frequency</label>
                    <select id="grid_frequency_hz">
                        <option value="50">50 Hz</option>
                        <option value="60">60 Hz</option>
                    </select>
                </div>
                <div>
                    <label>Regulatory Framework Body</label>
                    <input type="text" id="regulatory_body" value="Office for Nuclear Regulation (ONR)" required>
                </div>
            </div>
            <button type="submit">Run Engineering AI Analysis</button>
        </form>

        <div class="results" id="resultsBlock">
            <h2>AI Assessment Output</h2>
            <p><strong>Net Plant Capacity:</strong> <span id="outNet"></span> MWe</p>
            <p><strong>Overnight Capital Cost Estimate:</strong> $<span id="outCost"></span> Billion USD</p>
            
            <h3>Required Engineering Adaptations</h3>
            <ul id="outAdaptations"></ul>

            <h3>Calculated Project Milestone Deadlines</h3>
            <ul id="outMilestones"></ul>

            <h3>Project Responsibility Matrix (RACI)</h3>
            <div id="outMatrix"></div>
        </div>
    </div>

    <script>
        document.getElementById('evaluationForm').addEventListener('submit', async (e) => {
            e.preventDefault();
            const payload = {
                customer_id: document.getElementById('customer_id').value,
                country_code: document.getElementById('country_code').value,
                seismic_pga: parseFloat(document.getElementById('seismic_pga').value),
                ambient_temp: parseFloat(document.getElementById('ambient_temp').value),
                grid_frequency_hz: parseInt(document.getElementById('grid_frequency_hz').value),
                regulatory_body: document.getElementById('regulatory_body').value
            };

            const response = await fetch('/api/evaluate', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(payload)
            });
            const data = await response.json();

            document.getElementById('outNet').innerText = data.optimized_net_output_mwe.toFixed(2);
            document.getElementById('outCost').innerText = data.overnight_cost_usd_billion.toFixed(2);
            
            const adaptList = document.getElementById('outAdaptations');
            adaptList.innerHTML = '';
            data.engineering_adaptations.forEach(m => {
                adaptList.innerHTML += '<li>' + m + '</li>';
            });

            const milestoneList = document.getElementById('outMilestones');
            milestoneList.innerHTML = '';
            data.milestone_deadlines.forEach(m => {
                const dateStr = new Date(m.TargetDate).toLocaleDateString('en-US', { month: 'long', day: 'numeric', year: 'numeric' });
                milestoneList.innerHTML += '<li><strong>' + m.PhaseName + '</strong>: ' + dateStr + '</li>';
            });

            const matrixDiv = document.getElementById('outMatrix');
            matrixDiv.innerHTML = '';
            for (const [scope, owner] of Object.entries(data.responsibility_matrix)) {
                matrixDiv.innerHTML += '<div class="matrix-item"><strong>' + scope + '</strong> &rarr; Lead Owner: ' + owner + '</div>';
            }

            document.getElementById('resultsBlock').style.display = 'block';
        });
    </script>
</body>
</html>`

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
            http.NotFound(w, r)
            return
        }
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        w.Write([]byte(htmlDashboard))
    })

    http.HandleFunc("/api/evaluate", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        var payload struct {
            CustomerID      string  `json:"customer_id"`
            CountryCode     string  `json:"country_code"`
            SeismicPGA      float64 `json:"seismic_pga"`
            AmbientTemp     float64 `json:"ambient_temp"`
            GridFrequencyHz int     `json:"grid_frequency_hz"`
            RegulatoryBody  string  `json:"regulatory_body"`
        }

        if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        targetSite := models.SiteRequirements{
            CustomerID:         payload.CustomerID,
            CountryCode:        payload.CountryCode,
            SeismicPGA:         payload.SeismicPGA,
            AmbientTemp:        payload.AmbientTemp,
            GridFrequencyHz:    payload.GridFrequencyHz,
            RegulatoryBody:     payload.RegulatoryBody,
            TargetCommercialOp: time.Now().AddDate(9, 0, 0),
        }

        _, aiDrivenMods, netElectricalOutput := engine.EvaluateSiteWithAI(targetSite)
        totalEstimatedCost, deadlines, RACIMatrix := engine.GenerateProjectPlanWithAI(targetSite, netElectricalOutput, aiDrivenMods)

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "status":                     "success",
            "customer_id":                targetSite.CustomerID,
            "market_region":              targetSite.CountryCode,
            "optimized_net_output_mwe":   netElectricalOutput,
            "overnight_cost_usd_billion": totalEstimatedCost / 1e9,
            "engineering_adaptations":    aiDrivenMods,
            "milestone_deadlines":        deadlines,
            "responsibility_matrix":      RACIMatrix,
        })
    })

    println("=====================================================================")
    println(" A14INDUSTRY.BAA SMR Web Interface [INLINE TEMPLATE ENGINE MODE]")
    println(" Open your web browser and navigate directly to: http://localhost:8081")
    println("=====================================================================")
    
    if err := http.ListenAndServe(":8081", nil); err != nil {
        panic(err)
    }
}
