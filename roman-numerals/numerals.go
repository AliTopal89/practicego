package numerals

import (
	"fmt"
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

var symboltonum = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(roman int) string {

	var result strings.Builder
	for _, numerals := range symboltonum {
		for roman >= numerals.Value {
			result.WriteString(numerals.Symbol)
			roman -= numerals.Value
		}
	}
	fmt.Println(result.String())
	return result.String()
}

func ConvertToNum(num string) int {
	return 1
}
