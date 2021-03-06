package numerals

import (
	"fmt"
	"testing"
	"testing/quick"
)

var (
	cases = []struct {
		Description string
		Num         uint16
		Roman       string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
		{"4 gets converted to IV", 4, "IV"},
		{"5 gets converted to V", 5, "V"},
		{"9 gets converted to IX", 9, "IX"},
		{"10 gets converted to X", 10, "X"},
		{"14 gets converted to XIV", 14, "XIV"},
		{"18 gets converted to XVIII", 18, "XVIII"},
		{"20 gets converted to XX", 20, "XX"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
		{"40 gets converted to XL", 40, "XL"},
		{"47 gets converted to XLVII", 47, "XLVII"},
		{"49 gets converted to XLIX", 49, "XLIX"},
		{"50 gets converted to L", 50, "L"},
		{"1984 gets converted to MCMLXXXIV", 1984, "MCMLXXXIV"},
		{"2021 gets converted to MMXXI", 2021, "MMXXI"},
	}
)

func TestConvertingToRoman(t *testing.T) {
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Num)
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertingtoNumerals(t *testing.T) {
	for _, test := range cases[:1] {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Num), func(t *testing.T) {
			got := ConverttoNum(test.Roman)
			if got != test.Num {
				t.Errorf("got %d want %d", got, test.Num)
			}
		})
	}
}
func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(num uint16) bool {
		if num > 3999 {
			return true
		}
		t.Log("testing", num)
		roman := ConvertToRoman(num)
		fromRoman := ConverttoNum(roman)
		return fromRoman == num
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
