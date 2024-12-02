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

const (
	toleranceCap int = 1
	max          int = 3
)

func isRepeat(report []int, reportN int, tolerance *int) bool {
	for i := 1; i < len(report); i++ {
		if report[i] == report[i-1] {
			*tolerance++
			//			fmt.Printf("Repeat found at index %d: report=%d\n", i, reportN)
			if *tolerance > toleranceCap {
				return false
			}
		}
	}
	return false
}

func withinMax(report []int, max int, reportN int, tolerance *int) bool {
	for i := 1; i < len(report); i++ {
		if report[i] > report[i-1]+max || report[i] < report[i-1]-max {
			*tolerance++
			//			fmt.Printf("Max breach found at index %d and %d: report=%d\n", i, i-1, reportN)
			if *tolerance > toleranceCap {
				return false
			}
		}
	}
	return true
}

func isOrdered(report []int, reportN int, tolerance *int) bool {
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
		*tolerance++
		//		fmt.Printf("Report %d is not ordered\n", reportN)

		if *tolerance > toleranceCap {
			return false
		}
	}
	return true
}

func reportTest(report []int, reportN int, tolerance *int) bool {
	if withinMax(report, max, reportN, tolerance) && !isRepeat(report, reportN, tolerance) && isOrdered(report, reportN, tolerance) {
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
	noToleranceSafe := make(map[int]bool)
	reportN := 0
	for line.Scan() {
		tolerance := 0
		reportN++
		reportString := strings.Split(line.Text(), " ")
		report := []int{}
		for _, s := range reportString {
			level, _ := strconv.Atoi(s)
			report = append(report, level)
		}
		noToleranceSafe[reportN] = reportTest(report, reportN, &tolerance)
		reportTest(report, reportN, &tolerance)
		if tolerance <= toleranceCap {
			fmt.Printf("Report %d is apparrently safe with tolerance %d\n", reportN, tolerance)
			safeReports[reportN] = true
		}
	}
	toleranceSafeN := 0
	for _, v := range noToleranceSafe {
		if v {
			toleranceSafeN++
		}
	}
	noToleranceSafeN := 0
	for _, v := range safeReports {
		if v {
			noToleranceSafeN++
		}
	}
	fmt.Println("No Tolerance: ", noToleranceSafeN)
	fmt.Println("Tolerance: ", toleranceSafeN)
}
