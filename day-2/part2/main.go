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

func reportTest(report []int, reportN int) bool {
	if withinMax(report, max, reportN) && !isRepeat(report, reportN) && isOrdered(report, reportN) {
		fmt.Print(withinMax(report, max, reportN), !isRepeat(report, reportN), isOrdered(report, reportN), "\n")
		return true
	} else {
		return troubleShoot(report, reportN)
	}
}

func troubleShoot(report []int, reportN int) bool {
	for {
		safe := false
		for i := 0; i < len(report); i++ {
			holdingReport := make([]int, len(report))
			copy(holdingReport, report)
			holdingReport = append(holdingReport[:i], holdingReport[i+1:]...)

			if withinMax(holdingReport, max, reportN) && !isRepeat(holdingReport, reportN) && isOrdered(holdingReport, reportN) {
				safe = true
				return safe
			}
		}
		return false
	}
}

func isRepeat(report []int, reportN int) bool {
	for i := 1; i < len(report); i++ {
		if report[i] == report[i-1] {
			return true
		}
	}
	return false
}

func withinMax(report []int, max int, reportN int) bool {
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
		// fmt.Printf("Report %d is not ordered\n", reportN)
		return false
	}
	return true
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
	fmt.Print(safeReports)
	safeN := 0
	for _, v := range safeReports {
		if v {
			safeN++
		}
	}
	fmt.Println("\n", "\n", "\n", "\n", "\n", "\n", "\n", "\n", safeN)
}
