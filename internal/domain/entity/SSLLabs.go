package entity

type SSLLabsResponse struct {
	Host            string    `json:"host"`
	Port            int       `json:"port"`
	Protocol        string    `json:"protocol"`
	IsPublic        bool      `json:"isPublic"`
	Status          string    `json:"status"`
	StartTime       int64     `json:"startTime"`
	TestTime        int64     `json:"testTime"`
	EngineVersion   string    `json:"engineVersion"`
	CriteriaVersion string    `json:"criteriaVersion"`
	Endpoints       Endpoints `json:"endpoints"`
}

type Endpoints []Endpoint

func (e *Endpoints) GetServerNames() (serverNames []string) {
	for _, endpoint := range *e {
		serverNames = append(serverNames, endpoint.ServerName)
	}
	return serverNames
}

type Endpoint struct {
	IPAddress         string `json:"ipAddress"`
	ServerName        string `json:"serverName"`
	StatusMessage     string `json:"statusMessage"`
	Grade             string `json:"grade"`
	GradeTrustIgnored string `json:"gradeTrustIgnored"`
	HasWarnings       bool   `json:"hasWarnings"`
	IsExceptional     bool   `json:"isExceptional"`
	Progress          int    `json:"progress"`
	Duration          int    `json:"duration"`
	Delegation        int    `json:"delegation"`
}
