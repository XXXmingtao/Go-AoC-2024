package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to open input")
	}
	var result []int
	var equations []MyLinkedList
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ":")

		anwser, err := strconv.Atoi(text[0])
		if err != nil {
			fmt.Print("failed to convert  result")
		}
		result = append(result, anwser)

		rege := regexp.MustCompile(`\d+`)
		numbers := rege.FindAllString(text[1], -1)
		equation := MyLinkedList{}
		for _, number := range numbers {
			real_number, err := strconv.Atoi(number)
			if err != nil {
				fmt.Print("failed to convert number")
			}
			equation.insertAtLast(real_number)
		}
		equations = append(equations, equation)
	}

	var total int
	for index := 0; index < len(result); index++ {
		final := result[index]
		caliEquation := equations[index]

		var operations [][]string

		var rawOperators []string
		generateOperations(&rawOperators, "", caliEquation.size()-1)

		for _, line := range rawOperators {
			array := strings.Split(line, "")
			operations = append(operations, array)
		}

		for i := 0; i < len(operations); i++ {
			temp_result := caliEquation.head.data
			node := caliEquation.head

			for j := 0; j < len(operations[i]); j++ {
				if node.next != nil {
					temp_result = calculate(temp_result, node.next.data, operations[i][j])
					node = node.next
				}
			}

			if final == temp_result {
				total += final
				break
			}
		}
	}

	fmt.Println("total: ", total)

}

func calculate(x int, y int, operator string) int {
	if operator == "*" {
		return x * y
	} else {
		return x + y
	}
}

func generateOperations(operations *[]string, operation string, length int) {

	if len(operation) == length {
		// addArrayOperation(operations, operation)
		*operations = append(*operations, operation)
		return
	}

	generateOperations(operations, operation+"+", length)
	generateOperations(operations, operation+"*", length)
}
