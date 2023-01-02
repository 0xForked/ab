package lorem_test

import (
	"testing"
)

// BenchmarkGridLakes-10    	 4357220	       261.3 ns/op
// BenchmarkGridLakes-10    	 4501647	       260.0 ns/op
func BenchmarkGridLakes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gridLake(5, 4, [][]string{
			{"#", ".", "#", "#", "#"},
			{".", ".", "#", "#", "#"},
			{"#", "#", ".", "#", "#"},
			{"#", "#", "#", "#", "."},
		})
	}
}

// BenchmarkGridLakes2-10    	 5374088	       208.4 ns/op
// BenchmarkGridLakes2-10    	 5616960	       205.6 ns/op
func BenchmarkGridLakes2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gridLake2(5, 4, [][]string{
			{"#", ".", "#", "#", "#"},
			{".", ".", "#", "#", "#"},
			{"#", "#", ".", "#", "#"},
			{"#", "#", "#", "#", "."},
		})
	}
}

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

func gridLake2(m int, n int, lakes [][]string) int {
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
					for _, d := range []int{-1, 0, 1, -1, 1, -1, 0, 1} {
						nr := cr + d/2
						nc := cc + d%2
						if nr < 0 || nr >= n || nc < 0 || nc >= m {
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
	return c
}
