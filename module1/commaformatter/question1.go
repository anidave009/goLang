package commaformatter

import (
	"bytes"
)

func AddCommasToNumber(num string) string {

	// 1. Find the decimal point
	decimalIndex := -1

	for i := 0; i < len(num); i++ {
		if num[i] == '.' {
			decimalIndex = i
			break
		}
	}

	// If there is no decimal point,
	// the integer part is the entire number.
	if decimalIndex == -1 {
		decimalIndex = len(num)
	}

	// 2. Separate integer and decimal parts
	integerPart := num[:decimalIndex]
	decimalPart := num[decimalIndex:]

	// 3. Check for optional sign
	sign := ""

	if len(integerPart) > 0 && (integerPart[0] == '-' || integerPart[0] == '+') {
		sign = integerPart[:1]
		integerPart = integerPart[1:]
	}

	// 4. Build the integer part from right to left
	var buffer bytes.Buffer

	count := 0

	for i := len(integerPart) - 1; i >= 0; i-- {

		buffer.WriteByte(integerPart[i])
		count++

		// Add comma after every 3 digits,
		// but don't add one after the last digit.
		if count == 3 && i != 0 {
			buffer.WriteByte(',')
			count = 0
		}
	}

	// 5. The buffer currently contains the integer part backwards.
	// Reverse it.
	result := buffer.Bytes()

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	// 6. Construct final result
	var final bytes.Buffer

	final.WriteString(sign)
	final.Write(result)
	final.WriteString(decimalPart)

	return final.String()
}
