package main

import "log"

func toLow(s byte)byte{
	dur:='a'-'A'
	if 'A'<=s && s<='Z' {
		return s+byte(dur)
	}
	return s
}

func isPalindrome(s string) bool {
	pre := 0
	last := len(s) - 1
	
	for pre != last && pre < last {
		isOk := true
		if !(('a' <= s[pre] && s[pre] <= 'z') || ('A' <= s[pre] && 'Z' >= s[pre]) || ('0' <= s[pre] && '9' >= s[pre])) {
			pre++
			isOk = false
		}
		if !(('a' <= s[last] && s[last] <= 'z') || ('A' <= s[last] && 'Z' >= s[last]) || ('0' <= s[last] && '9' >= s[last])) {
			last--
			isOk = false
		}

		if isOk {
			if !(toLow(s[pre])==toLow(s[last])) {
				return false
			}
			pre++
			last--
		}
		
		
	}
	return true
}

func main() {
	log.Println(isPalindrome("0P"))
}
