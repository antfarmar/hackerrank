package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func decipher(key string, ctext string) (ans string) {
	abc := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//key = strip(key)
	//keyabc := strip(key + abc)
	//keyabc = genKey(keyabc, abc, len(key))
	//ans = subst(keyabc, ctext)
	ans = subst(genKey(strip(key+abc), abc, len(strip(key))), ctext)
	//debug(key, keyabc)
	return ans
}

// HELPERS ==============================================

func subst(key string, ctext string) string {
	pt := make([]rune, len(ctext))
	for i, chr := range ctext {
		if chr == ' ' {
			continue
		}
		for j, ltr := range key {
			if chr == ltr {
				pt[i] = 'A' + rune(j)
				break
			}
		}
	}
	return string(pt)
}

func genKey(key, abc string, n int) string {
	out := make([]rune, 0)
	for _, ltr := range abc { // Alphabet is sorted
		for i, chr := range key[:n] { // Search key for ltr
			if chr == ltr {
				for j := i; j < len(key); j += n {
					out = append(out, rune(key[j]))
				}
				break
			}
		}
	}
	return string(out)
}

func strip(s string) string {
	out := make([]rune, 0)
	set := make(map[rune]bool, len(s))
	for _, b := range []rune(s) {
		if !set[b] {
			set[b] = true
			out = append(out, b)
		}
	}
	return string(out)
}

// MAIN ==================================================

func main() {
	file, err := os.Open(`..\data.txt`)
	check(err)
	defer file.Close()

	process(solution, file)
	//process(solution, os.Stdin)
}

// IO UTILITIES ==========================================

func solution(input []string) (output []string) {
	for i := 1; i < len(input); i += 2 {
		key := input[i]
		ctext := input[i+1] //strings.Fields(input[i+1])
		output = append(output, decipher(key, ctext))
	}
	return output
}

func process(f func([]string) []string, file io.Reader) {
	var (
		input   []string
		scanner = bufio.NewScanner(file)
		writer  = bufio.NewWriter(os.Stdout)
	)
	for scanner.Scan() {
		input = append(input, scanner.Text()) // Lines
	}
	check(scanner.Err())

	for _, s := range f(input) {
		writer.WriteString(s + "\n")
	}
	writer.Flush()
}

func check(err error) {
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args) // println(...)
}
