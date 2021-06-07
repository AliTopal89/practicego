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

type windowedRoman string

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

func ConverttoNum(num string) (total int) {
	for _, symbols := range windowedRoman(num).Symbols() {
		total += allSymboltonum.ValueOf(symbols...)
	}
	fmt.Println("what the total", total)
	return
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

func (r romanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return true
		}
	}
	return false
}

func (w windowedRoman) Symbols() (symbols [][]byte) {
	//[][] byte - slice of slices of bytes

	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := i+1 < len(w)

		if notAtEnd && couldBeSubtractive(symbol) && allSymboltonum.Exists(symbol, w[i]) {
			symbols = append(symbols, []byte{symbol, w[i+1]})
			i++
		} else {
			symbols = append(symbols, []byte{symbol})
		}
	}
	return
}

func couldBeSubtractive(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}
