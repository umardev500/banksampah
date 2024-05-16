package types

type Pagination struct {
	Page   int `json:"page"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type Filter struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type Order struct {
	Field string `json:"field"`
	Dir   string `json:"dir"` // ASC or DESC
}

type QueryParam struct {
	Pagination Pagination `json:"pagination"`
	Filter     []Filter   `json:"filter"`
	Order      Order      `json:"order"`
}
