package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "time"
    "a14industry.baa/smr-engine/engine"
    "a14industry.baa/smr-engine/models"
)

func main() {
    // --- RESTORED SPECIFIC SERVICE ROUTE LOGIC ---
    // Endpoint handler explicitly pointing to the missing ai4industry.baawebservice gateway
    http.HandleFunc("/api/ai4industry.baawebservice", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "status":      "active",
            "service":     "ai4industry.baawebservice",
            "version":     "1.0.0",
            "description": "Enterprise AI Governance & Industrial Edge Automation Core Module",
            "timestamp":   time.Now().Format(time.RFC3339),
        })
    })

    // --- RETAINED SMR CALCULATOR ENGINE PIPELINE ---
    http.HandleFunc("/api/evaluate", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

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

    // --- FALLBACK INTERFACE ROUTING ---
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "message": "A14INDUSTRY.BAA Central AI Engine Active Gateway",
        })
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8081"
    }

    fmt.Printf("=====================================================================\n")
    fmt.Printf("   A14INDUSTRY.BAA Engine [ALL CORE INDUSTRIAL WEB SERVICES RESTORED]\n")
    fmt.Printf("   Active Endpoint 1: GET/POST /api/ai4industry.baawebservice\n")
    fmt.Printf("   Active Endpoint 2: POST     /api/evaluate\n")
    fmt.Printf("   Listening actively on host target port: :%s\n", port)
    fmt.Printf("=====================================================================\n")

    if err := http.ListenAndServe(":" + port, nil); err != nil {
        panic(err)
    }
}
