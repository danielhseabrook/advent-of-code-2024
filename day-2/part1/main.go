package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var safeReports int

const max = 3

func isRepeat(report []int, reportN int) bool {
	for i := 1; i < len(report); i++ {
		if report[i] == report[i-1] {
			fmt.Printf("Repeat found at index %d: report=%d\n", i, reportN)
			return true
		}
	}
	return false
}

func withinMax(report []int, max int) bool {
	for i := 1; i < len(report); i++ {
		if report[i] > report[i-1]+max || report[i] < report[i-1]-max {
			return false
		}
	}
	return true
}

func isOrdered(report []int, reportN int) bool {
	ascending := true
	descending := true

	for i := 1; i < len(report); i++ {
		if report[i] < report[i-1] {
			ascending = false
		}
		if report[i] > report[i-1] {
			descending = false
		}
	}

	if !ascending && !descending {
		fmt.Printf("Report %d is not ordered\n", reportN)
		return false
	}
	return true
}

func reportTest(report []int, reportN int) bool {
	if withinMax(report, max) && !isRepeat(report, reportN) && isOrdered(report, reportN) {
		return true
	}
	return false
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	line := bufio.NewScanner(file)

	safeReports := make(map[int]bool)
	reportN := 0
	for line.Scan() {
		reportN++
		reportString := strings.Split(line.Text(), " ")
		report := []int{}
		for _, s := range reportString {
			level, _ := strconv.Atoi(s)
			report = append(report, level)
		}
		safeReports[reportN] = reportTest(report, reportN)
	}
	safeN := 0
	for _, v := range safeReports {
		if v {
			safeN++
		}
	}
	fmt.Println(safeN)
}
