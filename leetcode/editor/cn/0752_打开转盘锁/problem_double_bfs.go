package p0752

func openLockDoubleBfs(deadends []string, target string) int {
	visited := make(map[string]bool)
	for _, v := range deadends {
		visited[v] = true
	}

	q1 := make(map[string]bool)
	q1["0000"] = true
	q2 := make(map[string]bool)
	q2[target] = true
	step := 0

	for len(q1) > 0 && len(q2) > 0 {
		if len(q1) > len(q2) {
			q1, q2 = q2, q1
		}

		tmp := make(map[string]bool)
		for k, _ := range q1 {
			if visited[k] {
				continue
			}
			visited[k] = true

			if q2[k] {
				return step
			}

			tmp[adjust(k, 3, true)] = true
			tmp[adjust(k, 3, false)] = true
			tmp[adjust(k, 2, true)] = true
			tmp[adjust(k, 2, false)] = true
			tmp[adjust(k, 1, true)] = true
			tmp[adjust(k, 1, false)] = true
			tmp[adjust(k, 0, true)] = true
			tmp[adjust(k, 0, false)] = true
		}

		q1 = q2
		q2 = tmp

		step++
	}

	return -1
}
