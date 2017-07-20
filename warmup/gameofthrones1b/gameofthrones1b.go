package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	//file = os.Stdin
	file, _ = os.Open(`.\data\gameofthrones1.txt`)
	data, _ = ioutil.ReadAll(file)
)

func main() {
	data = bytes.TrimSpace(data)

	oddCnt, ans := 0, "YES"

	for ascii := byte('a'); ascii <= byte('z'); ascii++ {

		chrCnt := 0
		for _, chr := range data {
			if chr == ascii {
				chrCnt++
			}
		}

		if chrCnt&1 == 1 {
			oddCnt++
		}

		if oddCnt > 1 {
			ans = "NO"
			break
		}
	}

	fmt.Println(ans)
}
