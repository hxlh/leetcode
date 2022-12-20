package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLongestWord("abce", []string{"abe","abc"}))
}

type mystring []string

func (m mystring) Len() int {
	return len(m)
}
 
func (m mystring) Less(i, j int) bool {
	if len(m[i]) == len(m[j]) {
		for k := 0; k < len(m[i]); k++ {
			if m[i][k] < m[j][k] {
				return true
			}else if m[i][k] == m[j][k]{
				continue
			}else{
				return false
			}
		}
	}
	return len(m[i]) > len(m[j])
}

func (m mystring) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

/**
思路：先排序从最长的dictionary开始找
判断s是否按顺序包含dictionary的所有字符
如果答案不止一个，返回长度最长且字母序最小的字符串。
如果答案不存在，则返回空字符串。
*/
func findLongestWord(s string, dictionary []string) string {
	sort.Sort(mystring(dictionary))

	// fmt.Println(dictionary)

	for i := 0; i < len(dictionary); i++ {
		dst := dictionary[i]

		sindex := 0
		dstIndex := 0
		for sindex < len(s) && dstIndex < len(dst) {
			if s[sindex] == dst[dstIndex] {
				dstIndex++
				sindex++
			} else {
				sindex++
			}
		}

		//找到
		if dstIndex>=len(dst) {
			return dst
		}
	}

	return ""
}
