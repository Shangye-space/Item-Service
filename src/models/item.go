package models

import "time"

// Item ...DB -> Item Schema
type Item struct {
	ItemID        *int
	ItemName      *string
	SubCategoryID *int
	InSale        *bool
	AddedTime     *time.Time
	LastUpdated   *time.Time
	RemovedTime   *time.Time
}
