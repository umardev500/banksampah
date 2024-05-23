package constant

type ErrCodeName string

const (
	ErrCodeNameDuplicate    ErrCodeName = "DUPLICATE"
	ErrCodeNameValidation   ErrCodeName = "VALIDATION"
	ErrCodeNameInvalidID    ErrCodeName = "INVALID_ID"
	ErrCodeNameInvalidIDs   ErrCodeName = "INVALID_IDS"
	ErrCodeNameConstraint   ErrCodeName = "CONSTRAINT"
	ErrCodeNameOutOfBalance ErrCodeName = "OUT_OF_BALANCE"
	ErrCodeNameEmpy         ErrCodeName = "EMPTY"
)
