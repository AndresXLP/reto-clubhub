package entity

import "time"

type WebInfo struct {
	FranchiseID int64
	UrlImage    string
	Protocol    string
	TraceRoutes []string
	Domain      Domain
}

type Domain struct {
	CreatedAt       time.Time
	ExpiredAt       time.Time
	RegistrantName  string
	RegistrantEmail string
}
