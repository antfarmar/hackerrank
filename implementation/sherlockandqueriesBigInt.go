package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N, M int
	fmt.Scan(&N)
	fmt.Scan(&M)
	buf := bufio.NewReader(os.Stdin)
	line, _ := buf.ReadString('\n')
	sVals := strings.Split(line, " ")

	A := make([]*big.Int, N)
	for i := 0; i < N; i++ {
		tVal, _ := strconv.Atoi(strings.Trim(sVals[i], "\n\t "))
		A[i] = big.NewInt(int64(tVal))
	}
	line, _ = buf.ReadString('\n')
	sVals = strings.Split(line, " ")
	B := make([]int, M)
	for i := 0; i < M; i++ {
		B[i], _ = strconv.Atoi(strings.Trim(sVals[i], "\n\t "))
	}
	line, _ = buf.ReadString('\n')
	sVals = strings.Split(line, " ")

	C := make([]*big.Int, M)
	for i := 0; i < M; i++ {
		tVal, _ := strconv.Atoi(strings.Trim(sVals[i], "\n\t "))
		C[i] = big.NewInt(int64(tVal))
	}

	modulus := big.NewInt(0)
	modulus.SetString("1000000007", 10)
	moduli := make(map[int]*big.Int)

	for i := 0; i < M; i++ {
		v, ok := moduli[B[i]]
		if ok {
			moduli[B[i]] = v.Mul(v, C[i]).Mod(v, modulus)
		} else {
			moduli[B[i]] = C[i]
		}
	}

	for k, v := range moduli {
		for j := k; j <= N; j += k {
			A[j-1] = A[j-1].Mul(A[j-1], v).Mod(A[j-1], modulus)
		}
	}

	for i := 0; i < N; i++ {
		fmt.Printf("%d ", A[i])
	}

	fmt.Println()
}
