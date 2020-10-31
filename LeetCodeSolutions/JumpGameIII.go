/*
Given an array of non-negative integers arr, you are initially positioned at
start index of the array. When you are at index i,
 you can jump to i + arr[i] or i - arr[i], check if you can reach to any index with value 0.

Notice that you can not jump outside of the array at any time.

*/
package leetcodesolutions

func CanReach(arr []int, start int) bool {
	//From any neighbor node can we reach our target
	//This is a classic DFS case.
	set := make(map[int]bool, 0)
	for _, val := range arr {
		set[val] = false
	}
	return dfs(arr, start, set)
}

func dfs(arr []int, curridx int, visited map[int]bool) bool {
	if len(arr) <= 0 || curridx >= len(arr) || curridx < 0 {
		return false
	}

	found := visited[curridx]
	if found {
		return false
	}

	if arr[curridx] == 0 {
		return true
	}

	visited[curridx] = true

	if dfs(arr, curridx+arr[curridx], visited) {
		return true
	}
	return dfs(arr, curridx-arr[curridx], visited)

}
