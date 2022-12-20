package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(frequencySort("tree"))
}

type intint [][]int

func (a intint) Len() int           { return len(a) }
func (a intint) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a intint) Less(i, j int) bool { return a[i][1] > a[j][1] }

func frequencySort(s string) string {
	c := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		c[s[i]]++
	}

	bulk := make([][]int, len(c))
	for i := 0; i < len(c); i++ {
		bulk[i] = make([]int, 2)
	}

	i := 0
	for k, v := range c {
		bulk[i][0] = int(k)
		bulk[i][1] = v
		i++
	}

	sort.Sort(intint(bulk))

	var buf bytes.Buffer

	for i := 0; i < len(bulk); i++ {
		for j := 0; j < bulk[i][1]; j++ {
			buf.WriteByte(byte(bulk[i][0]))
		}
	}

	return buf.String()
}
