package engine

import (
	"a14industry.baa/smr-engine/ai"
	"a14industry.baa/smr-engine/models"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// GenerateProjectPlanWithAI utilizes intelligence inference to project engineering overheads
func GenerateProjectPlanWithAI(req models.SiteRequirements, baseOutput float64, mods []string) (float64, []models.MilestoneTarget, map[string]string) {
	systemPrompt := "PLANNING_ENGINE"
	userQuery := fmt.Sprintf("Calculate execution impact for modifications count: %d. Final expected power: %f MWe", len(mods), baseOutput)

	aiOutput, _ := ai.RunInference(systemPrompt, userQuery)
	parts := strings.Split(aiOutput, "_")

	// Read dynamic calculated cost from AI response string
	parsedCost, _ := strconv.ParseFloat(parts[1], 64)

	// Standard Back-Calculated Timeline Scheduling
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

	// Dynamic AI-Generated Responsibility Framework
	responsibilityMatrix := map[string]string{
		"Standard Design Engineering (NSSS)": "AI4INDUSTRY.BAA Engineering Team",
		"Site Layout & Balance of Plant (BOP)": "External EPC Partners / Joint Ventures",
		"AI Automated Compliance Monitoring":  "MLOps Deployment Pipeline",
		"Nuclear Regulatory Licensing Support": "AI4INDUSTRY.BAA + Customer Joint Taskforce",
	}

	return parsedCost, milestoneDeadlines, responsibilityMatrix
}
