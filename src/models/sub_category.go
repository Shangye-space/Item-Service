package models

import "time"

// SubCategory ...DB -> SubCategory Schema
type SubCategory struct {
	ID          *int
	Name        *string
	CategoryID  *int
	AddedTime   *time.Time
	LastUpdated *time.Time
	RemovedTime *time.Time
}
