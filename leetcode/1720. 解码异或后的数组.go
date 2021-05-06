package main

func decode(encoded []int, first int) []int {
	var res []int
	res = append(res, first)
	for _, e := range encoded {
		end := res[len(res)-1]
		res = append(res, end^e)
	}
	return res
}
