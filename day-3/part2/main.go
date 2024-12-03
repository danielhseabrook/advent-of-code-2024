package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	total   int
	matches []string
	action  string = "do"
)
var allMatches [][]string

func main() {
	file, err := os.Open("input") // Reading the input file
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	line := bufio.NewScanner(file)

	re := regexp.MustCompile(`(mul\([0-9]{1,3}\,[0-9]{1,3}\))|\b(don't|do)\b`) // Regex pattern to match mul pairs, do and don't

	for line.Scan() {
		line := line.Text()
		matches := re.FindAllStringSubmatch(line, -1) // Find all pattern matches, store as list of strings

		for _, match := range matches { // Loop through the matches
			matchGroup := match[0]
			cleanString := strings.ReplaceAll(matchGroup, "mul", "")
			cleanString = strings.ReplaceAll(cleanString, "(", "") // Cleaning each match
			cleanString = strings.ReplaceAll(cleanString, ")", "")
			parts := strings.Split(cleanString, ",") // Spltting number pairs
			if len(parts) > 1 && action == "do" {    // Checking if the part is an action or a pair of numbers and checking current action
				firstInt, _ := strconv.Atoi(parts[0])
				secondInt, _ := strconv.Atoi(parts[1]) // Storing the pairs individually
				ans := firstInt * secondInt
				total += ans // Multiplying the pairs and adding to the total
			}
			if len(parts) == 1 { // Checking the part is an action and setting it if it is
				action = parts[0]
			}
		}
	}
	fmt.Println("Total: ", total)
}
