package models

type PageLimit struct {
	Offset   int `json:"offset"  example:"1"`
	PageSize int `json:"pageSize" example:"10"`
}

type PageResult struct {
	PageSize   int64 `json:"pageSize"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"totalPages"`
	PageNumber int64 `json:"pageNumber"`
}
