package checkers

import (
	"strings"
)

func IsEmpty(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}
