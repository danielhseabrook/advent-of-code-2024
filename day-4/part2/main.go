package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	p1Count int = 0
	p2Count int = 0
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	line := bufio.NewScanner(file)
	lines := make([][]string, 0)

	for line.Scan() {
		lines = append(lines, strings.Split(line.Text(), ""))
	}
	// Part 1 Solution
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == "X" {
				if j >= 3 { // Left
					if lines[i][j-1] == "M" && lines[i][j-2] == "A" && lines[i][j-3] == "S" {
						p1Count++
					}
				}
				if j <= len(lines[i])-4 { // Right
					if lines[i][j+1] == "M" && lines[i][j+2] == "A" && lines[i][j+3] == "S" {
						p1Count++
					}
				}
				if i >= 3 { // Up
					if lines[i-1][j] == "M" && lines[i-2][j] == "A" && lines[i-3][j] == "S" {
						p1Count++
					}
				}
				if i <= len(lines)-4 { // Down
					if lines[i+1][j] == "M" && lines[i+2][j] == "A" && lines[i+3][j] == "S" {
						p1Count++
					}
				}
				if i >= 3 && j <= len(lines[i])-4 { // Diagonal Right Up
					if lines[i-1][j+1] == "M" && lines[i-2][j+2] == "A" && lines[i-3][j+3] == "S" {
						p1Count++
					}
				}
				if i >= 3 && j >= 3 { // Diagonal Left Up
					if lines[i-1][j-1] == "M" && lines[i-2][j-2] == "A" && lines[i-3][j-3] == "S" {
						p1Count++
					}
				}
				if i <= len(lines)-4 && j >= 3 { // Diagonal Down Left
					if lines[i+1][j-1] == "M" && lines[i+2][j-2] == "A" && lines[i+3][j-3] == "S" {
						p1Count++
					}
				}

				if i <= len(lines)-4 && j <= len(lines[i])-4 {
					if lines[i+1][j+1] == "M" && lines[i+2][j+2] == "A" && lines[i+3][j+3] == "S" {
						p1Count++
					}
				}
			}
		}
	}
	fmt.Println("Part 1: ", p1Count)

	// Part 2 Solution

	for i := 1; i < len(lines); i++ {
		for j := 1; j < len(lines[i]); j++ {
			if lines[i][j] == "A" {
				if i < len(lines)-1 && j < len(lines[i])-1 {
					neighbors := lines[i-1][j-1] + lines[i-1][j+1] + lines[i+1][j-1] + lines[i+1][j+1]
					if neighbors == "MSMS" || neighbors == "MMSS" || neighbors == "SMSM" || neighbors == "SSMM" {
						p2Count++
					}
				}
			}
		}
	}
	fmt.Println("Part 2: ", p2Count)
}
