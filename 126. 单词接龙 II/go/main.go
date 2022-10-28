package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// 开启 CPU profiling
	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		pprof.StopCPUProfile()
		f.Close()
	}()
	for i := 0; i < 10000; i++ {
		findLadders("kite", "ashy", []string{"ante", "does", "jive", "achy", "kids", "kits", "cart", "ache", "ands", "ashe", "acne", "aunt", "aids", "kite", "ants", "anne", "ashy"})
	}
}


// 双向bfs的求解思路
// 既然是双向bfs，那么我们需要两个队列，分别代表两个方向的搜索。
// 然后，我们需要两个map，k表示已经遍历的节点，v表示当前遍历处于的层数。（双向bfs的map其实就是朴素bfs的变量+set）
// 遍历的核心逻辑和朴素bfs是一样的，只不过每次我们选择的队列应当是两个队列中长度最短的那个队列进行遍历。
// 遍历结束的两个标志：当遍历的过程中找到对方也搜索过的数据，说明已经找到最短路径，如果有一个队列为空退出循环，那么则没有找到目标数据。

// 判断是否word1修改一个字符即变为word2
func assist(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	diff := 0
	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			diff++
		}
	}
	if diff > 1 {
		return false
	}
	return true
}

func contain(s []string, ss string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == ss {
			return true
		}
	}
	return false
}

// 思路：广度优先遍历查找每一个string的只差一个字符的子[]string，从两端开始搜索直到该层存在相同的子string。
// 最后使用深度优先遍历重构路径
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	if assist(beginWord, endWord) {
		return [][]string{{beginWord, endWord}}
	}
	// 检查List中是否有endWord
	if !contain(wordList, endWord) {
		return [][]string{}
	}

	head_queue := make([]string, 0)
	tail_queue := make([]string, 0)

	head_visited := make(map[string]int, 0)
	tail_visited := make(map[string]int, 0)
	head_level := 0
	tail_level := 0

	nexts := make(map[string][]string, 0)
	// init
	head_queue = append(head_queue, beginWord)
	tail_queue = append(tail_queue, endWord)
	found := false
	// false-> head, true->tail
	reverse := false

	for len(head_queue) > 0 && len(tail_queue) > 0 {
		cnt := 0
		if !reverse {
			cnt = len(head_queue)
			head_level++
			Debug("head bfs %v\n", head_queue[:cnt])
		} else {
			cnt = len(tail_queue)
			tail_level++
			Debug("tail bfs %v\n", tail_queue[:cnt])
		}

		tmpNext := make(map[string][]string, 0)
		// 遍历这一层
		for cnt > 0 {
			cnt -= 1
			curWord := ""
			if !reverse {
				curWord = head_queue[0]
				head_queue = head_queue[1:]
			} else {
				curWord = tail_queue[0]
				tail_queue = tail_queue[1:]
			}

			for _, word := range wordList {

				if curWord == word {
					if !reverse {
						if curWord == endWord {
							if !found {
								if tail_level > 0 {
									head_level--
								}
							}
							found = true
							break
						}
					} else {
						if curWord == beginWord {
							if !found {
								if head_level > 0 {
									tail_level--
								}
							}
							found = true
							break
						}
					}
					continue
				}

				if !assist(curWord, word) {
					continue
				}
				if !reverse {
					if _, ok := head_visited[word]; ok {
						continue
					}
					if _, ok := tail_visited[curWord]; ok {
						if !found {
							head_level--
						}
						found = true
						break
					}
					head_queue = append(head_queue, word)
					tmpNext[curWord] = append(tmpNext[curWord], word)
					head_visited[curWord] = head_level
				} else {
					if _, ok := tail_visited[word]; ok {
						continue
					}
					if _, ok := head_visited[curWord]; ok {
						if !found {
							tail_level--
						}
						found = true
						break
					}
					if !contain(tail_queue, word) {
						tail_queue = append(tail_queue, word)
					}
					tmpNext[word] = append(tmpNext[word], curWord)
					tail_visited[curWord] = tail_level
				}
			}
		}

		if len(tmpNext) > 0 {
			for k, v := range tmpNext {
				for i := 0; i < len(v); i++ {
					if contain(nexts[k], v[i]) {
						continue
					}
					nexts[k] = append(nexts[k], v[i])
				}
			}
			Debug("rev %v\n arr %v\n", reverse, tmpNext)
		}

		if found {
			// len(ans)
			maxLevel := head_level + tail_level

			Debug("%v\n", maxLevel)
			ans := make([][]string, 0)
			path := make([]string, 0)
			ans, _ = dfs(maxLevel, 1, ans, path, nexts, beginWord, endWord)
			Debug("%v\n", ans)
			return ans
		}

		// 从小的开始bfs
		if len(head_queue) > len(tail_queue) {
			reverse = true
		} else {
			reverse = false
		}
	}

	return [][]string{}
}

func dfs(maxLevel int, level int, ans [][]string, path []string, next map[string][]string, src string, dst string) ([][]string, []string) {
	if level > maxLevel {
		return ans, path
	}
	if src == dst {
		tmp := make([]string, len(path)+1)
		copy(tmp[:len(path)], path)
		tmp[len(tmp)-1] = src

		Debug("%v\n", tmp)
		ans = append(ans, tmp)
		return ans, path
	}
	nexts := next[src]
	path = append(path, src)
	Debug("%v\n", path)
	for _, n := range nexts {
		ans, path = dfs(maxLevel, level+1, ans, path, next, n, dst)
	}
	path = path[:len(path)-1]
	Debug("%v\n", path)
	return ans, path
}

const debug = false

func Debug(s string, v ...interface{}) {
	if debug {
		fmt.Printf(s, v...)
	}
}