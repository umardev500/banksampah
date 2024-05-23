package constant

type SqlState string

const (
	SqlStateDuplicate SqlState = "23505"
	SqlConstraint     SqlState = "23503"
)

type SqlErrPattern string

const (
	SqlKeyValuePattern SqlErrPattern = `Key \(([^)]+)\)=\(([^)]+)\)`
)
