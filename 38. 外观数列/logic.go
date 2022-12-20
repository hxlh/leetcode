package main

import (
	"log"
	"strconv"
)

func countAndSay(n int) string {
	s := "11"
	if n < 2 {
		return s[:1]
	}
	pre := 0
	next := 0
	
	for i := 2; i < n; i++ {
		ans:=""
		for  next<len(s)-1{
			if s[next]!=s[next+1] {
				ans+=strconv.Itoa(next-pre+1)+string(s[pre])
				next++
				pre=next
				continue
			}
			next++
		}
		s=ans+strconv.Itoa(next-pre+1)+string(s[pre])
		pre=0
		next=0
		// log.Printf("index: %d  v: %s\n",i,s)
	}

	return s
}

func main() {
	log.Println(countAndSay(3))
}
