package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main01() {
	file, err := os.Open("01-input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	max := 0
	curr := 0
	for scanner.Scan() {
		val := scanner.Text()
		if val == "" {
			if curr > max {
				max = curr
			}
			curr = 0
			continue
		}

		intVal, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		curr += intVal
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(max)
}

func main() {
	file, err := os.Open("01-input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	vals := make([]int, 3)
	curr := 0
	for scanner.Scan() {
		val := scanner.Text()
		if val == "" {
			vals = append(vals, curr)
			curr = 0
			continue
		}

		intVal, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		curr += intVal
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(vals)))

	fmt.Println(vals[:3])

	sum := 0
	for _, v := range vals[:3] {
		sum += v
	}

	fmt.Println(sum)

}
