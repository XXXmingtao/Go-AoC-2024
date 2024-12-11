package main

import "fmt"

type position struct {
	x         int
	y         int
	direction string
}

func (pos *position) equal(compares *position) bool {
	if pos.x == compares.x &&
		pos.y == compares.y {
		return true
	}

	return false
}

func (pos *position) Print() {
	fmt.Printf("[%d, %d] - %s", pos.x, pos.y, pos.direction)
}
