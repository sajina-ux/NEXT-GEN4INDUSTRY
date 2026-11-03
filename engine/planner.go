package engine

import (
    "a14industry.baa/smr-engine/ai"
    "a14industry.baa/smr-engine/models"
    "fmt"
    "strconv"
    "strings"
)

func GenerateProjectPlanWithAI(req models.SiteRequirements, baseOutput float64, mods []string) (float64, []models.MilestoneTarget, map[string]string) {
    systemPrompt := "PLANNING_ENGINE"
    userQuery := fmt.Sprintf("Calculate execution impact for modifications count: %d. Final expected power: %f MWe", len(mods), baseOutput)

    aiOutput, _ := ai.RunInference(systemPrompt, userQuery)
    parts := strings.Split(aiOutput, "_")

    parsedCost, _ := strconv.ParseFloat(parts[len(parts)-1], 64)

    // Fetch dynamic branding context
    appConfig := models.LoadConfig()

    commercialOpTarget := req.TargetCommercialOp
    constructionStartTarget := commercialOpTarget.AddDate(-4, 0, 0)       
    manufacturingStartTarget := constructionStartTarget.AddDate(-3, 0, 0)  
    designStartTarget        := manufacturingStartTarget.AddDate(0, -18, 0) 
    licensingStartTarget     := designStartTarget.AddDate(-2, 0, 0)        

    milestoneDeadlines := []models.MilestoneTarget{
        {PhaseName: "Phase 1: Initiate Site Hazard Assessment & Licensing", TargetDate: licensingStartTarget},
        {PhaseName: "Phase 2: Commence Engineering Detail Design Adaptation", TargetDate: designStartTarget},
        {PhaseName: "Phase 3: Launch Nuclear Island Supply Chain & Manufacturing", TargetDate: manufacturingStartTarget},
        {PhaseName: "Phase 4: Pour First Structural Concrete (Construction Start)", TargetDate: constructionStartTarget},
        {PhaseName: "Phase 5: Commercial Operation Date (Grid Synchronization)", TargetDate: commercialOpTarget},
    }

    // Company identity is injected at execution runtime dynamically
    responsibilityMatrix := map[string]string{
        "Standard Design Engineering (NSSS)":    fmt.Sprintf("%s Engineering Team", appConfig.CompanyName),
        "Site Layout & Balance of Plant (BOP)":  "External EPC Partners / Joint Ventures",
        "AI Automated Compliance Monitoring":   "MLOps Deployment Pipeline",
        "Nuclear Regulatory Licensing Support":  fmt.Sprintf("%s + Customer Joint Taskforce", appConfig.CompanyName),
        "Reactor Fuel Supply and Lifecycle":    fmt.Sprintf("%s Fuel Services", appConfig.CompanyName),
    }

    return parsedCost, milestoneDeadlines, responsibilityMatrix
}
