package constant

type SqlState string

const (
	SqlStateDuplicate SqlState = "23505"
	SqlConstraint     SqlState = "23503"
)

type SqlErrPattern string

const (
	SqlErrPatternDuplicate  SqlErrPattern = `Key \(([^)]+)\)=\(([^)]+)\) already exists\.`
	SqlErrConstraintPattern SqlErrPattern = `Key \(([^)]+)\)=\(([^)]+)\) is not present in table`
)
