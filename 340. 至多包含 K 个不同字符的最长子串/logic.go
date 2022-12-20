package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstringKDistinct("cfaeabdac", 4))
}

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	start := 0
	end := 0

	m := make(map[byte]int)

	maxSize := 0

	maxL := 0

	for ; end < len(s); end++ {
		_, ok := m[s[end]]
		//插入k个不同字符
		if ok || k > 0 {
			m[s[end]]++
			if !ok && k > 0 {
				k--
			}
			continue
		}
		//不存在map中，并且k=0，即出现了第k+1个不同字符
		if !ok && k == 0 {
			//结束该k个字符的搜索
			if maxSize < end-start {
				maxSize = end - start
				maxL = start
			}
			//向前移动start，直到从map中删除一个字符
			for {
				m[s[start]]--
				if m[s[start]] == 0 {
					delete(m, s[start])
					k++
					start++
					break
				}
				start++
			}

			end--
		}
	}

	if k == 0 && end-start > maxSize {
		maxSize = end - start
		maxL = start
	}

	fmt.Println(s[maxL : maxL+maxSize])

	return maxSize
}
