package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to open input")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var first_col []int64
	var sec_col []int64

	for scanner.Scan() {
		re := regexp.MustCompile(`\d+`)
		var temp_line []string = re.FindAllString(scanner.Text(), -1)

		number, err := strconv.ParseInt(temp_line[0], 0, 64)
		sec_num, err_sec := strconv.ParseInt(temp_line[1], 10, 64)

		if err != nil || err_sec != nil {
			log.Fatalf("Cannot convert")
		}

		first_col = append(first_col, number)
		sec_col = append(sec_col, sec_num)
	}

	slices.Sort(first_col)
	slices.Sort(sec_col)

	//puzzle 1
	var distance int
	if len(first_col) == len(sec_col) {
		for i := 0; i < len(sec_col); i++ {
			if sec_col[i]-first_col[i] < 0 {
				distance += int(first_col[i] - sec_col[i])
			} else {
				distance += int(sec_col[i] - first_col[i])
			}
		}
	}

	var simliarity int64

	dict := make(map[int64]int64)
	for _, num := range sec_col {
		dict[num]++
	}

	for _, num := range first_col {
		occurence, ok := dict[num]
		if ok {
			simliarity += num * occurence
		}
	}

	file.Close()

	fmt.Print("Distance ", distance)
	fmt.Print("Simliar score ", simliarity)

}
