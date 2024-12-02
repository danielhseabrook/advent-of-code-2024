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
		for i := 1; i < len(report); i++ {
			if report[i] > report[i-1] || report[i-1] < report[i] {
				if report[i] > report[i-1]+3 || report[i] == report[i-1] {
					safeReports[reportN] = false
					break
				} else {
					safeReports[reportN] = true
				}
			} // report test
		} // report
	} // line
	safeN := 0
	for _, v := range safeReports {
		if v {
			safeN++
		}
	}
	fmt.Println(safeN)
} // main