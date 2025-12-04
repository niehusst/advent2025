package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
	"sort"
	"container/heap"
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
*/


type Battery struct {
	jolts    string // main comparison value
	priority int    // The index in original battery bank
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Battery

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].jolts < pq[j].jolts
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Battery)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Battery, jolts string, priority int) {
	item.jolts = jolts
	item.priority = priority
	heap.Fix(pq, item.index)
}



func getJoltage(bank string, numBats int) int {
	pq := make(PriorityQueue, 0)

	for i := 0; i < len(bank); i++ {
		heap.Push(&pq, &Battery{jolts: string(bank[i]), priority: i})
		if len(pq) > numBats {
			heap.Pop(&pq)
		}
	}

	sort.Slice(pq, func (i, j int) bool {
		return pq[i].priority < pq[j].priority
	})

	bats := make([]string, numBats)
	for i, bat := range pq {
		bats[i] = bat.jolts
	}

	joltage, err := strconv.Atoi(strings.Join(bats,""))
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
		fmt.Println(bank, joltage)
		combinedJoltage += joltage
	}

	fmt.Println("totoal jolts", combinedJoltage)
}

func main() {
	bytes, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := string(bytes)
	part2(data)
}
