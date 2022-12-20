package main

import (
	"fmt"
)

func main() {

	fmt.Println(minWindow("DOABECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	if len(t)>len(s) {
		return ""
	}
	
	tm := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tm[t[i]]++
	}

	l, r := 0, 0
	cnt := 0
	minSize := len(s) + 1
	finalL := 0
	for r < len(s) && l <= r {
		if v, ok := tm[s[r]]; ok {
			//r移动到包含t的位置，并且多出的相同字符会减为<0,例如-1则是该字符在窗口中多出现了一次
			if v > 0 {
				cnt++
			}
			tm[s[r]]--
			//l右移：排除无关的字符和重复多余的字符，但始终保证窗口中包含t的字符
			for cnt == len(t) {
				//如果l的字符在t中存在,v>0说明还要字符没出现,v==0说明字符出现了一次,v<0说明字符出现了多次（还不是最小）
				//cnt == len(t)已保证所有字符已出现,即v<=0
				if v, ok := tm[s[l]]; ok {
					if v == 0 {
						if r-l+1 < minSize {
							minSize = r - l + 1
							finalL = l
						}
						tm[s[l]]++
						cnt--
					} else {
						tm[s[l]]++
					}
				}
				l++
			}
		}
		r++
	}

	if minSize == len(s)+1 {
		return ""
	}

	return s[finalL : finalL+minSize]
}
