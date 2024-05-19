package util

import "github.com/umardev500/banksampah/types"

func MappingKey(src string) string {
	switch src {
	case string(types.UserID):
		return types.HumanUserID
	}

	return ""
}
