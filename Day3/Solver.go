package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	file, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatalf("Failed to open input")
	}

	var pureText = string(file)
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	var multisInput []string = re.FindAllString(pureText, -1)

	chars := strings.Split(pureText, "")

	total := processMultis(multisInput)
	fmt.Println("total: ", total)

	//var result int64
	var validMultis string
	start := true

	for i := 0; i < len(chars); i++ {
		var termination bool

		if string(chars[i]) == "d" {
			command := getNextFourChars(pureText, i)
			if command == "don't" {
				start = false
				termination = true
			} else if strings.Contains(command, "do") {
				start = true
				termination = false
			}
		}

		if !termination && start {
			validMultis = validMultis + chars[i]
		}

	}

	cleanMultis := extractMultis(validMultis)
	cleanResult := processMultis(cleanMultis)
	fmt.Println("clean result:", cleanResult)
}

func getNextFourChars(words string, index int) string {
	command := strings.TrimSpace(words[index : index+5])

	return command
}

func extractMultis(text string) []string {
	var multisInput []string

	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	multisInput = re.FindAllString(text, -1)

	return multisInput
}

func processMultis(inputs []string) int64 {
	re := regexp.MustCompile(`\d+`)
	var total int64
	for _, words := range inputs {
		numbersInWords := re.FindAllString(words, -1)
		var numbers []int64
		for _, number := range numbersInWords {
			if i, err := strconv.ParseInt(number, 10, 64); err == nil {
				numbers = append(numbers, i)
			}
		}
		total += numbers[0] * numbers[1]
	}

	return total
}
