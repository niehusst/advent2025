package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
	"regexp"
)

/*
part1

for each prob column, apply the op at the end of the col
*/


func part1(rawProbs []string) {
	probs := make([][]string, len(rawProbs))
	for i := 0; i < len(probs); i++ {
		probs[i] = regexp.MustCompile("\\s+").Split(strings.TrimSpace(rawProbs[i]), -1)
	}
	checkSum := 0

	for col := 0; col < len(probs[0]); col++ {
		op := probs[len(probs)-1][col]
		probRes := 1
		if op == "+" {
			probRes = 0
		}
		for row := 0; row < len(probs) - 1; row++ {
			val, err := strconv.Atoi(probs[row][col])
			if err != nil {
				log.Fatal(err)
			}
			if op == "+" {
				probRes += val
			} else if op == "*" {
				probRes *= val
			} else {
				log.Fatal("bad op", op)
			}
		}
		checkSum += probRes
	}

	fmt.Println("checksum:", checkSum)
}

/*
part2

whitespace matters fml
*/

func split(probs []string) [][]string {
	// find split lens; up to next non ws char - 1 (except EOL)
	seps := make([][]int, 0)
	start := 0
	s := probs[len(probs)-1]
	for start < len(s) {
		// find col end
		colEnd := start + 1
		for colEnd < len(s) && s[colEnd] == byte(' ') {
			colEnd++
		}
		end := colEnd
		if colEnd < len(s) {
			// if not EOL, remove whitespace sep between probs
			end--
		}
	
		seps = append(seps, []int{start,end})
	
		start = colEnd //+ 1
	}

	// extract col splits
	ret := make([][]string, len(probs))
	for row := 0; row < len(probs); row++ {
		prob := make([]string, 0)
		for _, pair := range seps {
			prob = append(prob, probs[row][pair[0]:pair[1]])
		}
		ret[row] = prob
	}
	
	return ret
}

func part2(rawProbs []string) {
	probs := split(rawProbs)
	checkSum := 0

	for col := 0; col < len(probs[0]); col++ {
		op := strings.TrimSpace(probs[len(probs)-1][col])
		probRes := 1
		if op == "+" {
			probRes = 0
		}


		lenNum := len(probs[0][col])
		for numCol := 0; numCol < lenNum; numCol++ {
			// get num
			num := make([]byte, len(probs)-1)
			for row := 0; row < len(probs) - 1; row++ {
				num[row] = probs[row][col][numCol]
			}

			// do op on num
			val, err := strconv.Atoi(strings.TrimSpace(string(num)))
			if err != nil {
				log.Fatal(err)
			}
			if op == "+" {
				probRes += val
			} else if op == "*" {
				probRes *= val
			} else {
				log.Fatal("bad op", op)
			}
		}
		checkSum += probRes
	}

	fmt.Println("checksum:", checkSum)
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := string(bytes)
	// cant trim, otherwise lose meaningful ws after last op, but need to remove \n before EOF
	rawProbs := strings.Split(data[:len(data)-1], "\n")

	//part1(strings.TrimSpace(rawProbs))
	part2(rawProbs)
}
