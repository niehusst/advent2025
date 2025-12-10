package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
	"math"
)

/*
part1
*/

func area(p1 [2]int, p2 [2]int) int {
	width := math.Abs(float64(p1[0] - p2[0])) + 1
	length := math.Abs(float64(p1[1] - p2[1])) + 1
	return int(width * length)
}

func part1(coords [][2]int) {
	maxArea := 0
	for i := 0; i < len(coords); i++ {
		for j := i+1; j < len(coords); j++ {
			rectArea := area(coords[i], coords[j])
			if rectArea > maxArea {
				maxArea = rectArea
			}
		}
	}
	fmt.Println("max area", maxArea)
}

/*
part2

coords are coonected. find largest rect (w/ corners) inside constrained area
just checking all corners in bounds isnt enough.. there could be a little cut out

either:
1. find way to check quickly if arb point is green/red. then check all square in rects big -> smallest
2. create a bunch of new points for the green tiles? would still have to fill in the shape
*/

func part2(coords [][2]int) {

}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rawCoords := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	coords := make([][2]int, len(rawCoords))
	for i := 0; i < len(rawCoords); i++ {
		point := strings.Split(rawCoords[i], ",")
		x, err := strconv.Atoi(point[0])
		if err != nil { log.Fatal(err) }
		y, err := strconv.Atoi(point[1])
		if err != nil { log.Fatal(err) }
		coords[i] = [2]int{x,y}
	}

	part1(coords)
}
