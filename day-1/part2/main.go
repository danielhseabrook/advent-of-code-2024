package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	listA           []int
	listB           []int
	similarityScore int
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	line := bufio.NewScanner(file)

	for line.Scan() {
		sepLists := strings.Split(line.Text(), "   ")
		valA, err := strconv.Atoi(sepLists[0])
		if err != nil {
			log.Fatal(err)
		}
		valB, err := strconv.Atoi(sepLists[1])
		listA = append(listA, valA)
		listB = append(listB, valB)

	}
	sort.Ints(listA)
	sort.Ints(listB)

	locationMap := make(map[int]int)
	for _, num := range listB {
		locationMap[num]++
	}

	for location, occurrence := range locationMap {
		for i := 0; i < len(listA); i++ {
			if listA[i] == location {
				similarityScore += (location * occurrence)
			}
		}
	}
	fmt.Print(similarityScore)
}
