package models

// SubCategory ...DB -> SubCategory Schema
type SubCategory struct {
	SubCategoryID   *int
	SubCategoryName *string
	CategoryID      *int
}
