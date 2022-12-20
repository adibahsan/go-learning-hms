package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Printf("Hello maps\n")

	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulations)

	statePopulationsAgain := make(map[string]int)
	statePopulationsAgain = map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulationsAgain)
	fmt.Println(statePopulationsAgain["Ohio"])

	pop, ok := statePopulations["Nevada"]

	fmt.Println(pop, ok)

	str := "Lorem Ipsum is simply dummy text of the printing and typesetting industry Lorem Ipsum has been the industry's standard dummy text ever since the 1500s when an unknown printer took a galley of type and scrambled it"

	arrOfStr := strings.Split(str, " ")
	fmt.Println(arrOfStr, len(arrOfStr))
	fmt.Printf("%T  %T", arrOfStr, str)

	stringCounter := make(map[string]int)

	for i := 0; i < len(arrOfStr); i++ {
		fmt.Println(arrOfStr[i])
		stringCounter[arrOfStr[i]]++
	}

	// Sorting by the keys

	arrOfStrKeys := make([]string, 0, len(stringCounter))

	for k := range stringCounter {
		// println(k, arrOfStrKeys)
		arrOfStrKeys = append(arrOfStrKeys, k)

	}

	sort.Strings(arrOfStrKeys)

	fmt.Println(arrOfStr)

	for _, k := range arrOfStrKeys {
		fmt.Println(k, stringCounter[k])
	}
}
