package measure

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var regex = regexp.MustCompile(`(\d{1,5}\.?\d*)(\s?)(\D{1,3})`)

const (
	valueIndex  = 1
	symbolIndex = 3
)

func ExtractSymbolAndValue(input string) (string, float64, error) {
	elements := regex.FindStringSubmatch(input)
	if elements == nil {
		return "", 0, fmt.Errorf("invalid wantValue: %s", input)
	}

	symbol := strings.ToLower(elements[symbolIndex])
	value, err := strconv.ParseFloat(elements[valueIndex], 64)
	return symbol, value, err
}
