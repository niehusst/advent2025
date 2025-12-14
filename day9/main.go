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


can check if point in area by iter 4 directions (until min/max coords) to see if hits permieter
- only need to check points in perimeter and in perp direction (e.g. top side doesnt need to check both ways horz)
*/

func coordToKey(p [2]int) string {
	return fmt.Sprintf("%d,%d", p[0], p[1])
}

func pointInbound(p [2]int, minp [2]int, maxp [2]int, edges map[string]bool) bool {
	//fmt.Printf("bounds? x: %d <= %d <= %d  y: %d <= %d <= %d\n", minp[0], p[0], maxp[0], minp[1], p[1], maxp[1])
	// right
	found := false
	for x := p[0]; x <= maxp[0]; x++ {
		if edges[coordToKey([2]int{x,p[1]})] {
			//fmt.Println("found right", x, p[1])
			found = true
			break
		}
	}
	if !found { return false }
	found = false
	// up
	for y := p[1]; y <= maxp[1]; y++ {
		if edges[coordToKey([2]int{p[0],y})] {
			//fmt.Println("found up", p[0], y)
			found = true
			break
		}
	}
	if !found { return false }
	found = false
	// left
	for x := p[0]; x >= minp[0]; x-- {
		if edges[coordToKey([2]int{x,p[1]})] {
			//fmt.Println("found left", x, p[1])
			found = true
			break
		}
	}
	if !found { return false }
	found = false
	// down
	for y := p[1]; y >= minp[1]; y-- {
		if edges[coordToKey([2]int{p[0],y})] {
			//fmt.Println("found down", p[0], y)
			found = true
			break
		}
	}
	if !found { return false }
	return true
}

func perimeterInbound(p1 [2]int, p2 [2]int, minp [2]int, maxp [2]int, edges map[string]bool) bool {
	var xmin, ymin, xmax, ymax int = 0,0,0,0
	// connect points
	if p1[0] > p2[0] {
		xmin = p2[0]
		xmax = p1[0]
	} else {
		xmin = p1[0]
		xmax = p2[0]
	}
	if p1[1] > p2[1] {
		ymin = p2[1]
		ymax = p1[1]
	} else {
		ymin = p1[1]
		ymax = p2[1]
	}

	// horz scan
	for x := xmin; x <= xmax; x++ {
		if !(pointInbound([2]int{x,ymin}, minp, maxp, edges) && pointInbound([2]int{x,ymax}, minp, maxp, edges)) {
			return false
		}
	}

	// vert scan
	for y := ymin; y <= ymax; y++ {
		if !(pointInbound([2]int{xmin,y}, minp, maxp, edges) && pointInbound([2]int{xmax,y}, minp, maxp, edges)) {
			return false
		}
	}
	return true
}


func part2(coords [][2]int) {
	// find min max coords
	// + create edge set
	minx := 0
	miny := 0
	maxx := 0
	maxy := 0
	edges := make(map[string]bool)
	for i := 0; i < len(coords); i++ {
		if coords[i][0] < minx {
			minx = coords[i][0]
		}
		if coords[i][0] > maxx {
			maxx = coords[i][0]
		}
		if coords[i][1] < miny {
			miny = coords[i][1]
		}
		if coords[i][1] > maxy {
			maxy = coords[i][1]
		}
		j := i + 1
		if i == len(coords)-1 {
			j = 0
		}
		// connect points
		var xmin, ymin, xmax, ymax int = 0,0,0,0
		if coords[i][0] > coords[j][0] {
			xmin = coords[j][0]
			xmax = coords[i][0]
		} else {
			xmin = coords[i][0]
			xmax = coords[j][0]
		}
		if coords[i][1] > coords[j][1] {
			ymin = coords[j][1]
			ymax = coords[i][1]
		} else {
			ymin = coords[i][1]
			ymax = coords[j][1]
		}

		// horz scan
		for x := xmin; x <= xmax; x++ {
			edges[coordToKey([2]int{x,ymin})] = true
			edges[coordToKey([2]int{x,ymax})] = true
		}

		// vert scan
		for y := ymin; y <= ymax; y++ {
			edges[coordToKey([2]int{xmin,y})] = true
			edges[coordToKey([2]int{xmax,y})] = true
		}
		
	}

	// iter all rects & check perim
	maxArea := 0
	minp := [2]int{minx,miny}
	maxp := [2]int{maxx,maxy}
	//fmt.Println(edges)
	//fmt.Println(pointInbound([2]int{2,2},minp,maxp,edges))
	for i := 0; i < len(coords); i++ {
		for j := i+1; j < len(coords); j++ {
			rectArea := area(coords[i], coords[j])
			if perimeterInbound(coords[i], coords[j], minp, maxp, edges) && rectArea > maxArea {
				maxArea = rectArea
			}
		}
	}

	fmt.Println("max contained area", maxArea)
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

	part2(coords)
}
