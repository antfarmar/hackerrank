package main

import "fmt"
import "bufio"
import "os"
import "bytes"
import "strings"

func scanForInput() chan string {
	lines := make(chan string)
	go func() {
		for {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Split(bufio.ScanLines)
			for scanner.Scan() {
				lines <- scanner.Text()
			}
		}
	}()
	return lines
}

const alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func uniqueString(keyword string) []byte {
	uniqueKeyword := make([]byte, 0, 1)
	for i := range keyword {
		b := []byte{byte(keyword[i])}
		if !bytes.Contains(uniqueKeyword, b) {
			uniqueKeyword = append(uniqueKeyword, keyword[i])
		}
	}
	//fmt.Println(string(uniqueKeyword))
	return uniqueKeyword
}

func printMatrix(taco [][]byte) {
	for i := range taco {
		fmt.Println(string(taco[i]))
	}
}

func sortByteMatrix(toSortPtr *[][]byte) [][]byte {
	toSort := *toSortPtr
	if len(toSort) == 0 {
		return toSort
	}

	pivot := toSort[0]
	//fmt.Printf("Pivot:")
	//fmt.Println(string(pivot))
	left := make([][]byte, 0)
	right := make([][]byte, 0)
	for i := 1; i < len(toSort); i++ {
		if toSort[i][0] < pivot[0] {
			left = append(left, toSort[i])
		} else {
			right = append(right, toSort[i])
		}
	}
	left = sortByteMatrix(&left)
	//printMatrix(left)
	right = sortByteMatrix(&right)
	//printMatrix(right)
	sorted := make([][]byte, 0)
	sorted = append(sorted, left...)
	sorted = append(sorted, pivot)
	sorted = append(sorted, right...)
	return sorted
}

func makeDecoder(keyword []byte) []byte {
	decoder := make([][]byte, len(keyword))
	for i := 0; i < len(decoder); i++ {
		decoder[i] = make([]byte, 1)
	}
	for i := 0; i < len(decoder); i++ {
		decoder[i][0] = keyword[i]
	}
	filledIn := len(keyword)
	for i := range alphabet {
		b := []byte{byte(alphabet[i])}
		if !bytes.Contains(keyword, b) {
			idx := filledIn % len(keyword)
			decoder[idx] = append(decoder[idx], byte(alphabet[i]))
			filledIn++
		}
	}
	decoder = sortByteMatrix(&decoder)
	decoderString := make([]byte, 0)
	for i := range decoder {
		//fmt.Println(string(decoder[i]))
		decoderString = append(decoderString, decoder[i]...)
	}

	return decoderString
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var tests int
	fmt.Scanf("%d", &tests)
	inputs := make([]string, 2*tests)
	lines := scanForInput()
	for i := 0; i < 2*tests; i++ {
		input := <-lines
		inputs[i] = input
		//fmt.Println(inputs[i])
	}

	for i := 0; i < tests; i++ {
		keyword := inputs[i*2]
		encoded := inputs[i*2+1]
		uniqKey := uniqueString(keyword)
		decoder := makeDecoder(uniqKey)
		decoded := make([]byte, 0)
		for i := range encoded {
			next := strings.Index(string(decoder), string(encoded[i]))
			if next != -1 {
				decoded = append(decoded, alphabet[next])
			} else {
				space := " "
				decoded = append(decoded, space[0])
			}
		}
		fmt.Println(string(decoded))
	}

}
