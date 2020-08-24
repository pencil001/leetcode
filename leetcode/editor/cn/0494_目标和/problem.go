package p0494

//leetcode submit region begin(Prohibit modification and deletion)
func findTargetSumWays(nums []int, S int) int {
	cnt := 0
	dfs(nums, S, &cnt)
	return cnt
}

func dfs(nums []int, target int, cnt *int) {
	if len(nums) == 0 {
		if target == 0 {
			*cnt++
		}

		return
	}

	n := nums[0]
	dfs(nums[1:], target-n, cnt)
	dfs(nums[1:], target+n, cnt)
}

//leetcode submit region end(Prohibit modification and deletion)
