package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	row int
	col int
	vst bool
}

var (
	//fileIn  = os.Stdin
	fileOut   = os.Stdout
	fileIn, _ = os.Open(`data\countluck.txt`)
	rw        = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))

	graph       map[*Node]bool
	start, exit *Node
)

// MAIN ==================================================

func main() {
	defer fileOut.Close()
	defer fileIn.Close()
	defer rw.Flush()

	T := atoi(readline())
	//T = 1
	for t := 0; t < T; t++ {
		NM := strings.Fields(readline())
		N, M := atoi(NM[0]), atoi(NM[1])
		grid := make([]string, N)
		for n := 0; n < N; n++ {
			grid[n] = readline()
		}
		K := atoi(readline())

		crossroadsCnt := exitForest(grid, N, M)
		println("CNT:", crossroadsCnt)
		if crossroadsCnt == K {
			out("Impressed\n")
		} else {
			out("Oops!\n")
		}
	}
}

// SOLUTION ==============================================

func exitForest(grid []string, row, col int) (waveCnt int) {
	//show(grid)
	graph = make(map[*Node]bool)
	//start, exit *Node

	// Build a graph map
	for r, rw := range grid {
		for c, co := range rw {
			switch co {
			case '.':
				graph[&Node{r, c, false}] = true
			case 'M':
				start = &Node{r, c, false}
				graph[start] = true

			case '*':
				exit = &Node{r, c, false}
				graph[exit] = true
			}
		}
	}
	//fmt.Println(graph)
	waveCnt = dfs(start, exit)
	return
}

func dfs(curr, exit *Node) (forks int) {
	(*curr).vst = true

	if curr == exit {
		fmt.Println("EXIT")
		return
	}

	cnt := 0
	for _, node := range getAdjacent(curr) {
		if !(*node).vst {
			fmt.Printf("visiting: %v\n", *node)
			cnt++
			forks += dfs(node, exit)
			//fmt.Println(node, cnt)
			if cnt > 1 {
				forks++
			}
		}
	}
	return
}

func getAdjacent(node *Node) (adj []*Node) {
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

func show(grid []string) {
	println("==== GRID =====")
	for _, row := range grid {
		fmt.Fprintln(os.Stderr, row)
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
	if err != nil {
		panic(err)
	}
}
