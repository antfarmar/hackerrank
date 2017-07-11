package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

var (
	tag_re = `<([[:alnum:]]+)(.*?)>` // don't be greedy!
	att_re = `\s+([[:alnum:]]+)=`
	attRE  = regexp.MustCompile(att_re)
	tagRE  = regexp.MustCompile(tag_re)
	table  = make(map[string]string)
)

func main() {
	lines := scanIntoSlice(01, 0)[1:] // skip # cases

	for _, ln := range lines {
		tagMatches := tagRE.FindAllStringSubmatch(ln, -1) // [[match0 group1=tag group2=body=att]...]
		debug(fmt.Sprintf("TAG SUBMATCHES: %s", tagMatches))

		for _, tm := range tagMatches {
			tag, body := tm[1], tm[2]
			if body == "" { // tag has no attributes
				table[tag+"+"] = ""
				continue
			}

			attMatches := attRE.FindAllStringSubmatch(body, -1)
			for _, am := range attMatches {
				att := am[1]
				table[tag+"+"+att] = "" // don't need values
			}
		}
		debug(fmt.Sprintf("ATT SUBMATCHES: %s", attRE.FindAllStringSubmatch(ln, -1)))
	}
	debug(fmt.Sprintln("\nTABLE: ", table))

	// Sort tag+att key pair
	keys := []string{}
	for k, _ := range table {
		keys = append(keys, k)
	}
	sort.Sort(sort.StringSlice(keys)) // only one sort needed

	// Print answer
	fmt.Println(keysToString(keys))
}

func keysToString(keys []string) (ans string) {
	pt := ""
	for _, k := range keys {
		t_a := strings.Split(k, "+")
		t, a := t_a[0], t_a[1]
		if t != pt {
			pt = t
			ans += "\n" + t + ":"
		} else {
			ans += ","
		}
		ans += a
	}
	return ans[1:]
}

// IO =========================================================================
func scanIntoSlice(stdin, linesWords int) []string {
	var file *os.File
	if stdin == 0 {
		file = os.Stdin
	} else {
		var err error
		file, err = os.Open(`..\data.txt`)
		check(err)
		defer file.Close()
	}
	scanner := bufio.NewScanner(file)
	if linesWords == 0 {
		scanner.Split(bufio.ScanLines)
	} else {
		scanner.Split(bufio.ScanWords)
	}
	check(scanner.Err())
	input := []string{} //make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func ssprint(ss []string) {
	fmt.Printf("SLICE: %s\n", ss)
	for i, s := range ss {
		fmt.Printf("[%v]: %v\n", i, s)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}

func debug(s string) {
	fmt.Fprintln(os.Stderr, s)
}
