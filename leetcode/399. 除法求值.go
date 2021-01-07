package main

import "fmt"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	dd := make(map[string]map[string]float64)
	for i := 0; i < len(equations); i++ {
		x, y := equations[i][0], equations[i][1]
		v := values[i]
		if dd[x] == nil {
			ddX := make(map[string]float64)
			dd[x] = ddX
		}
		ddX := dd[x]
		ddX[y] = v
		dd[x] = ddX
		if dd[y] == nil {
			ddY := make(map[string]float64)
			dd[y] = ddY
		}
		ddY := dd[y]
		ddY[x] = 1 / v
		dd[y] = ddY
	}
	r := []float64{}
	for _, queried := range queries {
		start, end := queried[0], queried[1]
		way := make(map[string]bool, 0)
		qr := dfs(start, end, way, dd)
		r = append(r, qr)
	}
	return r
}

func dfs(start string, end string, way map[string]bool, dd map[string]map[string]float64) float64 {
	way[start] = true
	if dd[start] == nil {
		return -1
	}
	if start == end {
		return 1
	}
	for key, _ := range dd[start] {
		if key == end {
			return dd[start][key]
		}
		if way[key] == false {
			way[key] = true
			v := dfs(key, end, way, dd)
			if v != -1 {
				return v * dd[start][key]
			}
		} else {

		}
	}
	return -1
}

func main() {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	r := calcEquation(equations, values, queries)
	fmt.Println(r)
}
