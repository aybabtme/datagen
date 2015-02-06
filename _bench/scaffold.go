package bench

import (
	"strconv"
)

// scaffolding

func makeStrings(n int) []string {
	vals := make([]string, 0, n)
	for i := 0; i < n; i++ {
		vals = append(vals, strconv.Itoa(i))
	}
	return vals
}
func makeBytes(n int) [][]byte {
	vals := make([][]byte, 0, n)
	for i := 0; i < n; i++ {
		vals = append(vals, []byte(strconv.Itoa(i)))
	}
	return vals
}

func makeInts(n int) []int {
	vals := make([]int, 0, n)
	for i := 0; i < n; i++ {
		vals = append(vals, i)
	}
	return vals
}
func makeFloat64s(n int) []float64 {
	vals := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		vals = append(vals, float64(i))
	}
	return vals
}
