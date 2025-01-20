package model

// NewsCategory links News and Category in a many-to-many table
type NewsCategory struct {
	NewsID     int64 `reform:"NewsId"`
	CategoryID int64 `reform:"CategoryId"`
}
