package main

func maxPoints(points [][]int) int {
	n := len(points)
	// 优化
	if n == 1 || n == 2 {
		return n
	}
	r := 1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x1, y1 := points[i][0], points[i][1]
			x2, y2 := points[j][0], points[j][1]
			a := x1 - x2
			b := y1 - y2
			tmp := 2
			for k := j + 1; k < n; k++ {
				x3, y3 := points[k][0], points[k][1]
				c := x2 - x3
				d := y2 - y3
				if a*d == c*b {
					tmp += 1
				}
			}
			if r < tmp {
				r = tmp
			}
			// 优化
			if r > n/2 {
				break
			}
		}
	}
	return r
}
