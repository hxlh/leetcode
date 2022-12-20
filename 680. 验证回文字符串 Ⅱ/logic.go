package main

import "fmt"

func main() {
	fmt.Println(validPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"))
	// fmt.Println(validPalindrome("abca"))
}

//检测start~end之间字符串是否回文
func check(ss *string,start,end int) (int,int) {
	for start <= end {
		if (*ss)[start] != (*ss)[end] {
			return start,end
		}
		start++
		end--
	}
	start=end
	return start,end
}

func validPalindrome(s string) bool {
	start := 0
	end := len(s) - 1
	//第一次检测冲突位置
	start,end=check(&s,start,end)
	if start==end {
		return true
	}else {
		//删除start位置，并检查剩余字符串是否回文
		if tstart,tend:=check(&s,start+1,end);tstart==tend {
			return true
		}
		//删除end位置，并检查剩余字符串是否回文
		if tstart,tend:=check(&s,start,end-1);tstart==tend {
			return true
		}
	}

	return false
}
