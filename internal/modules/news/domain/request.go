package domain

// UpdateNewsRequest for POST /edit/:id
// Using pointer fields to see if a field was provided or not.
type UpdateNewsRequest struct {
	ID         *int64   `json:"Id,omitempty"`
	Title      *string  `json:"Title,omitempty"`
	Content    *string  `json:"Content,omitempty"`
	Categories *[]int64 `json:"Categories,omitempty"`
}
