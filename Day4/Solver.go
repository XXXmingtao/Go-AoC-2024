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
	var matrix []string

	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}

	var total int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			xmasValue := checkXmasCordniates(matrix, i, j)
			if strings.Count(xmasValue, "MAS") >= 2 {
				total++
			}
		}
	}

	fmt.Println("total: ", total)

	// for _, line := range matrix {
	// 	fmt.Println(line)
	// }
}

func checkXmasCordniates(martrix []string, i int, j int) string {
	var validSequence [][][]int
	if martrix[i][j] == 'A' {

		if i-1 >= 0 && j-1 >= 0 && i+2 <= len(martrix) && j+2 <= len(martrix[i]) {
			validUpLeftForward := [][]int{{i - 1, j - 1}, {i, j}, {i + 1, j + 1}}
			validUpleftBackward := [][]int{{i + 1, j + 1}, {i, j}, {i - 1, j - 1}}
			validSequence = append(validSequence, validUpLeftForward)
			validSequence = append(validSequence, validUpleftBackward)
			validDownRightForward := [][]int{{i + 1, j - 1}, {i, j}, {i - 1, j + 1}}
			validDownRightBackwards := [][]int{{i - 1, j + 1}, {i, j}, {i + 1, j - 1}}
			validSequence = append(validSequence, validDownRightForward)
			validSequence = append(validSequence, validDownRightBackwards)
		}
	}

	var msg string
	for _, sequence := range validSequence {
		for _, cord := range sequence {
			msg = msg + retriveChar(martrix, cord)
		}
		msg = msg + ","
	}

	return msg
}

func checkCordniates(martrix []string, i int, j int) int {
	var validSequence [][][]int

	if martrix[i][j] == 'X' {

		if i+4 <= len(martrix) {
			validPosDown := [][]int{{i, j}, {i + 1, j}, {i + 2, j}, {i + 3, j}}
			validSequence = append(validSequence, validPosDown)
		}

		if i-3 >= 0 {
			validPosUp := [][]int{{i, j}, {i - 1, j}, {i - 2, j}, {i - 3, j}}
			validSequence = append(validSequence, validPosUp)
		}

		if j+4 <= len(martrix[i]) {
			validPosRight := [][]int{{i, j}, {i, j + 1}, {i, j + 2}, {i, j + 3}}
			validSequence = append(validSequence, validPosRight)
		}

		if j-3 >= 0 {
			validPosLeft := [][]int{{i, j}, {i, j - 1}, {i, j - 2}, {i, j - 3}}
			validSequence = append(validSequence, validPosLeft)
		}

		if j-3 >= 0 && i+4 <= len(martrix) {
			validPosDownLeftDia := [][]int{{i, j}, {i + 1, j - 1}, {i + 2, j - 2}, {i + 3, j - 3}}

			validSequence = append(validSequence, validPosDownLeftDia)
		}

		if i+4 <= len(martrix) && j+4 <= len(martrix[i]) {
			validPosDownRightDia := [][]int{{i, j}, {i + 1, j + 1}, {i + 2, j + 2}, {i + 3, j + 3}}
			validSequence = append(validSequence, validPosDownRightDia)
		}

		if j+4 <= len(martrix[i]) && i-3 >= 0 {
			validPosUpRightDia := [][]int{{i, j}, {i - 1, j + 1}, {i - 2, j + 2}, {i - 3, j + 3}}
			validSequence = append(validSequence, validPosUpRightDia)
		}

		if i-3 >= 0 && j-3 >= 0 {
			validUpLeftDia := [][]int{{i, j}, {i - 1, j - 1}, {i - 2, j - 2}, {i - 3, j - 3}}
			validSequence = append(validSequence, validUpLeftDia)
		}
	}

	var msg string
	for _, sequence := range validSequence {
		for _, cord := range sequence {
			msg = msg + retriveChar(martrix, cord)
		}

		msg = msg + ","
	}

	numbers := strings.Count(msg, "XMAS")

	return numbers
}

func retriveChar(matrix []string, cord []int) string {
	data := matrix[cord[0]][cord[1]]
	return string(data)
}
