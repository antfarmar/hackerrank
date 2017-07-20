package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	//file        = os.Stdin
	file, _     = os.Open(`data\almostsorted.txt`)
	rawBytes, _ = ioutil.ReadAll(file)
)

func main() {
	defer file.Close()

	A := atois(strings.Fields(string(rawBytes)))[1:]
	IND := findDecreasing(A)
	fmt.Println(canSort(A, IND))
}

// Try swapping or reversing elements of A found at the
// smallest & largest indices of non-increasing elements
func canSort(A, IND []int) (ans string) {
	ans = "yes"

	switch len(IND) {
	case 0: // Already sorted
		return
	default: // Swap or reverse: len(IND) >=2 and even (pairs)
		// Get smallest and largest indices of dec elements
		beg, end := IND[0], IND[len(IND)-1]

		// Try swapping first
		A[beg], A[end] = A[end], A[beg]
		if isSorted(A) {
			ans += fmt.Sprintf("\nswap %v %v", beg+1, end+1)
			return
		}

		// Unswap, try reversing
		A[beg], A[end] = A[end], A[beg]
		reverse(A, beg, end)
		if isSorted(A) {
			ans += fmt.Sprintf("\nreverse %v %v", beg+1, end+1)
			return
		}

		// Not possible at this point
		ans = "no"
	}
	return
}

// Returns beg,end index pairs of decreasing subarrays in A
func findDecreasing(A []int) (IND []int) {
	prev := A[0]
	inc := true

	for i := 1; i < len(A); i++ {
		if inc && A[i] < prev {
			IND = append(IND, i-1)
			inc = false
		} else if !inc && A[i] > prev {
			IND = append(IND, i-1)
			inc = true
		}
		prev = A[i]
	}

	if !inc {
		IND = append(IND, len(A)-1)
	}

	return
}

// Is A sorted?
func isSorted(A []int) (inc bool) {
	prev := A[0]
	inc = true
	for _, n := range A[1:] {
		if n < prev {
			inc = false
			return
		}
		prev = n
	}
	return
}

// Reverses a subarray of A
func reverse(A []int, beg, end int) {
	for i, j := beg, end; i < j; i, j = i+1, j-1 {
		A[i], A[j] = A[j], A[i]
	}
	return
}

// IO ==============================================
func atois(S []string) (I []int) {
	I = make([]int, len(S))
	for i, a := range S {
		I[i], _ = strconv.Atoi(a)
	}
	return
}
