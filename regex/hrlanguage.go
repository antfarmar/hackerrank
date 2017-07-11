package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var (
	langs = ":C:CPP:JAVA:PYTHON:PERL:PHP:RUBY:CSHARP:HASKELL:CLOJURE:BASH:SCALA:ERLANG:CLISP:LUA:BRAINFUCK:JAVASCRIPT:GO:D:OCAML:R:PASCAL:SBCL:DART:GROOVY:OBJECTIVEC:"
	noRE  = regexp.MustCompile(`[0-9]+`)
	lgRE  *regexp.Regexp
)

func main() {
	data := readData(0)
	for _, req := range data[1:] {
		if !noRE.MatchString(req) { // skip id#
			fmt.Printf("%vVALID\n", []string{"IN", ""}[match(req)])
		}
	}
}

func match(lg string) (ans int) {
	lgRE = regexp.MustCompile(`:` + lg + `:`)
	if lgRE.MatchString(langs) {
		ans++
	}
	return
}

func readData(file int) []string {
	b, err := []byte{}, error(nil)
	switch file {
	case 0:
		b, err = ioutil.ReadFile(`..\data.txt`)
	default:
		b, err = ioutil.ReadAll(os.Stdin)
	}
	if err != nil {
		panic(err)
	}
	return strings.Fields(string(b))
}
