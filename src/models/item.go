package models

import "time"

// Item ...DB -> Item Schema
type Item struct {
	ID            *int
	Name          *string
	Price         *float32
	SubCategoryID *int
	InSale        *bool
	AddedTime     *time.Time
	LastUpdated   *time.Time
	RemovedTime   *time.Time
}
