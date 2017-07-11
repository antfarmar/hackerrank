package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	//fileIn    = os.Stdin
	fileOut   = os.Stdout
	fileIn, _ = os.Open(`data\icecream.txt`)
	rw        = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))
)

// MAIN ==================================================

func main() {
	defer fileOut.Close()
	defer fileIn.Close()
	defer rw.Flush()

	data, _ := rw.ReadString('\n')
	T, _ := strconv.Atoi(strings.TrimSpace(data))

	for t := 0; t < T; t++ {
		data, _ = rw.ReadString('\n')
		M, _ := strconv.Atoi(strings.TrimSpace(data))

		_, _ = rw.ReadString('\n') // skip N
		data, _ = rw.ReadString('\n')
		temp := strings.Fields(data)

		nums := make([]int, len(temp))
		for i, s := range temp {
			num, _ := strconv.Atoi(s)
			nums[i] = num
		}

		idx1, idx2 := subsetSum(nums, M)
		rw.WriteString(fmt.Sprintln(idx1+1, idx2+1))
	}
}

// SOLUTION ==============================================

func subsetSum(A []int, tgt int) (idx1, idx2 int) {
	S := make([]int, len(A))
	copy(S, A)
	sort.Ints(S)

	for lo, hi := 0, len(S)-1; ; {
		if S[lo]+S[hi] > tgt {
			hi--
		}
		if S[lo]+S[hi] < tgt {
			lo++
		}
		if S[lo]+S[hi] == tgt {
			idx1, idx2 = getIndex(A, S[lo], S[hi])
			break
		}

	}
	return
}

func getIndex(A []int, lo, hi int) (idx1, idx2 int) {
	for idx1 = 0; A[idx1] != lo; idx1++ {
	}
	for idx2 = 0; A[idx2] != hi || idx2 == idx1; idx2++ {
	}
	if idx1 > idx2 {
		idx1, idx2 = idx2, idx1
	}
	return
}
