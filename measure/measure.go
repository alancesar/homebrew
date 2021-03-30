package measure

import (
	"regexp"
	"strconv"
	"strings"
)

var regex = regexp.MustCompile(`(\d{1,5}\.?\d*)(\s?)(\D{1,7})`)

const (
	valueIndex  = 1
	symbolIndex = 3
)

type (
	Measure interface {
		IsZero() bool
	}

	BindFn func(symbol string, value float64)
)

func ExtractSymbolAndValue(input string) (string, float64, bool) {
	elements := regex.FindStringSubmatch(input)
	if elements == nil {
		return "", 0, false
	}

	symbol := strings.ToLower(elements[symbolIndex])
	value, err := strconv.ParseFloat(elements[valueIndex], 64)
	return symbol, value, err == nil
}

func NewFrom(input string, fn BindFn) {
	if symbol, value, match := ExtractSymbolAndValue(input); match {
		fn(symbol, value)
	}
}

func HasSomeZeroValue(measures ...Measure) bool {
	for _, measure := range measures {
		if measure.IsZero() {
			return true
		}
	}

	return false
}
