package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
	"regexp"
)

type ProductRange struct{
	first int
	last int
}

func parseInput(data string) []ProductRange {
	ranges := make([]ProductRange, 0)
	rawRanges := strings.Split(data, ",")	
	for _, prodRange := range rawRanges {
		parts := strings.Split(prodRange, "-")
		first, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil { log.Fatal(err) }
		last, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil { log.Fatal(err) }
		ranges = append(ranges, ProductRange{first: first, last: last})
	}

	return ranges
}

/*
part 1

only need to check numbers w/ even number of digits
*/
func isBad(id int) bool {
	num := strconv.Itoa(id)
	if len(num) & 1 == 1 {
		return false
	}
	var (
		ptr1 = 0
		ptr2 = len(num) / 2
	)
	for ptr2 < len(num) {
		if num[ptr1] != num[ptr2] {
			return false
		}
		ptr1++
		ptr2++
	}
	return true
}

func part1(data string) {
	// overflow risk?
	sumBadID := 0
	ranges := parseInput(data)
	for _, idRange := range ranges {
		// skip good ranges
		fNum := strconv.Itoa(idRange.first)
		lNum := strconv.Itoa(idRange.last)
		if len(fNum) == len(lNum) && len(lNum) & 1 == 1 {
			continue
		}

		// manually check
		for id := idRange.first; id <= idRange.last; id++ {
			if isBad(id) {
				//fmt.Println(id)
				sumBadID += id
			}
		}
	}
	fmt.Println("bad id sum", sumBadID)
}

/*
part2

cant skip anything anymore
*/
func isBad2(id int) bool {
	// build regex for each possible way the bad ID could be constructed.
	// unfortunately a bit slow, since most ID are not bad, in which case we do the most work
	num := strconv.Itoa(id)
	for split := 2; split <= len(num); split++ {
		if len(num) % split == 0 {
			pattern := fmt.Sprintf("(%s){%d}", num[:len(num)/split], split)
			match, err := regexp.MatchString(pattern, num)
			if err != nil {
				log.Fatal(err)
			}
			if match { return true }
		}
	}
	return false
}

func part2(data string) {
	// overflow risk?
	sumBadID := 0
	ranges := parseInput(data)
	for _, idRange := range ranges {
		// manually check
		for id := idRange.first; id <= idRange.last; id++ {
			if isBad2(id) {
				//fmt.Println(id)
				sumBadID += id
			}
		}
	}
	fmt.Println("bad id sum", sumBadID)
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := string(bytes)
	part2(data)
}
