package currency

import (
	"strconv"
	"strings"
)

func ToRupiahFormat(amount float64) string {
	// Convert float64 to int64 (for currency formatting purposes)
	amountInt := int64(amount * 100) // Multiply by 100 to handle decimal places

	// Format as currency with commas
	amountStr := strconv.FormatInt(amountInt, 10)
	n := len(amountStr)
	if n <= 3 {
		return "Rp" + amountStr
	}
	result := strings.Builder{}
	// Insert commas every 3 digits from the end
	for i, c := range amountStr {
		if i > 0 && (n-i)%3 == 0 {
			result.WriteRune(',')
		}
		result.WriteRune(c)
	}

	return "Rp" + result.String()
}
