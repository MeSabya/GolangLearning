func countSubstrings(s string) int {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	count := 0
	n := len(s)
	for idx := 0; idx < n; idx++ {
		dp[idx][idx] = true
		count++
	}

	for startidx := n - 1; startidx >= 0; startidx-- {
		for endidx := startidx + 1; endidx < n; endidx++ {
			if s[startidx] == s[endidx] {
				if endidx-startidx == 1 || dp[startidx+1][endidx-1] {
					dp[startidx][endidx] = true
					count++
				}
			}
		}
	}

	return count
}

//Complexity of the above algorithm is O(n2)