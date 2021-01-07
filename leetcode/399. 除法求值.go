package main

import "fmt"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	dd := make(map[string]map[string]float64)
	for i := 0; i < len(equations); i++ {
		x, y := equations[i][0], equations[i][1]
		v := values[i]
		dd[x][y] = v
		dd[y][x] = 1 / v
	}
	fmt.Println(dd)
	return []float64{0.1}
}

func main() {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	r := calcEquation(equations, values, queries)
	fmt.Println(r)
}
