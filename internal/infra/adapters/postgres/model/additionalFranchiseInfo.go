package model

import (
	"time"

	"github.com/lib/pq"
)

type AdditionalFranchiseInfo struct {
	FranchiseId           int64
	UrlImage              string
	Protocol              string
	TraceRoutes           pq.StringArray `gorm:"type:text[]"`
	DomainCreatedAt       time.Time
	DomainExpiredAt       time.Time
	DomainRegistrantName  string
	DomainRegistrantEmail string
}
