package models

// Image ...DB -> Image Schema
type Image struct {
	ID           *int
	ItemID       *int
	PrimaryImage *int
	Path         *string
}
