package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

/*
Part 1

start at pos 50
number of times the dial is left pointing at 0 after any rotation in the sequence.
*/

const DIAL_SIZE = 100

func part1(data string) {
	zeros := 0
	dial := 50
	lines := strings.Split(data, "\n")

	for _, line := range lines {
		if len(line) < 2 {
			continue
		}
		val, err := strconv.Atoi(line[1:])
		val = val % DIAL_SIZE
		if err != nil {
			log.Fatal(err)
		}
		if line[0] == 'R' {
			// inc
			dial = (dial + val) % DIAL_SIZE
		} else {
			// dec
			dial = dial - val
			if dial < 0 {
				dial += DIAL_SIZE
			}
		}

		if dial == 0 {
			zeros += 1
		}
		//fmt.Println(dial)
	}
	fmt.Println("zeros", zeros)
}

/*
Part 2

any time 0 is passed
*/
func part2(data string) {
	zeros := 0
	dial := 50
	lines := strings.Split(data, "\n")

	for _, line := range lines {
		if len(line) < 2 {
			continue
		}
		val, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		zeros += int(val / DIAL_SIZE)
		val = val % DIAL_SIZE
		var new_dial int

		if line[0] == 'R' {
			// inc
			new_dial = (dial + val) % DIAL_SIZE
			if new_dial < dial {
				zeros += 1
			}
		} else {
			// dec
			new_dial = dial - val
			if new_dial < 0 {
				if dial != 0 {
					zeros += 1
				}
				new_dial += DIAL_SIZE
			}
			if new_dial == 0 && dial != 0 {
				zeros += 1
			}
		}
		dial = new_dial

		//fmt.Println(line, dial, zeros)
	}
	fmt.Println("zeros", zeros)
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := string(bytes)
	part2(data)
}
