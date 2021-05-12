package main

func decode(encoded []int) []int {
	var r []int
	n := len(encoded)
	if n < 2 {
		return r
	}
	p2_n := 0
	for i := 1; i < n+2; i++ {
		p2_n ^= i
	}
	p1_n := 0
	for i := 1; i < n; i += 2 {
		p1_n ^= encoded[i]
	}
	r = append(r, p2_n^p1_n)
	for _, e := range encoded {
		num := r[len(r)-1]
		r = append(r, num^e)
	}
	return r
}
