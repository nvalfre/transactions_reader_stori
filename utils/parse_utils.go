package utils

import (
	"fmt"
	"strings"
	"transactions_reader_stori/validators"
)

// ParseAmount parses the amount string into a float64 value
func ParseAmount(amount string) float64 {
	value := strings.ReplaceAll(amount, "+", "")
	value = strings.ReplaceAll(value, "-", "")
	value = strings.TrimSpace(value)

	result := 0.0
	fmt.Scan(value, &result)

	if validators.IsCredit(amount) {
		return result
	}
	return -result
}
