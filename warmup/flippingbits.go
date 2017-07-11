package main

import "fmt"

func main() {
	var cases, u uint32
	fmt.Scanf("%v", &cases)
	for ; cases > 0; cases-- {
		fmt.Scanf("%v", &u)
		fmt.Printf("%v\n", flipBitwise(u))
	}
}

func flipBitwise(u uint32) uint32 {
	return ^u
}

func flipBitString(u uint32) uint32 {
	bits := fmt.Sprintf("%032b", u)
	//fmt.Printf("U: %032b\n", u)
	n := len(bits) - 1
	u = 0
	for i := 0; i <= n; i++ {
		if bits[i] == '0' {
			u += (1 << uint8(n-i))
		}
	}
	//fmt.Printf("U: %032b\n", u)
	return u
}

/*
Sample Input
4
4294967295
2147483647
1
0

Sample Output
0
2147483648
4294967294
4294967295
*/
