package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	total   int
	matches []string
)
var allMatches [][]string

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile(`(mul\([0-9]{1,3}\,[0-9]{1,3}\))`)
	line := bufio.NewScanner(file)

	for line.Scan() {
		line := line.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			matchGroup := match[0]
			cleanString := strings.ReplaceAll(matchGroup, "mul", "")
			cleanString = strings.ReplaceAll(cleanString, "(", "")
			cleanString = strings.ReplaceAll(cleanString, ")", "")
			parts := strings.Split(cleanString, ",")
			firstInt, _ := strconv.Atoi(parts[0])
			secondInt, _ := strconv.Atoi(parts[1])
			ans := firstInt * secondInt
			total += ans
		}
	}
}
