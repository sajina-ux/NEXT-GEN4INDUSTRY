package engine

import (
	"fmt"
	"a14industry.baa/smr-engine/models"
)

// EvaluateSite checks localized hazards against the BWRX-300 standard design envelope.
func EvaluateSite(req models.SiteRequirements) (bool, []string) {
	standard := models.GetStandardDesign()
	var modifications []string
	compatible := true

	// External Hazard Screening: Seismic
	if req.SeismicPGA > standard.MaxSeismicPGA {
		compatible = false
		modifications = append(modifications, fmt.Sprintf("High Seismic Risk: Local PGA (%vG) exceeds standard envelope (%vG). Reinforced foundation design required.", req.SeismicPGA, standard.MaxSeismicPGA))
	}

	// Grid adaptation check
	if req.GridFrequencyHz != 50 && req.GridFrequencyHz != 60 {
		compatible = false
		modifications = append(modifications, "Invalid grid frequency configuration specified.")
	} else if req.GridFrequencyHz == 50 {
		modifications = append(modifications, "Standard design runs at 60Hz. Adjusting turbine generator electrical subsystem for 50Hz operations.")
	}

	// Regulatory Adaptation Framework
	if req.CountryCode != "US" && req.CountryCode != "CA" {
		modifications = append(modifications, fmt.Sprintf("Non-standard regulatory compliance required for %s (%s framework). Licensing schedule buffer added.", req.CountryCode, req.RegulatoryBody))
	}

	return compatible, modifications
}
