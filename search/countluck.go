package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Graph map[*Node][]*Node

type Node struct {
	row int
	col int
	frk int
	vst bool
	dad *Node
}

var (
	//fileIn  = os.Stdin
	fileOut   = os.Stdout
	fileIn, _ = os.Open(`data\countluck.txt`)
	rw        = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))
)

// MAIN ==================================================

func main() {
	defer fileOut.Close()
	defer fileIn.Close()
	defer rw.Flush()

	T := atoi(readline())
	// For each test case, parse and solve...
	for t := 0; t < T; t++ {
		// Parse...
		NM := strings.Fields(readline())
		N, M := atoi(NM[0]), atoi(NM[1])

		grid := make([][]byte, N)
		for n := 0; n < N; n++ {
			grid[n] = readbytes()
		}
		K := atoi(readline())

		// Solve...
		forkCnt := searchGrid(grid, N, M)
		if forkCnt == K {
			out("Impressed\n")
		} else {
			out("Oops!\n")
		}
	}
}

// SOLUTION ==============================================

func searchGrid(grid [][]byte, row, col int) (forkCnt int) {
	graph, start, exit := makeAdjMap(grid)
	bfs(graph, grid, start, exit)

	// Traverse shortest path from exit to start
	for n := exit; n != start; {
		n = (*n).dad
		grid[(*n).row][(*n).col] = '|'
		if (*n).frk >= 2 {
			forkCnt++
		}
	}
	show(grid)
	return
}

// Breadth-first search the graph for the exit while counting forks
// i.e. Nodes along path with edge count/degree >= 2 - visited
func bfs(graph Graph, grid [][]byte, start, exit *Node) {
	Q := []*Node{start} //queue
	(*start).vst = true

	for len(Q) > 0 {
		t := Q[0]      //dequeue
		Q = Q[1:]      //slice head off
		if t == exit { //found shortest path
			return //so exit
		}

		for _, u := range graph[t] {
			if !(*u).vst {
				grid[(*u).row][(*u).col] = 'v'
				//show(grid)
				(*t).frk++
				(*u).dad = t
				(*u).vst = true
				Q = append(Q, u)
			}
		}
	}
	return
}

func makeAdjMap(grid [][]byte) (graph Graph, start, exit *Node) {
	graph = make(Graph)
	for r, row := range grid {
		for c, col := range row {
			switch col {
			case '.':
				graph[&Node{row: r, col: c}] = []*Node{}
			case 'M':
				start = &Node{row: r, col: c}
				graph[start] = []*Node{}

			case '*':
				exit = &Node{row: r, col: c}
				graph[exit] = []*Node{}
			}
		}
	}
	for node := range graph {
		graph[node] = getAdjNodes(graph, node)
	}
	return
}

func getAdjNodes(graph Graph, node *Node) (adj []*Node) {
	row, col := (*node).row, (*node).col
	for np := range graph {
		if np != node {
			if (*np).row == row {
				if (*np).col == col-1 || (*np).col == col+1 {
					adj = append(adj, np)
				}
			}
			if (*np).col == col {
				if (*np).row == row-1 || (*np).row == row+1 {
					adj = append(adj, np)
				}
			}
		}
	}
	return
}

func show(grid [][]byte) {
	println("==== GRID =====")
	for _, row := range grid {
		fmt.Fprintf(os.Stderr, "%s\n", row)
	}
}

// IO ====================================================

func out(s string) {
	n, err := rw.WriteString(s)
	check(err)
	if n != len(s) {
		panic("full write failed")
	}
}

func readbytes() (line []byte) {
	line, err := rw.ReadBytes('\n')
	line = bytes.TrimSpace(line)
	check(err)
	return
}

func readline() (line string) {
	line, err := rw.ReadString('\n')
	line = strings.TrimSpace(line)
	check(err)
	return
}

func atoi(s string) (i int) {
	i, err := strconv.Atoi(s)
	check(err)
	return
}

func check(err error) {
	defer func() {
		if err := recover(); err == io.EOF {
			println("reached EOF:", err)
		} else if err != nil {
			panic(err)
		}
	}()

	if err != nil || err == io.EOF {
		panic(err)
	}
}
