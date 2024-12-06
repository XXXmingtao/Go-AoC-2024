package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to open input")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var rules = make(map[string][]string)
	var orders [][]string
	for scanner.Scan() {

		if strings.Contains(scanner.Text(), ",") {
			input := strings.Split(scanner.Text(), ",")
			orders = append(orders, input)

		} else if strings.Contains(scanner.Text(), "|") {
			rule := strings.Split(scanner.Text(), "|")

			key := rule[0]
			value := rule[1]

			rules[key] = append(rules[key], value)
		}
	}

	file.Close()

	var validOrders [][]string
	var totalPages int
	var incorrectPagesUpdates int
	for _, order := range orders {
		var middle int
		isValid := true
		for i := 0; i <= len(order)-1; i++ {
			if len(order)%2 != 0 {
				if i == (len(order)-1)/2 {
					middle, err = strconv.Atoi(order[i])
				}
			}

			if i == len(order)-1 {
				validOrders = append(validOrders, order)
				if err == nil {
					// fmt.Println(middle)
					totalPages += middle
				}
				break
			}

			if len(rules[order[i+1]]) > 0 {
				if checkPrequsite(rules[order[i+1]], order[i]) {
					// fmt.Println("not valid") : true
					isValid = false
					break
				}
			}
		}

		if !isValid {

			newList := MyLinkedList{}
			newList.insertAtFront(order[0])

			for i := 1; i < len(order); i++ {

				if len(rules[order[i]]) == 0 {
					//check if node is in this key
					newList.insertAtLast(order[i])
				} else if checkPrequsite(rules[newList.getLast().data], order[i]) {
					newList.insertAfterValue(newList.getLast().data, order[i])
				} else {
					tempNode := newList.head
					for tempNode != nil {
						if checkPrequsite(rules[order[i]], tempNode.data) {
							newList.inserBeforeValue(tempNode.data, order[i])
							break
						}
						tempNode = tempNode.next
					}
					// newList.inserBeforeValue(newList.getLast().data, order[i])
				}
			}
			midPoint := (len(order))/2 + 1
			correctedNum, err := strconv.Atoi(newList.findNodeAt(midPoint).data)
			if err == nil {
				fmt.Println("correct: ", correctedNum)
				incorrectPagesUpdates += correctedNum
			}

			fmt.Println("Original: ", order)
			newList.Print()
		}

	}

	// for key := range rules {
	// 	fmt.Println(key, " | ", rules[key])
	// }

	// fmt.Println(orders)
	//fmt.Println("valid, ", validOrders)
	fmt.Println("total pages, ", totalPages)
	fmt.Println("incorrectPagesUpdates number", incorrectPagesUpdates)
}

func checkPrequsite(rule []string, page string) bool {

	for _, i := range rule {
		if i == page {
			return true
		}
	}

	return false
}
