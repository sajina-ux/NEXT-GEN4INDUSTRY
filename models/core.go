package models

import "time"

type StandardBWRX300 struct {
    MaxSeismicPGA     float64
    MaxAmbientTemp    float64
    StandardNetOutput float64
    BaseCostUSD       float64
}

type SiteRequirements struct {
    CustomerID         string
    CountryCode        string
    SeismicPGA         float64
    AmbientTemp        float64
    GridFrequencyHz    int
    RegulatoryBody     string
    TargetCommercialOp time.Time
}

type AdaptationReport struct {
    CustomerID            string
    CountryCode           string
    IsCompatible          bool
    RequiredModifications []string
    AdjustedNetOutput     float64
    EstimatedCostUSD      float64
    MilestoneDeadlines    []MilestoneTarget
    ResponsibilityMatrix  map[string]string
}

type MilestoneTarget struct {
    PhaseName  string
    TargetDate time.Time
}

func GetStandardDesign() StandardBWRX300 {
    return StandardBWRX300{
        MaxSeismicPGA:     0.3,
        MaxAmbientTemp:    40.0,
        StandardNetOutput: 300.0,
        BaseCostUSD:       1000000000,
    }
}
