package models

import "time"

// SubCategory ...DB -> SubCategory Schema
type SubCategory struct {
	SubCategoryID   *int
	SubCategoryName *string
	CategoryID      *int
	AddedTime       *time.Time
	LastUpdated     *time.Time
	RemovedTime     *time.Time
}
