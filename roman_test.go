package tdd_kata

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	testcases := []struct {
		Number int
		Roman  string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{30, "XXX"},
		{40, "XL"},
		{49, "XLIX"},
		{50, "L"},
		{90, "XC"},
		{100, "C"},
		{400, "CD"},
		{500, "D"},
		{900, "CM"},
		{1000, "M"},
		{1984, "MCMLXXXIV"},
		{2022, "MMXXII"},
	}
	for _, test := range testcases {
		t.Run(fmt.Sprintf("%d gets converted to %s", test.Number, test.Roman), func(t *testing.T) {
			actual := ConvertToRoman(test.Number)
			assert.Equal(t, test.Roman, actual)
		})
	}
}
