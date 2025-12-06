package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
	"sort"
)

type Range struct {
	start int
	end int
}

/*
part1
count ids to see how many in fresh ranges
*/
func part1(freshRanges []Range, ids []int) {
	// naive n*m approach
	fresh := 0
	for i := 0; i < len(ids); i++ {
		for j := 0; j < len(freshRanges); j++ {
			if ids[i] >= freshRanges[j].start && ids[i] <= freshRanges[j].end {
				fresh++
				break
			}
		}
	}

	fmt.Println("total fresh ids:", fresh)
}


/*
part 2

merge ranges. count range lens
*/
func part2(ranges []Range) {
	sort.Slice(ranges, func (i int, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	merged := make([]Range, 1)
	merged[0] = ranges[0]
	for i, ran := range ranges {
		if i == 0 || ran.end < ran.start {
			continue
		}
		last := len(merged)-1
		if ran.start <= merged[last].end {
			if ran.end > merged[last].end {
				merged[last].end = ran.end
			}
		} else {
			merged = append(merged, ran)
		}
	}
//	fmt.Println(merged)

	totalFresh := 0
	for _, ran := range merged {
		totalFresh += (ran.end - ran.start) + 1
	}
	fmt.Println("total fresh", totalFresh)
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := string(bytes)
	
	inputs := strings.Split(strings.TrimSpace(data), "\n\n")
	if len(inputs) != 2 {
		log.Fatal("Split input doesnt have 2 parts. Has", len(inputs))
	}
	ranges := strings.Split(inputs[0], "\n")
	ids := strings.Split(inputs[1], "\n")

	freshRanges := make([]Range, len(ranges))
	for i, ran := range ranges {
		ends := strings.Split(strings.TrimSpace(ran), "-")
		start, err := strconv.Atoi(ends[0])
		if err != nil { log.Fatal(err) }
		end, err := strconv.Atoi(ends[1])
		if err != nil { log.Fatal(err) }
		freshRanges[i] = Range{start:start, end:end}
	}
	testIds := make([]int, len(ids))
	for i, id := range ids {
		intId, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
		}
		testIds[i] = intId
	}

	//part1(freshRanges, testIds)
	part2(freshRanges)
}
