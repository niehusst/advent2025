package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

/*
part 1

find the 2 highest numbers in each battery bank, and glue them together (in order they were found in the bank)
sum max joltage from each bank
*/

func part1(data string) {
	combinedJoltage := 0
	banks := strings.Split(strings.TrimSpace(data), "\n")

	for _, bank := range banks {
		first := bank[0]
		second := bank[1]

		// start from 1 since each attempted battery can be assigned to `second`
		for i := 1; i < len(bank); i++ {
			if bank[i] > first {
				// cant update first when this is last battery in bank
				if i != len(bank) - 1 {
					first = bank[i]
					second = bank[i+1]
					// DONT inc i again here, since i+1 may actually be greater than current bank[i], so we need to check it again
				} else {
					second = bank[i]
				}
			} else if bank[i] > second {
				second = bank[i]
			}
		}

		joltage, err := strconv.Atoi(string(first) + string(second))
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(bank, joltage)
		combinedJoltage += joltage
	}

	fmt.Println("totoal jolts", combinedJoltage)
}

/*
part 2

12 batteries instead of 2 now

.. heap-like data structure that store n largest values, and pops them out in insertion order
^ this doesn't work because we would rather have a higher first number at the expense of including a lesser later one

largest subseqence len 12
*/


func getJoltage(bank string, numBats int) int {
	bats := make([]byte, numBats)
	start := 0

	for bat := 0; bat < len(bats); bat++ {
		maxVal := byte('0')
		idx := -1
		// find max battery in allowed range
		for i := start; i < len(bank) - (len(bats) - (1+bat)); i++ {
			if bank[i] > maxVal {
				idx = i
				maxVal = bank[i]
			}
		}

		// reset start
		start = idx + 1

		// set the max value in bats
		bats[bat] = maxVal
	}

	joltage, err := strconv.Atoi(string(bats))
	if err != nil {
		log.Fatal(err)
	}
	return joltage
}

func part2(data string) {
	combinedJoltage := 0
	banks := strings.Split(strings.TrimSpace(data), "\n")

	for _, bank := range banks {
		joltage := getJoltage(bank, 12)
		//fmt.Println(bank, joltage)
		combinedJoltage += joltage
	}

	fmt.Println("totoal jolts", combinedJoltage)
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := string(bytes)
	part2(data)
}
