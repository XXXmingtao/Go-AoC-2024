package main

type coordinate struct {
	x int
	y int
}

func (cord *coordinate) equal(compare *coordinate) bool {
	if cord.x == compare.x &&
		cord.y == compare.y {
		return true
	}

	return false
}
