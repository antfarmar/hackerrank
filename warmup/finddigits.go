package main

import "fmt"

func main() {
	var cases, data uint64
	fmt.Scanf("%v", &cases)

	for ; cases > 0; cases-- {
		fmt.Scanf("%v", &data)
		fmt.Printf("%v\n", find(data))
	}
}

func find(n uint64) (cnt int) {
	for _, d := range fmt.Sprint(n) {
		div := uint64(d - '0')
		if div > 0 && n%div == 0 {
			cnt++
		}
	}
	return
}
