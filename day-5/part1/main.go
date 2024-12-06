package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	ruleCount      int
	pages          [][]string
	goodPages      [][]string
	badPages       [][]string
	goodPageNumber []int
	badPageNumber  []int
	midPageTotal   int
	badPageTotal   int
	rules          map[string]int
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()
	raw := bufio.NewScanner(file)
	rules = make(map[string]int)

	for raw.Scan() {
		line := raw.Text()
		if strings.Contains(line, "|") {
			fmt.Println(line)
			rules[line] = 1
		}
		if strings.Contains(line, ",") {
			splitPages := strings.Split(line, ",")
			pages = append(pages, splitPages)

		}
	}
	for i := 0; i < len(pages); i++ {
		if ruleCheck(rules, pages[i]) == 1 {
			badPages = append(badPages, pages[i])
		} else {
			midPageTotal += returnMiddlePage(pages[i])
		}
	}
	for i := 0; i < len(badPages); i++ {
		sort(badPages[i])
		badPageTotal += returnMiddlePage(badPages[i])
	}
	fmt.Println("Total: ", midPageTotal)
	fmt.Println("Total: ", badPageTotal)
}

func ruleCheck(rules map[string]int, page []string) int {
	for rule := range rules {
		rule := strings.Split(rule, "|")
		if !slices.Contains(page, rule[0]) || !slices.Contains(page, rule[1]) {
			continue
		} else if slices.Index(page, rule[0]) > slices.Index(page, rule[1]) {
			return 1
		}
	}
	return -1
}

func returnMiddlePage(page []string) int {
	midPage, _ := strconv.Atoi(page[len(page)/2])
	return midPage
}

func sort(page []string) { // Thanks Kim ._.
	slices.SortStableFunc(page, func(a, b string) int {
		if rules[a+"|"+b] == 1 {
			return -1
		} else if rules[b+"|"+a] == 1 {
			return 1
		}
		return 0
	})
}
