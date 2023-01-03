package _042_Check_if_Numbers_Are_Ascending_in_a_Sentence

import (
	"strconv"
	"strings"
	"unicode"
)

func areNumbersAscending(s string) bool {
	prev := 0
	for _, token := range strings.Split(s, " ") {
		if unicode.IsDigit(rune(token[0])) {
			v, _ := strconv.Atoi(token)
			if v <= prev {
				return false
			}
			prev = v
		}
	}
	return true
}
