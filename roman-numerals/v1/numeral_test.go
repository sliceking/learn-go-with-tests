package v1

import (
	"strings"
	"testing"
	"testing/quick"
)

func TestArabicNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Want        uint16
		Roman       string
	}{
		{"1 is I", 1, "I"},
		{"2 is II", 2, "II"},
		{"3 gets converted to III", 3, "III"},
		{"4 gets converted to IV (can't repeat more than 3 times)", 4, "IV"},
		{"5 gets converted to V", 5, "V"},
		{"6 gets converted to VI", 6, "VI"},
		{"7 gets converted to VII", 7, "VII"},
		{"8 gets converted to VIII", 8, "VIII"},
		{"9 gets converted to IX", 9, "IX"},
		{"10 gets converted to X", 10, "X"},
		{"10 gets converted to X", 10, "X"},
		{"14 gets converted to XIV", 14, "XIV"},
		{"18 gets converted to XVIII", 18, "XVIII"},
		{"20 gets converted to XX", 20, "XX"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
		{"40 gets converted to XL", 40, "XL"},
		{"47 gets converted to XLVII", 47, "XLVII"},
		{"49 gets converted to XLIX", 49, "XLIX"},
		{"50 gets converted to L", 50, "L"},
		{"400 gets converted to CD", 400, "CD"},
		{"500 gets converted to D", 500, "D"},
		{"900 gets converted to CM", 900, "CM"},
		{"1000 gets converted to M", 1000, "M"},
		{"1984 gets converted to MCMLXXXIV", 1984, "MCMLXXXIV"},
		{"3999 gets converted to MMMCMXCIX", 3999, "MMMCMXCIX"},
		{"2014 gets converted to MMXIV", 2014, "MMXIV"},
		{"1006 gets converted to MVI", 1006, "MVI"},
		{"798 gets converted to DCCXCVIII", 798, "DCCXCVIII"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Want {
				t.Errorf("got %d, want %d", got, test.Want)
			}
		})
	}

}

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      uint16
		Want        string
	}{
		{"1 is I", 1, "I"},
		{"2 is II", 2, "II"},
		{"3 gets converted to III", 3, "III"},
		{"4 gets converted to IV (can't repeat more than 3 times)", 4, "IV"},
		{"5 gets converted to V", 5, "V"},
		{"6 gets converted to VI", 6, "VI"},
		{"7 gets converted to VII", 7, "VII"},
		{"8 gets converted to VIII", 8, "VIII"},
		{"9 gets converted to IX", 9, "IX"},
		{"10 gets converted to X", 10, "X"},
		{"10 gets converted to X", 10, "X"},
		{"14 gets converted to XIV", 14, "XIV"},
		{"18 gets converted to XVIII", 18, "XVIII"},
		{"20 gets converted to XX", 20, "XX"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
		{"40 gets converted to XL", 40, "XL"},
		{"47 gets converted to XLVII", 47, "XLVII"},
		{"49 gets converted to XLIX", 49, "XLIX"},
		{"50 gets converted to L", 50, "L"},
		{"400 gets converted to CD", 400, "CD"},
		{"500 gets converted to D", 500, "D"},
		{"900 gets converted to CM", 900, "CM"},
		{"1000 gets converted to M", 1000, "M"},
		{"1984 gets converted to MCMLXXXIV", 1984, "MCMLXXXIV"},
		{"3999 gets converted to MMMCMXCIX", 3999, "MMMCMXCIX"},
		{"2014 gets converted to MMXIV", 2014, "MMXIV"},
		{"1006 gets converted to MVI", 1006, "MVI"},
		{"798 gets converted to DCCXCVIII", 798, "DCCXCVIII"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %s, want %s", got, test.Want)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

type RomanNumerals []RomanNumeral

var allRomanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

func (r RomanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return true
		}
	}

	return false
}

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) (total uint16) {
	for _, symbols := range windowedRoman(roman).Symbols() {
		total += allRomanNumerals.ValueOf(symbols...)
	}
	return
}

type windowedRoman string

func (w windowedRoman) Symbols() (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := i+1 < len(w)

		if notAtEnd && isSubtractive(symbol) && allRomanNumerals.Exists(symbol, w[i+1]) {
			symbols = append(symbols, []byte{byte(symbol), byte(w[i+1])})
			i++
		} else {
			symbols = append(symbols, []byte{byte(symbol)})
		}
	}
	return
}

func isSubtractive(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}
