package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var count int = 0

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
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == "X" {
				if j >= 3 { // Left
					if lines[i][j-1] == "M" && lines[i][j-2] == "A" && lines[i][j-3] == "S" {
						count++
					}
				}
				if j <= len(lines[i])-4 { // Right
					if lines[i][j+1] == "M" && lines[i][j+2] == "A" && lines[i][j+3] == "S" {
						count++
					}
				}
				if i >= 3 { // Up
					if lines[i-1][j] == "M" && lines[i-2][j] == "A" && lines[i-3][j] == "S" {
						count++
					}
				}
				if i <= len(lines)-4 { // Down
					if lines[i+1][j] == "M" && lines[i+2][j] == "A" && lines[i+3][j] == "S" {
						count++
					}
				}
				if i >= 3 && j <= len(lines[i])-4 { // Diagonal Right Up
					if lines[i-1][j+1] == "M" && lines[i-2][j+2] == "A" && lines[i-3][j+3] == "S" {
						count++
					}
				}
				if i >= 3 && j >= 3 { // Diagonal Left Up
					if lines[i-1][j-1] == "M" && lines[i-2][j-2] == "A" && lines[i-3][j-3] == "S" {
						count++
					}
				}
				if i <= len(lines)-4 && j >= 3 { // Diagonal Down Left
					if lines[i+1][j-1] == "M" && lines[i+2][j-2] == "A" && lines[i+3][j-3] == "S" {
						count++
					}
				}

				if i <= len(lines)-4 && j <= len(lines[i])-4 {
					if lines[i+1][j+1] == "M" && lines[i+2][j+2] == "A" && lines[i+3][j+3] == "S" {
						count++
					}
				}
			}
		}
	}
	print(count)
}
