package models

import "time"

// Category ...DB -> Category Schema
type Category struct {
	CategoryID   *int
	CategoryName *string
	AddedTime    *time.Time
	LastUpdated  *time.Time
	RemovedTime  *time.Time
}
