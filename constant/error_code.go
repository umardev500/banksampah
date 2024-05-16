package constant

type ErrCodeName string

const (
	ErrCodeNameDuplicate  ErrCodeName = "DUPLICATE"
	ErrCodeNameValidation ErrCodeName = "VALIDATION"
	ErrCodeNameInvalidID  ErrCodeName = "INVALID_ID"
)
