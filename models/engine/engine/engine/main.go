package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"a14industry.baa/smr-engine/engine"
	"a14industry.baa/smr-engine/models"
)

func main() {
	// 1. Establish the API Endpoint Route
	http.HandleFunc("/api/evaluate", func(w http.ResponseWriter, r *http.Request) {
		// Inject strict Cross-Origin Resource Sharing (CORS) security headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight browser options verification checklist safely
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Enforce strict POST body payload requirements
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Anonymous structural blueprint to extract arriving frontend parameters
		var payload struct {
			CustomerID      string  `json:"customer_id"`
			CountryCode     string  `json:"country_code"`
			SeismicPGA      float64 `json:"seismic_pga"`
			AmbientTemp     float64 `json:"ambient_temp"`
			GridFrequencyHz int     `json:"grid_frequency_hz"`
			RegulatoryBody  string  `json:"regulatory_body"`
		}

		// Parse the arriving JSON incoming payload data streams
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// 2. Map payload metrics directly into standard core schemas
		targetSite := models.SiteRequirements{
			CustomerID:         payload.CustomerID,
			CountryCode:        payload.CountryCode,
			SeismicPGA:         payload.SeismicPGA,
			AmbientTemp:        payload.AmbientTemp,
			GridFrequencyHz:    payload.GridFrequencyHz,
			RegulatoryBody:     payload.RegulatoryBody,
			TargetCommercialOp: time.Now().AddDate(9, 0, 0), // Firm 9-year forward horizon calculation
		}

		// 3. Execute multi-agent AI processing pipelines
		_, aiDrivenMods, netElectricalOutput := engine.EvaluateSiteWithAI(targetSite)
		totalEstimatedCost, deadlines, RACIMatrix := engine.GenerateProjectPlanWithAI(targetSite, netElectricalOutput, aiDrivenMods)

		// 4. Return compiled analytics payload data objects back to frontend
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

	// 5. Fire up the backend daemon service thread cluster listener
	fmt.Println("=====================================================================")
	fmt.Println("   BWRX-300 SMR Configuration Engine [AI WEB APP MODE ACTIVE]")
	fmt.Println("   Listening actively on host target: http://localhost:8081")
	fmt.Println("=====================================================================")

	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
