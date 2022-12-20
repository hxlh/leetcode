package main

import "log"

func isAnagram(s string, t string) bool {
	if len(s)!=len(t) {
		return false
	}
	m := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		_, ok := m[s[i]]
		if ok {
			m[s[i]] += 1
		} else {
			m[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		_, ok := m[t[i]]
		if ok {
			m[t[i]] -= 1
		} else {
			return false
		}
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	log.Println(isAnagram("rat","car"))
}