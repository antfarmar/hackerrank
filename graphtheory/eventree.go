package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph map[int][]int

var (
	//fileIn       = os.Stdin
	fileOut      = os.Stdout
	fileIn, _    = os.Open(`.\data\eventree.txt`)
	rw           = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))
	bytesWritten = 0
)

// MAIN ==================================================

func main() {
	defer report()
	defer fileOut.Close()
	defer fileIn.Close()
	defer rw.Flush()

	VE := atois(readfields())
	V, E := VE[0], VE[1]
	graph := make(Graph, V)

	// Build graph (adj list map)
	for i := 0; i < E; i++ {
		uv := atois(readfields())
		u, v := uv[0], uv[1]
		if u > v {
			u, v = v, u
		}
		graph[u] = append(graph[u], v)
	}

	write(fmt.Sprintln(evenTreeForest(graph)))
}

// SOLUTION ==============================================

func evenTreeForest(graph Graph) (rmCnt int) {
	for vert, _ := range graph {
		for _, adjv := range graph[vert] {
			childCnt := countChildRec(graph, adjv) + 1
			if childCnt%2 == 0 {
				rmCnt++
			}
		}
	}
	return
}

func countChildRec(graph Graph, vert int) (cnt int) {
	for _, v := range graph[vert] {
		cnt += countChildRec(graph, v) + 1
	}
	return
}

// IO ====================================================

func write(s string) {
	n, err := rw.WriteString(s)
	check(err)
	bytesWritten += n
}

func report() {
	println("<wrote", bytesWritten, "bytes>")
}

func readline() (line string) {
	line, err := rw.ReadString('\n')
	line = strings.TrimSpace(line)
	check(err)
	return
}

func readfields() (fields []string) {
	line, err := rw.ReadString('\n')
	fields = strings.Fields(line)
	check(err)
	return
}

func atois(S []string) (I []int) {
	I = make([]int, len(S))
	for i, a := range S {
		n, err := strconv.Atoi(a)
		check(err)
		I[i] = n
	}
	return
}

func check(err error) {
	if err != nil {
		if err.Error() == "EOF" {
			println("<EOF reached>")
		} else {
			panic(err)
		}
	}
}
