package numerals

import (
	"fmt"
	"strings"
)

type romanNumeral struct {
	Value  int
	Symbol string
}

type romanNumerals []romanNumeral

var allSymboltonum = romanNumerals{
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
	for _, numerals := range allSymboltonum {
		for roman >= numerals.Value {
			result.WriteString(numerals.Symbol)
			roman -= numerals.Value
		}
	}
	fmt.Println(result.String())
	return result.String()
}

func (r romanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

func ConvertToNum(num string) int {
	total := 0

	for i := 0; i < len(num); i++ {
		symbol := num[i]

		if couldBeSubtractive(i, symbol, num) {
			if value := allSymboltonum.ValueOf(symbol, num[i+1]); value != 0 {
				total += value
				i++ // move past this character too for the next loop
			} else {
				total += allSymboltonum.ValueOf(symbol)
			}
		} else {
			total += allSymboltonum.ValueOf(symbol)
		}
	}
	return total
}

func couldBeSubtractive(index int, currentSymbol uint8, num string) bool {
	isSubtractiveSymbol := currentSymbol == 'I' || currentSymbol == 'X' || currentSymbol == 'C'
	return index+1 < len(num) && isSubtractiveSymbol
}
