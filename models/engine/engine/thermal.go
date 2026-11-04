package engine

import "a14industry.baa/smr-engine/models"

// CalculateHeatBalance evaluates ambient thermal impact on overall cycle efficiency.
func CalculateHeatBalance(req models.SiteRequirements, modifications *[]string) float64 {
	standard := models.GetStandardDesign()
	adjustedOutput := standard.StandardNetOutput

	// If site ambient temperature is higher than design parameters, cooling efficiency decreases
	if req.AmbientTemp > 30.0 {
		efficiencyLoss := (req.AmbientTemp - 30.0) * 0.55 // 0.55 MWe loss per degree above 30Â°C
		adjustedOutput -= efficiencyLoss
		*modifications = append(*modifications, "Elevated ambient intake temperatures detected. Condenser/cooling tower sizing optimized to prevent deep thermal de-rating.")
	}

	return adjustedOutput
}
