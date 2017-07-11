package main

import "fmt"

func main() {
	var L, R uint
	fmt.Scanf("%v%v", &L, &R)
	fmt.Printf("%v\n", maxXor(L, R))
}

func maxXor(A, B uint) (max uint) {
	for i := A; i <= B; i++ {
		for j := i; j <= B; j++ {
			if x := i ^ j; x > max {
				max = x
			}
		}
	}
	return
}

/*
Sample Input
1 10

Sample Output
15
*/
