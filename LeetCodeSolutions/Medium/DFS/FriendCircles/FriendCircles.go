/*
https://leetcode.com/problems/friend-circles/
There are N students in a class. Some of them are friends, while some are not.
Their friendship is transitive in nature. For example, if A is a direct friend of B,
 and B is a direct friend of C, then A is an indirect friend of C.
 And we defined a friend circle is a group of students who are direct or indirect friends.

Given a N*N matrix M representing the friend relationship between students in the class.
If M[i][j] = 1, then the ith and jth students are direct friends with each other,
otherwise not. And you have to output the total number of friend circles among all the students.

Example 1:

Input:
[[1,1,0],
 [1,1,0],
 [0,0,1]]
Output: 2
Explanation:The 0th and 1st students are direct friends, so they are in a friend circle.
The 2nd student himself is in a friend circle. So return 2.


Example 2:

Input:
[[1,1,0],
 [1,1,1],
 [0,1,1]]
Output: 1
Explanation:The 0th and 1st students are direct friends,
the 1st and 2nd students are direct friends,
so the 0th and 2nd students are indirect friends.
All of them are in the same friend circle, so return 1.
*/
package main

import "fmt"

func findCircleNum(M [][]int) int {
	visited := make(map[int]bool)
	n := len(M)

	fmt.Println("Size of the matrix", n)
	for idx := 0; idx < n; idx++ {
		visited[idx] = false
	}

	circleNum := 0
	for idx := 0; idx < n; idx++ {
		if found, _ := visited[idx]; !found {
			circleNum++
			findCircleNumUtil(M, visited, idx, n)
			visited[idx] = true

		}
	}

	return circleNum
}

func findCircleNumUtil(M [][]int, visited map[int]bool, u int, size int) {
	for v := 0; v < size; v++ {
		if found, _ := visited[v]; !found && M[u][v] == 1 {
			visited[v] = true
			findCircleNumUtil(M, visited, v, size)
		}
	}
}

func main() {
	matrix := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	num := findCircleNum(matrix)
	fmt.Println(num)

}
