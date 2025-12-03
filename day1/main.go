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
	const size = 100
	counter := 0
	starto := 50

	file, err := os.Open("input_day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "R") {
			r_case, _ := strconv.Atoi(strings.TrimPrefix(line, "R"))
			starto = (starto + r_case) % size

		} else if strings.HasPrefix(line, "L") {
			l_case, _ := strconv.Atoi(strings.TrimPrefix(line, "L"))
			starto = (starto - (l_case % size) + size) % size
		}
		if starto == 0 {
			counter++
		}
	}
	fmt.Println(counter)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
