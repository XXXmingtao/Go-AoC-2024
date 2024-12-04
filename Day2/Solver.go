package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to open input")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var reports [][]int64

	for scanner.Scan() {
		re := regexp.MustCompile(`\d+`)
		var str_floors []string = re.FindAllString(scanner.Text(), -1)
		var int_floors []int64

		for _, i := range str_floors {
			i, err := strconv.ParseInt(i, 10, 64)
			if err != nil {
				log.Fatalf("cannot convert")
			}
			int_floors = append(int_floors, i)
		}

		reports = append(reports, int_floors)
	}

	var safeFloorCount int64
	var failedReports [][]int64
	var reportsWithIndex = make(map[int][][]int64)

	for _, floor := range reports {
		var isFloorSafe bool
		var errorIndex []int

		isFloorSafe = checkFloors(floor, reportsWithIndex)

		if isFloorSafe && errorIndex == nil {
			safeFloorCount++
		} else {
			failedReports = append(failedReports, floor)
		}
	}

	file.Close()

	var failed int
	for key, value := range reportsWithIndex {
		failed = failed + len(reportsWithIndex[key])
		for _, floor := range value {
			var firstArray []int64
			var secondArray []int64
			var thirdArray []int64
			firstArray = append(firstArray, floor...)
			secondArray = append(secondArray, floor...)
			thirdArray = append(thirdArray, floor...)

			firstArray = removeIndex(firstArray, key)
			secondArray = removeIndex(secondArray, key+1)
			thirdArray = removeIndex(thirdArray, key-1)

			if checkFloors(firstArray, map[int][][]int64{}) || checkFloors(secondArray, map[int][][]int64{}) || checkFloors(thirdArray, map[int][][]int64{}) {
				safeFloorCount++
			}
		}
	}

	fmt.Println("bads: ", failed)
	fmt.Println("safe reports+++: ", safeFloorCount)

}

func checkFloors(floor []int64, failedWithIndex map[int][][]int64) bool {

	var isSafe bool
	// var errorIndex []int
	isAscending := floor[0]-floor[1] < 0

	for i := 0; i <= len(floor)-2; i++ {

		if isAscending {
			if floor[i]-floor[i+1] > 0 {
				failedWithIndex[i] = append(failedWithIndex[i], floor)
				//fmt.Println("floor is descending but we going up, unsafe")
				break
			}
		} else {
			if floor[i]-floor[i+1] < 0 {
				failedWithIndex[i] = append(failedWithIndex[i], floor)
				//fmt.Println("floor is ascending but we going down, unsafe")
				break
			}
		}

		gap := math.Abs(float64(floor[i] - floor[i+1]))

		if gap <= 3 && gap > 0 {
			if i == len(floor)-2 {
				//fmt.Println("report is safe")
				isSafe = true
			}
			continue
		} else {
			failedWithIndex[i] = append(failedWithIndex[i], floor)
			//fmt.Println("report is unsafe")
			break
		}
	}
	return isSafe
}

func removeIndex(floor []int64, index int) []int64 {
	var removedArray []int64
	if index == len(floor) {
		index = len(floor) - 1
	} else if index < 0 {
		removedArray = floor[1:]
		return removedArray
	}
	removedArray = append(floor[:index], floor[index+1:]...)
	return removedArray
}
