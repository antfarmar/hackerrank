package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

// SOLUTION ==============================================

func cryptanalysis(ctext, dict []string) (soln []string) {

	t0, t1 := time.Now(), time.Now() // Start timers

	var buf bytes.Buffer
	key := []byte("wztbeackxoqspvlugdrijhfynm")
	tmp := []byte("abcdefghijklmnopqrstuvwxyz")
	maxKey := []byte("abcdefghijklmnopqrstuvwxyz")

	shuffle(key)
	scoreMap := makeMap(dict) // score is word length
	wordMatch, attempts, fitness, maxfit := 0, 0, 0, 0
	tlimit := time.Millisecond * time.Duration(3900)

	// Main Loop
	// Keep improving key until it deciphers all ciphertext
	for amtWords := len(ctext); wordMatch < amtWords; attempts++ {
		t1 = time.Now()
		if t1.Sub(t0) > tlimit {
			break
		}

		// Restart to get out of local maxima
		if attempts > 1000 {
			maxfit, attempts = 0, 0
			shuffle(key)
			//debug(fmt.Sprint("RESTART"))
		}

		buf.Reset()
		fitness, wordMatch = 0, 0

		// Restore unfit child tmp key to parent and reshuffle
		_ = copy(tmp, key)
		//numSwaps := 3 - 3*wordMatch/amtWords
		numSwaps := 1 + 3*wordMatch/amtWords
		swap(tmp, numSwaps)

		// Evaluate fitness of tmp key
		for _, cword := range ctext {
			plain := subst(tmp, []byte(cword))
			score, isInDict := scoreMap[plain]
			if isInDict {
				wordMatch++
				fitness += wordMatch + score/wordMatch
				//buf.WriteString(cword + "(" + plain + ") ")
			}
		}

		// Keep the most fit key
		if fitness > maxfit {
			maxfit = fitness
			_ = copy(key, tmp)
			_ = copy(maxKey, tmp)
			attempts = 0

			//debug(fmt.Sprintf("%v/%v %s %v %v",
			//	wordMatch, amtWords, maxKey, maxfit, buf.String()))
		}
	}

	// Finally decipher ciphertext with best key for output
	for _, cword := range ctext {
		soln = append(soln, subst(maxKey, []byte(cword)))
	}

	return soln
}

// HELPER FUNCTIONS =======================================

func makeMap(ss []string) map[string]int {
	out := make(map[string]int, len(ss))
	for _, str := range ss {
		str = strings.ToLower(str)
		out[str] = len(str) // score is word length
	}
	return out
}

func subst(abc []byte, txt []byte) string {
	for i, b := range txt {
		txt[i] = abc[b-'a']
	}
	return string(txt)
}

func swap(b []byte, n int) {
	for ; n > 0; n-- {
		r1 := rand.Intn(26)
		r2 := rand.Intn(26)
		b[r1], b[r2] = b[r2], b[r1]
	}
}

func shuffle(b []byte) {
	// ~= swap(b,len(b))
	for i, n := range rand.Perm(len(b)) {
		b[i] = byte(n) + 'a'
	}
}

// MAIN ==================================================

func main() {
	//t0 := time.Now()
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

	//t1 := time.Now()
	//debug(fmt.Sprintf("TIME: %v", t1.Sub(t0)))
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
