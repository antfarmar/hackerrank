package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

// SOLUTION ==============================================

func cryptanalysis(text, dict []string) (soln []string) {
	dict = []string{"test", "case"}
	text = []string{"uftu", "dbtf"}

	key := []byte("abcdefghijklmnopqrstuvwxyz")
	scoreMap := makeMap(dict)
	wordMatch := 0

	for wordMatch < len(text) {
		wordMatch = 0
		shuffle(key)
		//fmt.Printf("%s\n", key)
		for _, cipher := range text {
			plain := subs(key, []byte(cipher))
			_, isInDict := scoreMap[plain]
			if isInDict {
				wordMatch++
				//fmt.Printf("%s%v\n", key, wordMatch)
			} else {
				break
			}
		}
	}
	return dict
}

func shuffle(b []byte) {
	for i, n := range rand.Perm(len(b)) {
		b[i] = byte(n) + 'a'
	}
}

func makeMap(ss []string) map[string]int {
	out := make(map[string]int, len(ss))
	for _, str := range ss {
		out[str] = len(str)
	}
	return out
}

func subs(abc []byte, txt []byte) string {
	for i, b := range txt {
		txt[i] = abc[b-'a']
	}
	return string(txt)
}

// MAIN ==================================================

func main() {
	t0 := time.Now()
	rand.Seed(time.Now().Unix())

	dictFile, err := os.Open(`dictionary.lst`)
	check(err)
	defer dictFile.Close()

	//inputFile := os.Stdin
	inputFile, err := os.Open(`cipherwords.lst`)
	check(err)
	defer inputFile.Close()

	dict := read(dictFile)
	input := read(inputFile)

	write(os.Stdout, cryptanalysis(input, dict))

	t1 := time.Now()
	debug(fmt.Sprintf("TIME: %v", t1.Sub(t0)))
}

// IO UTILITIES ==========================================

func read(file io.Reader) (input []string) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) // default is ScanLine

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	check(scanner.Err())
	return
}

func write(file io.Reader, data []string) {
	writer := bufio.NewWriter(os.Stdout)
	for _, s := range data {
		writer.WriteString(s + " ")
	}
	writer.WriteString("\n")
	writer.Flush()
}

func check(err error) {
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args)
}
