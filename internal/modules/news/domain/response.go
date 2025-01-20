package domain

// NewsListResponse for GET /list
type NewsListResponse struct {
	Success bool         `json:"Success"`
	News    []NewsDetail `json:"News"`
}

// NewsDetail provides extended info for each news item
type NewsDetail struct {
	ID         int64   `json:"Id"`
	Title      string  `json:"Title"`
	Content    string  `json:"Content"`
	Categories []int64 `json:"Categories"`
}
