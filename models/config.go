package models

import "os"

// Config holds dynamic enterprise variable parameters loaded at initialization
type Config struct {
    CompanyName string
}

// LoadConfig reads variables directly from the host infrastructure profile
func LoadConfig() Config {
    companyName := os.Getenv("SMR_COMPANY_NAME")
    if companyName == "" {
        // Secure default fallback identity
        companyName = "A14INDUSTRY.BAA"
    }

    return Config{
        CompanyName: companyName,
    }
}
