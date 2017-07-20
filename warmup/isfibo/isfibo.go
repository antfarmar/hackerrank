package main

import "fmt"

//fib(50) = 12,586,269,025 > max = 10^10
var fibmap = makefibmap(50)

func main() {
	var cases, data uint64
	fmt.Scanf("%v", &cases)

	for ; cases > 0; cases-- {
		fmt.Scanf("%v", &data)
		fmt.Printf("%v\n", isfibo(data))
	}
}

func isfibo(n uint64) string {
	if fibmap[n] {
		return "IsFibo"
	}
	return "IsNotFibo"
}

func makefibmap(n int) (fibs map[uint64]bool) {
	f := fibonacci()
	fibs = make(map[uint64]bool, n)
	for len(fibs) < n {
		fibs[f()] = true
	}
	return
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() uint64 {
	x := uint64(0)
	y := uint64(1)
	return func() uint64 {
		x, y = y, x+y
		return x
	}
}

/*
Sample Input
3
2147483647
1
0

Sample Output
2147483648
4294967294
4294967295
*/
