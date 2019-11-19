package models

import "time"

// Category ...DB -> Category Schema
type Category struct {
	ID          *int
	Name        *string
	AddedTime   *time.Time
	LastUpdated *time.Time
	RemovedTime *time.Time
}
