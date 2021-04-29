package main

func canCross(stones []int) bool {
	stonesMap := make(map[int]bool)
	for _, s := range stones {
		stonesMap[s] = true

	}
	n := len(stones)
	hashMap := make(map[[2]int]bool)
	diff := []int{-1, 0, 1}
	var foo func(stone int, step int) bool
	foo = func(stone int, step int) bool {
		if stone == stones[n-1] {
			return true
		}
		if _, ok := hashMap[[2]int{stone, step}]; ok {
			return hashMap[[2]int{stone, step}]
		}
		for _, i := range diff {
			_, ok := stonesMap[i+step+stone]
			if i+step >= 1 && ok {
				if foo(i+step+stone, i+step) {
					hashMap[[2]int{stone, step}] = true
					return true
				}
			}
		}
		hashMap[[2]int{stone, step}] = false
		return false
	}
	return foo(0, 0)
}

func main() {

}
