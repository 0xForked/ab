package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(gridLake(5, 4, [][]string{
		{"#", ".", "#", "#", "#"},
		{".", ".", "#", "#", "#"},
		{"#", "#", ".", "#", "#"},
		{"#", "#", "#", "#", "."},
	}))

	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(1221))
	fmt.Println(isPalindrome(2579752))

	fmt.Println(timeConversion("12:01:00PM"))
	fmt.Println(timeConversion("12:01:00AM"))
}

// An `m` x `n` map grid is provided to you. Cells are marked
// either `#` or `.` where `#` represents `land` and `.` represents `water`.
// You can assume that the following conditions are all true:
// grid cells are connected horizontally / vertically / diagonally
// the map grid is completely surrounded by land
// Function Description
// Create a function that returns the number of lakes represented by a set of `water`
// Input Format
// The first line contains a single integer, `m`, the width of the map.
// The second line contains a single integer, `n`, the height of map.
// The following lines represent the map grid
// Constraints
// 1 <= `m` <= 100
// 1 <= `n` <= 100
func gridLake(m int, n int, lakes [][]string) int {
	v := make([][]bool, n)
	for i := range v {
		v[i] = make([]bool, m)
	}

	var c int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if lakes[i][j] == "." && !v[i][j] {
				c++
				v[i][j] = true
				q := []int{i, j}
				for len(q) > 0 {
					cr := q[0]
					cc := q[1]
					q = q[2:]
					for _, d := range []int{-1, 0, 1} {
						nr := cr + d
						if nr < 0 || nr >= n {
							continue
						}

						for _, d := range []int{-1, 0, 1} {
							nc := cc + d
							if nc < 0 || nc >= m {
								continue
							}

							if lakes[nr][nc] == "." && !v[nr][nc] {
								v[nr][nc] = true
								q = append(q, nr, nc)
							}
						}
					}
				}
			}
		}
	}
	return c
}

//func dfs(g [][]string, i, j, m, n int, v [][]bool) {
//	if i < 0 || i >= m || j < 0 || j >= n || v[i][j] || g[i][j] != "." {
//		return
//	}
//
//	v[i][j] = true
//
//	dfs(g, i-1, j, m, n, v)
//	dfs(g, i+1, j, m, n, v)
//	dfs(g, i, j-1, m, n, v)
//	dfs(g, i, j+1, m, n, v)
//	dfs(g, i-1, j-1, m, n, v)
//	dfs(g, i-1, j+1, m, n, v)
//	dfs(g, i+1, j-1, m, n, v)
//	dfs(g, i+1, j+1, m, n, v)
//}

// Do Not convert value to date
// Given a time in 12-hour AM/PM format, convert it to military (24-hour) time.
// Note:
// - 12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
// - 12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock.
// Example 1
// s = '12:01:00PM'
// Return '12:01:00'
// Example 2
// s = '12:01:00AM'
// Return '00:01:00'
// Function Description
// Complete the timeConversion function in the editor below.
// It should return a new string representing the input time in 24 hour format.
// timeConversion has the following parameter(s):
// • string s: a time in 12 hour format
// Returns
// • string: the time in 24 hour format
// Input Format
// A single string s that represents a time in 12-hour clock format (i.e.: hh:mm:ssAM or hh:mm:ssPM).
// Constraints
// All input times are valid
// Sample Input
// Sample Output
//
//07:05:45PM
//19:05:45
func timeConversion(s string) string {
	time := strings.Split(s, ":")
	h, m, s := time[0], time[1], time[2]
	day := strings.HasSuffix(s, "AM")
	night := strings.HasSuffix(s, "PM")

	fh, _ := strconv.Atoi(h)
	if night && fh < 12 {
		fh += 12
	}

	if day && fh == 12 {
		fh = 0
	}

	return fmt.Sprintf("%02d:%s:%s", fh, m, s[:2])
}

// Create a function that takes a positive integer and returns a boolean that then given
// positive integer is palindrome or not. For example, if a value is palindrome return `true`:
// Do Not convert value to string
// Example
// 121 ==> true
// 1221 ==> true
// 2579752 ==> true
// isPalindrome(num: 121)     // returns true
// isPalindrome(num: 1221)    // returns true
// isPalindrome(num: 2579752) // returns true
// If the digits are not a palindrome, return `false` :
// 98    ==> false
// 123   ==> false
// 12341 ==> false
// isPalindrome(num: 98)    // returns false
// isPalindrome(num: 123)   // returns false
// isPalindrome(num: 12341) // returns false
func isPalindrome(num int) bool {
	d := 0
	for i := num; i > 0; i /= 10 {
		d++
	}

	for i := 0; i < d/2; i++ {
		if (num/int(math.Pow10(i)))%10 != (num/int(math.Pow10(d-i-1)))%10 {
			return false
		}
	}

	return true
}
