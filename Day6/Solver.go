package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to open input")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var maze [][]string
	var mazeQ2 [][]string
	for scanner.Scan() {
		maze = append(maze, strings.Split(scanner.Text(), ""))
		mazeQ2 = append(mazeQ2, strings.Split(scanner.Text(), ""))
	}

	file.Close()

	// fmt.Println(maze[39][46])
	// fmt.Println(maze[39][46])

	var guard_x, guard_y int
	guard_x = 39
	guard_y = 46

	var path []*position
	start_pos := &position{guard_x, guard_y, "^"}
	path = append(path, start_pos)
	// start_direction := "^"
	var turningPoint []position

	for guard_x != 0 && guard_y != 0 && guard_x != len(maze)-1 && guard_y != len(maze[0])-1 {

		isDirectionChanged, direction := changeDirection(maze[guard_x][guard_y], &maze, guard_x, guard_y, &turningPoint)

		// println("loop size", len(turningPoint))

		// if direction == "<" && isDirectionChanged && len(loop_obsctuction) == 3 {
		// 	placement_x := loop_obsctuction[2].x
		// 	placement_y := loop_obsctuction[1].y - loop_obsctuction[0].y
		// 	new_placement := position{placement_x, placement_y}
		// 	loop_obsctuction = append(loop_obsctuction, new_placement)
		// 	total_placements = append(total_placements, new_placement)
		// }

		// if len(loop_obsctuction) == 4 {
		// 	fmt.Println("loop cordinates: ")
		// 	for _, pos := range loop_obsctuction {
		// 		fmt.Printf("[%d, %d]", pos.x, pos.y)
		// 	}

		// 	loop_obsctuction = []position{}
		// }

		if !isDirectionChanged {
			isHereBefore := false
			new_cordinatesX, new_cordinatesY := move(maze[guard_x][guard_y], maze, guard_x, guard_y)
			//move to new position
			guard_x = new_cordinatesX
			guard_y = new_cordinatesY
			// fmt.Println(guard_x, guard_y, direction)
			// println("before", maze[guard_x][guard_y])

			// println()
			maze[guard_x][guard_y] = direction
			// println(guard_x, guard_y)
			pos := &position{new_cordinatesX, new_cordinatesY, direction}
			// fmt.Printf("new_cordinatesX: %d , new_cordinatesY: %d", new_cordinatesX, new_cordinatesY)
			for _, visted_pos := range path {
				if visted_pos.equal(pos) {
					isHereBefore = true
				}
			}

			if !isHereBefore {
			}
			path = append(path, pos)

		}
	}

	// for _, pos := range turningPoint {
	// 	pos.Print()
	// 	fmt.Println()
	// }

	// fmt.Println(maze)

	println(len(path))

	var total int
	// temp_maze := deepCopyMaze(mazeQ2)
	// mazeQ2[39][46] = "."

	for i := 0; i < len(path); i++ {
		if path[i].x == 7 && path[i].y == 6 {
			fmt.Println("index", i)
		}
	}
	obstructs_place := []position{}

	for i := 1; i < len(mazeQ2); i++ {
		start_point := &position{39, 46, "^"}
		for j := 0; j < len(mazeQ2[0]); j++ {
			var temp_maze = deepCopyMaze(mazeQ2)
			repeat := 0

			if temp_maze[i][j] == "." {
				temp_maze[i][j] = "O"
			}

			fmt.Println("currently at: path ", i, j)
			// for _, line := range temp_maze {
			// 	fmt.Println(line)
			// }
			if workOutMaze(path, &temp_maze, start_point, repeat) {
				total++
			}

			// for _, line := range temp_maze {
			// 	fmt.Println(line)
			// }
		}
	}

	for i := 0; i < 0; i++ {
		var temp_maze = deepCopyMaze(mazeQ2)
		repeat := 0
		temp_maze[path[i].x][path[i].y] = path[i].direction
		pter, next_x, next_y := getNextPosition(path[i].direction, *path[i], temp_maze)
		if *pter != "#" {
			*pter = "O"
			start_point := path[i]

			// fmt.Println("currently at: path ", i)
			if workOutMaze(path, &temp_maze, start_point, repeat) {
				// fmt.Println(start_point.x, start_point.y)
				duplicated := false
				for _, place := range obstructs_place {
					if place.x == next_x && place.y == next_y {
						duplicated = true
					}
				}

				if !duplicated {
					obstructs_place = append(obstructs_place, position{next_x, next_y, "O"})
				}

				// total++
				// for _, line := range temp_maze {
				// 	fmt.Println(line)
				// }
			}

		}

	}
	println("total obs", total)
	// println("total", len(obstructs_place))
	// possible_obstructs := findAllPlaceableObstructs(turningPoint, mazeQ2, path)
	// fmt.Println("outcome", possible_obstructs)

}

func move(direction string, maze [][]string, position_x int, position_y int) (int, int) {
	x := position_x
	y := position_y

	if direction == "^" {
		x--
	} else if direction == ">" {
		y++
	} else if direction == "v" {
		x++
	} else if direction == "<" {
		y--
	}
	return x, y
}

func changeDirection(original string, maze *[][]string, position_x int, position_y int, loop_obstructs *[]position) (bool, string) {
	var direction string
	var next_pos_item string
	var isDirectionChanged bool
	var turningPoint position

	if original == "^" {
		next_pos_item = (*maze)[position_x-1][position_y]
		direction = ">"
	} else if original == ">" {
		next_pos_item = (*maze)[position_x][position_y+1]
		direction = "v"
	} else if original == "v" {
		next_pos_item = (*maze)[position_x+1][position_y]
		direction = "<"
	} else if original == "<" {
		next_pos_item = (*maze)[position_x][position_y-1]
		direction = "^"
	}

	if next_pos_item == "#" || next_pos_item == "O" {
		(*maze)[position_x][position_y] = direction

		// println("adding turning point")
		turningPoint = position{position_x, position_y, direction}
		*loop_obstructs = append(*loop_obstructs, turningPoint)

		isDirectionChanged = true
	} else {
		return false, original
	}

	return isDirectionChanged, direction
}

func willGuardReachThisPoint(guard_direction string, guard_position position, targetPoint position, maze [][]string, traveled_points []*position) int {
	temp_maze := maze
	// guard_x := guard_position.x
	// guard_y := guard_position.y
	//temp_maze[guard_x][guard_y] = guard_direction
	println("traveled point: ", traveled_points)
	total_number := 0
	for temp_maze[targetPoint.x][targetPoint.y] != guard_direction {
		// for _, line := range temp_maze {
		// 	fmt.Println(line)
		// }
		if getNextPositionChar(guard_direction, guard_position, temp_maze) == "#" {
			break
		} else {
			x, y := move(guard_direction, temp_maze, guard_position.x, guard_position.y)
			// fmt.Println(x, y, guard_direction)
			guard_position.x = x
			guard_position.y = y
			temp_maze[x][y] = guard_direction

			if guard_direction == "<" {
				for i := x; i != -1; i-- {
					if (temp_maze[i][y]) == "#" {
						fmt.Println("point at: ", x, y)

						// temp_point := position{i + 1, y, ""}
						// fmt.Print("# key found")
						// temp_point.Print()
						temp_maze[x][y-1] = "#"
						if workOutMaze(traveled_points, &temp_maze, &guard_position, 0) {
							// beforeTravelPoint.Print()
							total_number++
						}
					}
				}
			} else if guard_direction == "^" {
				for i := y; i < len(temp_maze[x])-1; i++ {
					if temp_maze[x][i] == "#" {
						//temp_point := position{x, i - 1, ""}
						// fmt.Print("# key found")
						// temp_point.Print()
						temp_maze[x-1][y] = "#"
						beforeTravelPoint := guard_position
						if workOutMaze(traveled_points, &temp_maze, &guard_position, 0) {
							total_number++
							fmt.Println("obstrcut placeable at: ", beforeTravelPoint.x-1, beforeTravelPoint.y)
							break
						}
					}
				}
			} else if guard_direction == ">" {
				for i := x; i < len(temp_maze); i++ {
					if (temp_maze[i][y]) == "#" {
						//temp_point := position{i - 1, y, ""}
						// fmt.Print("# key found")
						// temp_point.Print()
						temp_maze[x][y+1] = "#"
						beforeTravelPoint := guard_position
						if workOutMaze(traveled_points, &temp_maze, &guard_position, 0) {
							total_number++
							fmt.Println("obstrcut placeable at: ", beforeTravelPoint.x, beforeTravelPoint.y+1)
							break
						}
					}
				}
			} else if guard_direction == "v" {
				for i := y; i != -1; i-- {
					// fmt.Println("should look to left: ", x, i)
					if (temp_maze[x][i]) == "#" {
						// temp_point := position{x, i + 1, ""}
						// fmt.Print("# key found")
						// temp_point.Print()
						temp_maze[x+1][y] = "#"
						beforeTravelPoint := guard_position
						if workOutMaze(traveled_points, &temp_maze, &guard_position, 0) {
							total_number++
							fmt.Println("obstrcut placeable at: ", beforeTravelPoint.x+1, beforeTravelPoint.y)
							break
						}
					}
				}
			}

			// if guard_position.x == targetPoint.x && guard_position.y == targetPoint.y {
			// 	total_number++
			// }
		}
	}

	return total_number

	// return false
}

func deepCopyMaze(maze [][]string) [][]string {

	matrix := make([][]string, len(maze))
	for i := 0; i < len(maze); i++ {
		matrix[i] = make([]string, len(maze[0]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j] = maze[i][j]
		}
	}

	return matrix
}

func workOutMaze(traveledPath []*position, maze *[][]string, startPosition *position, repeat int) bool {

	// start_x := startPosition.x
	// start_y := startPosition.y
	// original_dirc := startPosition.direction

	guard_x := startPosition.x
	guard_y := startPosition.y
	var obstrcuts []position

	de_maze := *maze
	times := 0
	// var ch string
	this_line := []position{}
	for guard_x != 0 && guard_y != 0 && guard_x != len(de_maze)-1 && guard_y != len(de_maze[0])-1 {
		// println("before", guard_x, guard_y)
		isHereBefore := false
		// fmt.Println(ch)
		isDirectionChanged, direction := changeDirection(de_maze[guard_x][guard_y], &de_maze, guard_x, guard_y, &obstrcuts)

		//checkPoint := &position{guard_x, guard_y, direction}

		// if guard_x == start_x && guard_y == start_y {
		// 	repeat++
		// }

		// if checkPoint.x == start_x && checkPoint.y == start_y && (ch == "#" || ch == "O") {
		// 	times++
		// }

		// fmt.Println(ch)
		if times >= 2 || repeat == 2 {
			return true
		}

		if !isDirectionChanged {

			// println(guard_x, guard_y)
			new_cordinatesX, new_cordinatesY := move(de_maze[guard_x][guard_y], de_maze, guard_x, guard_y)
			guard_x = new_cordinatesX
			guard_y = new_cordinatesY
			de_maze[guard_x][guard_y] = direction

			// ch = getNextPositionChar(de_maze[guard_x][guard_y], position{guard_x, guard_y, de_maze[guard_x][guard_y]}, de_maze)

			// fmt.Println("start point", start_x, start_y)
			// fmt.Println("check point", guard_x, guard_y)
			// fmt.Println("original direction", original_dirc)
			// fmt.Println(ch)
			// println(checkPoint.x == start_x && checkPoint.y == start_y)

			// for _, point := range traveledPath {
			// 	if point.equal(checkPoint) && ch == "#" {
			// 		return true
			// 	}
			// }
			// var checkThisPostion position
			for _, visted_pos := range this_line {
				if visted_pos.equal(&position{guard_x, guard_y, direction}) && visted_pos.direction == direction {
					fmt.Println(visted_pos.x, visted_pos.y, visted_pos.direction)
					fmt.Println(guard_x, guard_y, direction)
					isHereBefore = true
					// checkThisPostion = visted_pos
				}
			}

			// if ch == "O" {
			// 	times++
			// }
			// checkThisPostion.direction == direction && (ch == "#" || ch == "O")
			if isHereBefore {
				// println("did we arrive here")
				return true
			} else {
				this_line = append(this_line, position{guard_x, guard_y, direction})
				// println("Length: ", len(this_line))
			}

		}

		// for _, line := range de_maze {
		// 	fmt.Println(line)
		// }

		// fmt.Println("------------------------------------------------------------------------")

	}

	return false
}

func checkContent(posList []position, pos position) bool {
	for _, point := range posList {
		if point.equal(&pos) {
			return true
		}
	}

	return false
}

func findAllPlaceableObstructs(allTurningPoint []position, maze [][]string, path []*position) int {
	total_possible := 0
	for i := 0; i < len(allTurningPoint)-2; i++ {
		current_turningSeq := []position{allTurningPoint[i], allTurningPoint[i+1], allTurningPoint[i+2]}
		// traveled_points = append(traveled_points, allTurningPoint[i], allTurningPoint[i+1], allTurningPoint[i+2])
		fmt.Println(current_turningSeq)

		start_point := current_turningSeq[2]
		// placement_x := current_turningSeq[2].x
		// placement_y := current_turningSeq[1].y - current_turningSeq[0].y

		target_point := workOutTargetPoint(current_turningSeq)

		fmt.Println("start point", start_point.x, start_point.y)
		fmt.Println("target", target_point.x, target_point.y)

		total_possible += willGuardReachThisPoint(start_point.direction, start_point, target_point, maze, path)
		fmt.Println("total possible:", total_possible)
	}

	return total_possible
}

func workOutTargetPoint(turningSeq []position) position {
	current_turningSeq := []position{turningSeq[0], turningSeq[1], turningSeq[2]}
	start_point := current_turningSeq[2]
	var placement_x int
	var placement_y int

	if start_point.direction == "^" {
		placement_y = start_point.y
		placement_x = current_turningSeq[0].x
	}

	if start_point.direction == "<" {
		placement_x = start_point.x
		placement_y = current_turningSeq[0].y
	}

	if start_point.direction == ">" {
		placement_x = start_point.x
		placement_y = current_turningSeq[0].y
	}

	if start_point.direction == "v" {
		placement_y = start_point.y
		placement_x = current_turningSeq[0].x
	}

	return position{placement_x, placement_y, ""}
}

// func createObsuctByturningPoint(direction string, placement_x int, placement_y int) position {

// 	var obstruct position

// 	if direction == "^" {
// 		obstruct = position{placement_x - 1, placement_y}
// 	} else if direction == ">" {
// 		obstruct = position{placement_x, placement_y + 1}
// 	} else if direction == "v" {
// 		obstruct = position{placement_x + 1, placement_y}
// 	} else if direction == "<" {
// 		obstruct = position{placement_x, placement_y - 1}
// 	}

// 	return obstruct
// }

func getNextPositionChar(direction string, guard_position position, maze [][]string) string {

	var next_pos_item string

	if direction == "^" {
		if guard_position.x == 0 {
			return ""
		}
		next_pos_item = maze[guard_position.x-1][guard_position.y]
	} else if direction == ">" {
		if guard_position.y == len(maze[0])-1 {
			return ""
		}
		next_pos_item = maze[guard_position.x][guard_position.y+1]
	} else if direction == "v" {
		if guard_position.x == len(maze)-1 {
			return ""
		}
		next_pos_item = maze[guard_position.x+1][guard_position.y]
	} else if direction == "<" {
		if guard_position.y == 0 {
			return ""
		}
		next_pos_item = maze[guard_position.x][guard_position.y-1]
	}

	return next_pos_item
}

func getNextPosition(direction string, guard_position position, maze [][]string) (*string, int, int) {
	var next_pos_item *string
	var next_x int
	var next_y int
	if direction == "^" {
		next_pos_item = &maze[guard_position.x-1][guard_position.y]
		next_x = guard_position.x - 1
		next_y = guard_position.y
	} else if direction == ">" {
		next_pos_item = &maze[guard_position.x][guard_position.y+1]
		next_x = guard_position.x
		next_y = guard_position.y + 1
	} else if direction == "v" {
		next_pos_item = &maze[guard_position.x+1][guard_position.y]
		next_x = guard_position.x + 1
		next_y = guard_position.y
	} else if direction == "<" {
		next_pos_item = &maze[guard_position.x][guard_position.y-1]
		next_x = guard_position.x
		next_y = guard_position.y - 1
	}

	return next_pos_item, next_x, next_y

}
