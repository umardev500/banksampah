package constant

type SqlState string

const (
	SqlStateDuplicate SqlState = "23505"
)

type SqlErrPattern string

const (
	SqlErrPatternDuplicate SqlErrPattern = `Key \(([^)]+)\)=\(([^)]+)\) already exists\.`
)
