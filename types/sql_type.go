package types

type SqlErrDetail struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
	Error string      `json:"error"`
}
