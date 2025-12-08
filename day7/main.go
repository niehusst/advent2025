package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

/*
part1

count num times split
*/

func part1(tach []string) {
	// find start
	start := 0
	for start < len(tach) && tach[0][start] != byte('S') {
		start++
	}

	// explore, count splits
	numSplits := 0
	beams := make(map[int]bool)
	beams[start] = true
	splitter := byte('^')
	for level := 1; level < len(tach); level++ {
		// advance all beams
		newBeams := make([]int, 0)
		deadBeams := make([]int, 0)
		for beam, _ := range beams {
			if tach[level][beam] == splitter {
				if beam == 0 || beam == len(tach[0])-1 {
					log.Fatal("split ob")
				}
				numSplits++
				deadBeams = append(deadBeams, beam)
				newBeams = append(newBeams, beam-1)
				newBeams = append(newBeams, beam+1)
			}
		}


		// update beam set, removing split beams and adding new beams
		for _, beam := range deadBeams {
			delete(beams, beam)
		}
		for _, beam := range newBeams {
			beams[beam] = true
		}

		// debug
		//for i, c := range tach[level] {
		//	if beams[i] {
		//		fmt.Printf("|")
		//	} else {
		//		fmt.Printf(string(c))
		//	}
		//}
		//fmt.Println()
	}

	fmt.Println("Num splits:", numSplits)
}

/*
part2

count total possible paths through a graph

 <
<
 <
*/

func part2(tach []string) {
	// find start
	start := 0
	for start < len(tach) && tach[0][start] != byte('S') {
		start++
	}

	// explore, count splits
	beams := make(map[int]int)
	beams[start] = 1
	splitter := byte('^')
	for level := 1; level < len(tach); level++ {
		// advance all beams
		newBeams := make(map[int]int)
		deadBeams := make([]int, 0)
		for beam, count := range beams {
			if tach[level][beam] == splitter {
				if beam == 0 || beam == len(tach[0])-1 {
					log.Fatal("split ob")
				}
				deadBeams = append(deadBeams, beam)
				newBeams[beam-1] += count
				newBeams[beam+1] += count
			}
		}


		// update beam set, removing split beams and adding new beams
		for _, beam := range deadBeams {
			delete(beams, beam)
		}
		for beam, count := range newBeams {
			beams[beam] += count
		}

		// debug
		//fmt.Println(beams)
		//for i, c := range tach[level] {
		//	if beams[i] > 0 {
		//		fmt.Printf("|")
		//	} else {
		//		fmt.Printf(string(c))
		//	}
		//}
		//fmt.Println()
	}

	// sum paths
	pathSum := 0
	for _, count := range beams {
		pathSum += count
	}
	fmt.Println("num paralel universes:", pathSum)
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	part2(data)
}
