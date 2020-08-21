package p0752

//leetcode submit region begin(Prohibit modification and deletion)
type Node struct {
	Val  string
	Step int
}

func openLock(deadends []string, target string) int {
	visited := make(map[string]bool)
	for _, v := range deadends {
		visited[v] = true
	}

	var cur Node
	queue := []Node{{Val: "0000", Step: 0}}
	for len(queue) > 0 {
		cur, queue = queue[0], queue[1:]
		if !visited[cur.Val] {
			visited[cur.Val] = true
		} else {
			continue
		}

		if cur.Val == target {
			return cur.Step
		}

		// 1st up
		queue = append(queue, Node{Val: adjust(cur.Val, 3, true), Step: cur.Step + 1})
		// 1st down
		queue = append(queue, Node{Val: adjust(cur.Val, 3, false), Step: cur.Step + 1})
		// 2nd up
		queue = append(queue, Node{Val: adjust(cur.Val, 2, true), Step: cur.Step + 1})
		// 2nd down
		queue = append(queue, Node{Val: adjust(cur.Val, 2, false), Step: cur.Step + 1})
		// 3rd up
		queue = append(queue, Node{Val: adjust(cur.Val, 1, true), Step: cur.Step + 1})
		// 3rd down
		queue = append(queue, Node{Val: adjust(cur.Val, 1, false), Step: cur.Step + 1})
		// 4th up
		queue = append(queue, Node{Val: adjust(cur.Val, 0, true), Step: cur.Step + 1})
		// 4th down
		queue = append(queue, Node{Val: adjust(cur.Val, 0, false), Step: cur.Step + 1})
	}

	return -1
}

func adjust(v string, i int, up bool) string {
	bytes := []byte(v)
	d := bytes[i] - '0'
	if up {
		bytes[i] = (d+1)%10 + '0'
	} else {
		bytes[i] = (10+d-1)%10 + '0'
	}
	return string(bytes)
}

//leetcode submit region end(Prohibit modification and deletion)
