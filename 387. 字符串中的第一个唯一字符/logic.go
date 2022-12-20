package main

import (
	"log"
)

func firstUniqChar(s string) int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		_, ok := m[s[i]]
		if !ok {
			m[s[i]] = i
		} else {
			m[s[i]] = -1
		}
	}
	for i := 0; i < len(s); i++ {
		v, ok := m[s[i]]
		if ok && v >= 0 {
			return i
		}
	}

	return -1
}

func main() {
	log.Println(firstUniqChar("leetcode"))
}
