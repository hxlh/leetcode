package main

type Node struct {
	x    int
	step int
}

func minimumOperations(nums []int, start int, goal int) int {
	var visited [1005]bool
	queue := make([]*Node, 0)
	queue = append(queue, &Node{start, 0})
	visited[start] = true

	if start == goal {
		return 0
	}

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		for i := 0; i < len(nums); i++ {
			x1 := cur.x + nums[i]
			x2 := cur.x - nums[i]
			x3 := cur.x ^ nums[i]
			if !(x1 < 0 || x1 > 1000 || visited[x1]) {
				visited[x1] = true
				queue = append(queue, &Node{x1, cur.step + 1})
			}
			if !(x2 < 0 || x2 > 1000 || visited[x2]) {
				visited[x2] = true
				queue = append(queue, &Node{x2, cur.step + 1})
			}
			if !(x3 < 0 || x3 > 1000 || visited[x3]) {
				visited[x3] = true
				queue = append(queue, &Node{x3, cur.step + 1})
			}
			if x1 == goal || x2 == goal || x3 == goal {
				return cur.step + 1
			}

		}
	}
	return -1
}

func main() {
	
}
