package engine

import "a14industry.baa/smr-engine/models"

func CalculateHeatBalance(req models.SiteRequirements, modifications *[]string) float64 {
    standard := models.GetStandardDesign()
    adjustedOutput := standard.StandardNetOutput

    if req.AmbientTemp > 30.0 {
        efficiencyLoss := (req.AmbientTemp - 30.0) * 0.55
        adjustedOutput -= efficiencyLoss
        *modifications = append(*modifications, "Elevated ambient intake temperatures detected. Condenser/cooling tower sizing optimized to prevent deep thermal de-rating.")
    }

    return adjustedOutput
}
