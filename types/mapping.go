package types

type Keys string

var (
	UserID Keys = "user_id"
)

var HumanKeys string

var (
	HumanUserID = "User ID"
)

func MappingKey(src string) string {
	switch src {
	case string(UserID):
		return HumanUserID
	}

	return ""
}
