package validators

import "strings"

// IsCredit checks if the amount is a credit transaction_service
func IsCredit(amount string) bool {
	return strings.Contains(amount, "+")
}
