package utils

import (
	"strings"

	"github.com/BeepLoop/go-audiobook/types"
)

func IsInputEmpty(input *types.Admin) bool {
	if strings.TrimSpace(input.Username) == "" || strings.TrimSpace(input.Password) == "" {
		return true
	}

	return false
}
