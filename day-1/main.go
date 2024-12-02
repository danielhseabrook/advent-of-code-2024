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
	listA               []int
	listB               []int
	listDifference      int
	listDifferenceTotal int
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
		// fmt.Print("valB = ", valB, "\n")
		// fmt.Print("valA = ", valA, "\n")
		listA = append(listA, valA)
		listB = append(listB, valB)

	}
	sort.Ints(listA)
	sort.Ints(listB)

	for i := 0; i < len(listA); i++ {
		listDifference = listA[i] - listB[i]
		if listDifference < 0 {
			listDifference = listDifference * -1
		}
		listDifferenceTotal += listDifference
		fmt.Print(i, ". ", listA[i], "+", listB[i], "=", listA[i]-listB[i], "\n")
	}
	fmt.Print("listDifference = ", listDifferenceTotal, "\n")
	if err := line.Err(); err != nil {
		log.Fatal(err)
	}
}
