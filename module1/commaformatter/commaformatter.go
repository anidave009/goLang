package commaformatter

import (
	"bytes"
	"fmt"
	"unicode"
)

func AddCommas(num string) (string, error) {
	endIndex, err := findEndIndex(num)
	if err != nil {
		return "", err
	}

	sign := ""
	start := 0

	// Keep the sign separate from the number.
	if num[0] == '+' || num[0] == '-' {
		sign = string(num[0])
		start = 1
	}

	var buf bytes.Buffer
	count := 0

	// Traverse integer part from right to left.
	for i := endIndex - 1; i >= start; i-- {
		buf.WriteByte(num[i])
		count++

		if count == 3 && i > start {
			buf.WriteByte(',')
			count = 0
		}
	}

	// Reverse the integer portion.
	result := reverse(buf.String())

	// Add sign back.
	result = sign + result

	// Add fractional portion unchanged.
	if endIndex < len(num) {
		result += num[endIndex:]
	}

	return result, nil
}

func findEndIndex(num string) (int, error) {
	if len(num) == 0 {
		return 0, fmt.Errorf("empty input")
	}

	dotFound := false
	signFound := false
	endIndex := len(num)

	for i, ch := range num {
		if ch == '+' || ch == '-' {
			if signFound || i != 0 {
				return 0, fmt.Errorf("invalid sign")
			}
			signFound = true
			continue
		}

		if ch == '.' {
			if dotFound {
				return 0, fmt.Errorf("more than one decimal point")
			}
			dotFound = true
			endIndex = i
			continue
		}

		if !unicode.IsDigit(ch) {
			return 0, fmt.Errorf("invalid character: %c", ch)
		}
	}

	return endIndex, nil
}

func reverse(s string) string {
	r := []byte(s)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}
