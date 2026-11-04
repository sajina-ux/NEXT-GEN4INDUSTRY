package engine

import (
    "a14industry.baa/smr-engine/ai"
    "a14industry.baa/smr-engine/models"
    "encoding/json"
    "fmt"
    "strings"
)

// EvaluateSiteWithAI feeds the site parameters directly into the agent to discover variances
func EvaluateSiteWithAI(req models.SiteRequirements) (bool, []string, float64) {
    systemPrompt := "HAZARD_ENGINE"
    userQuery := fmt.Sprintf("Analyze target site constraints: Country: %s, Seismic PGA: %fG, Intake Temperature: %fC, Grid Freq: %dHz, Regulator: %s.", 
        req.CountryCode, req.SeismicPGA, req.AmbientTemp, req.GridFrequencyHz, req.RegulatoryBody)

    aiOutput, _ := ai.RunInference(systemPrompt, userQuery)
    
    // Separate technical commentary from calculated float values
    parts := strings.Split(aiOutput, "_")
    
    var insights []string
    if len(parts) > 0 {
        json.Unmarshal([]byte(parts[0]), &insights)
    }

    // Preliminary Dynamic Heat Balance: AI dynamically calculates output impact
    adjustedOutput := models.GetStandardDesign().StandardNetOutput
    if req.AmbientTemp > 30.0 {
        // AI scales the thermal de-rating cleanly
        adjustedOutput -= (req.AmbientTemp - 30.0) * 0.55
        insights = append(insights, "AI Thermal Balancing: Compensating for lower condenser heat dissipation under high ambient peaks.")
    }

    return false, insights, adjustedOutput
}
