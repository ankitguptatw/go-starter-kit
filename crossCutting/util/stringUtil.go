package util

import (
	"fmt"
	"strings"
)

func Format(template string, args ...interface{}) string {
	return fmt.Sprintf(template, args...)
}

func IsNilEmptyOrWhiteSpace(val string) bool {
	if val == "" || len(val) == 0 || strings.TrimSpace(val) == "" {
		return true
	}
	return false
}
