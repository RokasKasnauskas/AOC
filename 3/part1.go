package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func findPrioritizedItem(rucksack string) (string, error) {
	compartmentSize := len(rucksack) / 2
	compartmentOne, compartmentTwo := rucksack[0:compartmentSize], rucksack[compartmentSize:]
	compartmentOneItems := strings.Split(compartmentOne, "")
	compartmentTwoItems := strings.Split(compartmentTwo, "")
	for i := 0; i < len(compartmentOneItems); i++ {
		for j := 0; j < len(compartmentTwoItems); j++ {
			if compartmentOneItems[i] == compartmentTwoItems[j] {
				return compartmentOneItems[i], nil
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
func main() {
	dat, _ := os.ReadFile("input.txt")
	input := string(dat)
	rucksacks := strings.Split(input, "\n")
	prioritySum := 0
	for _, rucksack := range rucksacks {
		item, err := findPrioritizedItem(rucksack)
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
