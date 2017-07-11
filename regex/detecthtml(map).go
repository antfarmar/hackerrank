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
	table  = make(map[string]bool) //make(map[string]map[string]string) //table[tag][att] = int(tag[0]att[0]) (sort)
)

func main() {
	lines := scanIntoSlice(01, 0)[1:]
	//sprint(lines)
	for _, ln := range lines {
		tags := tagRE.FindAllStringSubmatch(ln, -1)
		//fmt.Printf("SUBM: %s\n\n", tags)
		for _, t := range tags { // t[1] == tag (submatch1)
			//ssprint(t)
			atts := attRE.FindAllStringSubmatch(t[2], -1) // t[2] == atts (submatch2)
			if atts == nil {
				table[t[1]+"+"] = true //add(table, t[1], "")
			} else {
				for _, a := range atts { // misses ""
					//table[t[1]] = append(table[t[1]], a[1])
					table[t[1]+"+"+a[1]] = true //add(table, t[1], a[1])
				}
			}
		}
		//fmt.Printf("SUBM: %s\n\n", attRE.FindAllStringSubmatch(ln, -1))
	}
	//fmt.Printf("%v\n", table)
	keys := []string{}
	for k, _ := range table {
		keys = append(keys, k)
	}
	sort.Sort(sort.StringSlice(keys))
	//fmt.Println(keys)
	ans, pt := "", ""
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
	fmt.Println(ans[1:])
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
