package types

type Pagination struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	Offset    int `json:"offset"`
	Total     int `json:"total"`
	PageTotal int `json:"page_total"`
}

type Filter struct {
	Field           string `json:"field"`
	Operator        string `json:"operator"`
	Value           string `json:"value"`
	LogicalOperator string `json:"logicalOperator"`
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
