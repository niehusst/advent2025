package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

/*
part1
count num positinos w/ less than 4 surrounding @
*/

func part1(data string) {
	rows := strings.Split(strings.TrimSpace(data), "\n")
	accessiblePapers := 0
	paper := byte('@')

	for i, row := range rows {
		for j, item := range row {
			if item == '@' {
				// check surrounding
				n := 0
				if i > 0 && rows[i-1][j] == paper {
					n += 1
				}
				if j > 0 && rows[i][j-1] == paper {
					n += 1
				}
				if i > 0 && j > 0 && rows[i-1][j-1] == paper {
					n += 1
				}
				if i < len(rows)-1 && rows[i+1][j] == paper {
					n += 1
				}
				if j < len(row)-1 && rows[i][j+1] == paper {
					n += 1
				}
				if i < len(rows)-1 && j < len(row)-1 && rows[i+1][j+1] == paper {
					n += 1
				}
				if i > 0 && j < len(row)-1 && rows[i-1][j+1] == paper {
					n += 1
				}
				if j > 0 && i < len(rows)-1 && rows[i+1][j-1] == paper {
					n += 1
				}

				if n < 4 {
					accessiblePapers++
				}
			}
		}
	}

	fmt.Println("accessible papers:", accessiblePapers)
}

/*
part2
remove as many rolls as possible recursively
*/

/// returns number of rolls removed
func removeRolls(rows [][]byte) int {
	accessiblePapers := 0
	paper := byte('@')
	mark := byte('.')

	for i, row := range rows {
		for j, item := range row {
			if item == '@' {
				// check surrounding
				n := 0
				if i > 0 && rows[i-1][j] == paper {
					n += 1
				}
				if j > 0 && rows[i][j-1] == paper {
					n += 1
				}
				if i > 0 && j > 0 && rows[i-1][j-1] == paper {
					n += 1
				}
				if i < len(rows)-1 && rows[i+1][j] == paper {
					n += 1
				}
				if j < len(row)-1 && rows[i][j+1] == paper {
					n += 1
				}
				if i < len(rows)-1 && j < len(row)-1 && rows[i+1][j+1] == paper {
					n += 1
				}
				if i > 0 && j < len(row)-1 && rows[i-1][j+1] == paper {
					n += 1
				}
				if j > 0 && i < len(rows)-1 && rows[i+1][j-1] == paper {
					n += 1
				}

				if n < 4 {
					accessiblePapers++
					rows[i][j] = mark
				}
			}
		}
	}

	return accessiblePapers
}

func part2(data string) {
	strRows := strings.Split(strings.TrimSpace(data), "\n")
	rows := make([][]byte, len(strRows))
	// convert to 2d byte slice
	for i, row := range strRows {
		rows[i] = []byte(row)
	}
	totalRemoved := 0

	removed := 1
	for removed > 0 {
		removed = removeRolls(rows)
		totalRemoved += removed
	}

	fmt.Println("total removed:", totalRemoved)
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := string(bytes)
	part2(data)
}
