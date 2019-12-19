package models

import "time"

// ItemInfo ...DB -> Item info
type ItemInfo struct {
	ItemID       *int
	Quantity     *int
	Description  *string
	Price        *float32
	Discount     *float32
	Size         *string
	Color        *string
	Manufacturer *string
	ItemCode     *string
	Material     *string
	AddedTime    *time.Time
	LastUpdated  *time.Time
	RemovedTime  *time.Time
}
