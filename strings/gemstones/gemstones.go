package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type Set map[uint8]bool

// MAIN ==================================================

func main() {
	//input,_ := ioutil.ReadAll(os.Stdin)
	input, _ := ioutil.ReadFile(`data\gemstones.txt`)
	stones := bytes.Fields(input)[1:]
	fmt.Println(gemCount(stones))
}

// SOLUTION ==============================================

func gemCount(stones [][]uint8) (gemCnt int) {
	amtStones := len(stones)
	elemSets := make([]Set, amtStones)

	// Map the composition of each stone
	for i, stone := range stones {
		elemSets[i] = make(Set)
		for _, e := range stone {
			elemSets[i][e] = true
		}
	}

	// Count gems (elements that occur >=1 in each of the stones)
	for i := uint8('a'); i <= 'z'; i++ {
		elemCnt := 0
		for _, elemSet := range elemSets {
			if elemSet[i] {
				elemCnt++
			}
		}
		if elemCnt == amtStones {
			gemCnt++
		}
	}
	return
}
