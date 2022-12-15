package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func findItemInRucksack(rucksack string, item string) bool {
	rucksackItems := strings.Split(rucksack, "")
	for _, rucksackItem := range rucksackItems {
		if rucksackItem == item {
			return true
		}
	}
	return false
}
func findPrioritizedItem(rucksacks []string) (string, error) {
	firstRucksack := strings.Split(rucksacks[0], "")
	for _, item := range firstRucksack {
		foundInSecond := findItemInRucksack(rucksacks[1], item)
		if foundInSecond {
			foundInThird := findItemInRucksack(rucksacks[2], item)
			if foundInThird {
				return item, nil
			}
		}
	}
	return "", errors.New("prioritize item not found")
}
func findPriority(prioritizedItem string) (int, error) {
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for i, letter := range letters {
		if letter == prioritizedItem {
			return i + 1, nil
		}
	}
	return 0, errors.New("Letter was not found in alphabet")
}
func groupRucksacks(rucksacks []string) ([][]string, error) {
	var groups [][]string
	var group []string
	count := 0
	if len(rucksacks) < 1 {
		return groups, errors.New("No rucksacks provided")
	}
	for _, rucksack := range rucksacks {
		group = append(group, rucksack)
		count++
		if count == 3 {
			groups = append(groups, group)
			group = nil
			count = 0
		}
	}
	return groups, nil
}
func main() {
	dat, _ := os.ReadFile("input.txt")
	input := string(dat)
	rucksacks := strings.Split(input, "\n")
	groups, err := groupRucksacks(rucksacks)
	prioritySum := 0
	if err != nil {
		fmt.Println(err)
	}
	for _, group := range groups {
		item, err := findPrioritizedItem(group)
		if err != nil {
			fmt.Print(err)
		} else {
			priority, err := findPriority(item)
			if err != nil {
				fmt.Print(err)
			} else {
				prioritySum += priority
			}
		}
	}
	fmt.Println(prioritySum)
}
