package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to open input")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var cityMap [][]string
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), "")
		cityMap = append(cityMap, text)
	}

	file.Close()

	antennas := make(map[string][]coordinate)

	rege := regexp.MustCompile(`\w`)

	//Grab all antennas
	for i := 0; i < len(cityMap); i++ {
		for j := 0; j < len(cityMap[i]); j++ {
			if rege.MatchString(cityMap[i][j]) {
				coord := coordinate{i, j}
				key := cityMap[i][j]
				antennas[key] = append(antennas[key], coord)
			}
		}
	}

	for _, line := range cityMap {
		fmt.Println(line)
	}

	var antiNodesList []coordinate
	for key := range antennas {
		fmt.Println(key, ": ", antennas[key])
		inputs := antennas[key]
		for i := 0; i < len(inputs); i++ {
			for j := i + 1; j < len(inputs); j++ {
				workOutAntiNodes(inputs[i], inputs[j], cityMap, &antiNodesList)
			}
		}
	}

	// workOutAntiNodes(coordinate{3, 4}, coordinate{5, 5}, cityMap, &antiNodesList)

	// workOutAllTheWay(2, 1, 2, 1, &antiNodesList)
	// fmt.Println(antiNodesList)
	fmt.Println(len(antiNodesList))

}

func workOutAllTheWay(coordniate_x int, coordinate_y int, distance_x int, distance_y int, resultList *[]coordinate, cityMap [][]string) {
	if coordniate_x < 0 || coordniate_x > len(cityMap)-1 || coordinate_y < 0 || coordinate_y > len(cityMap[0])-1 {
		return
	}

	antiNode := coordinate{coordniate_x, coordinate_y}
	canAdd := true
	for _, item := range *resultList {
		if antiNode.equal(&item) {
			canAdd = false
		}
	}

	if canAdd {
		*resultList = append(*resultList, antiNode)
		fmt.Println("x:", coordniate_x, "y: ", coordinate_y)
	}

	workOutAllTheWay(coordniate_x+distance_x, coordinate_y+distance_y, distance_x, distance_y, *&resultList, cityMap)
}

func workOutAntiNodes(first coordinate, second coordinate, cityMap [][]string, resultList *[]coordinate) {

	//first antiNode
	// canAddOriginalOne := true
	// canAddOriginaltwo := true
	// for _, item := range *resultList {
	// 	if item.equal(&first) {
	// 		canAddOriginalOne = false
	// 	}

	// 	if item.equal(&second) {
	// 		canAddOriginaltwo = false
	// 	}
	// }

	// if canAddOriginalOne {
	// 	*resultList = append(*resultList, first)
	// }

	// if canAddOriginaltwo {
	// 	*resultList = append(*resultList, second)
	// }

	first_antiNodes_distance_x := first.x - second.x
	first_antiNodes_distance_y := first.y - second.y
	// antiNode_x := first.x + first_antiNodes_distance_x
	// antiNode_y := first.y + first_antiNodes_distance_y
	// first_antiNode := coordinate{antiNode_x, antiNode_y}
	// fmt.Println("first antiNode: ", antiNode_x, antiNode_y)
	fmt.Println("distance_x: ", first_antiNodes_distance_x)
	fmt.Println("distance_y: ", first_antiNodes_distance_y)

	//second antiNode
	second_antiNode_distance_x := second.x - first.x
	second_antiNode_distance_y := second.y - first.y
	// sec_antiNode_x := second.x + second_antiNode_distance_x
	// sec_antiNode_y := second.y + second_antiNode_distance_y
	// sec_antiNode := coordinate{sec_antiNode_x, sec_antiNode_y}
	// fmt.Println("second antiNode: ", sec_antiNode_x, sec_antiNode_y)

	// if antiNode_x >= 0 && antiNode_x <= len(cityMap)-1 &&
	// 	antiNode_y >= 0 && antiNode_y <= len(cityMap[0])-1 {
	// 	canAddFirst := true
	// 	for _, item := range *resultList {
	// 		if item.equal(&first_antiNode) {
	// 			canAddFirst = false
	// 		}
	// 	}

	// 	if canAddFirst {
	// 		*resultList = append(*resultList, first_antiNode)
	// 	}
	// }

	workOutAllTheWay(first.x, first.y, first_antiNodes_distance_x, first_antiNodes_distance_y, resultList, cityMap)
	workOutAllTheWay(first.x, first.y, -first_antiNodes_distance_x, -first_antiNodes_distance_y, resultList, cityMap)

	// if sec_antiNode_x >= 0 && sec_antiNode_x <= len(cityMap)-1 &&
	// 	sec_antiNode_y >= 0 && sec_antiNode_y <= len(cityMap[0])-1 {

	// 	canAddSecond := true
	// 	for _, item := range *resultList {
	// 		if item.equal(&sec_antiNode) {
	// 			canAddSecond = false
	// 		}
	// 	}

	// 	if canAddSecond {
	// 		*resultList = append(*resultList, sec_antiNode)
	// 	}
	// }
	fmt.Println("sec_distance_x: ", second_antiNode_distance_x)
	fmt.Println("sec_distance_y: ", second_antiNode_distance_y)

	workOutAllTheWay(second.x, second.y, second_antiNode_distance_x, second_antiNode_distance_y, resultList, cityMap)
	workOutAllTheWay(second.x, second.y, -second_antiNode_distance_x, -second_antiNode_distance_y, resultList, cityMap)
}
