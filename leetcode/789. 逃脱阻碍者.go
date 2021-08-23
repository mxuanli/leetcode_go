package main

func escapeGhosts(ghosts [][]int, target []int) bool {

	var abs func(x int) int
	abs = func(x int) int {
		if x < 0 {
			return 0 - x
		}
		return x
	}

	var manhattanDistance func(x []int, y []int) int
	manhattanDistance = func(x []int, y []int) int {
		return abs(x[0]-y[0]) + abs(x[1]-y[1])
	}

	playerDistance := manhattanDistance([]int{0, 0}, target)
	for _, ghost := range ghosts {
		if manhattanDistance(ghost, target) <= playerDistance {
			return false
		}
	}
	return true
}
