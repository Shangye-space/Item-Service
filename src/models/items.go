package models

import "time"

// Items ...DB -> Items Schema
type Items struct {
	ItemID      int
	ItemName    string
	Quantity    int
	Description string
	Price       float32
	Discount    int
	InSale      bool
	Category    string
	SubCategory *string
	AddedTime   time.Time
	RemovedTime *time.Time
}
