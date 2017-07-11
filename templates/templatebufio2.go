package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	//f, _ := os.Open("..\data.txt")
	//reader := bufio.NewReader(f)
	l, _ := reader.ReadString(byte('\n')) // seperator newline
	num, _ := strconv.Atoi(strings.TrimSpace(l))
	//fmt.Fscanln(stdin, &num)
	for i := 0; i < num; i++ {
		l, _ = reader.ReadString(byte('\n'))
		str := strings.Fields(l)
	}
}
