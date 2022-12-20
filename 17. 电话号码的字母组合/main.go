package main

import (
	"log"
)

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	key := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	if len(digits) == 1 {
		return key[digits[0]]
	}

	v, _ := key[digits[0]]
	ret := make([]string, 0)
	lastLen := 0
	for i := 1; i < len(digits); i++ {
		v2, _ := key[digits[i]]

		for j := 0; j < len(v); j++ {
			for k := 0; k < len(v2); k++ {
				ret = append(ret, v[j]+v2[k])
			}
		}
		ret = ret[lastLen:]
		lastLen = len(ret)
		v = ret
	}
	return ret
}

func main() {
	log.Println(letterCombinations("2"))
}
