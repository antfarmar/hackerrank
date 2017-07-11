package main

import (
	"bufio"
	"math"
	"os"
	"strings"
)

var (
	//fileIn  = os.Stdin
	fileOut   = os.Stdout
	fileIn, _ = os.Open(`data\encryption.txt`)
	rw        = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))
)

// MAIN ==================================================

func main() {
	defer fileOut.Close()
	defer fileIn.Close()
	defer rw.Flush()

	data, _ := rw.ReadString('\n')
	plainText := strings.TrimSpace(data)
	encrypt(plainText)
}

// SOLUTION ==============================================

func encrypt(pt string) {
	width, _ := minimizeArea(len(pt))
	cnt := 0

	for h := 0; cnt < len(pt); h++ {
		for i := h; i < len(pt); i += width {
			rw.WriteByte(pt[i])
			cnt++
		}
		rw.WriteByte(' ')
	}
	rw.WriteByte('\n')
}

func minimizeArea(area int) (width, height int) {
	ceil := int(math.Ceil(math.Sqrt(float64(area))))
	floor := int(math.Sqrt(float64(area)))
	width, height = ceil, floor

	for width*height < area {
		if (width+1)*height < width*(height+1) && width+1 <= ceil {
			width++
		} else {
			height++
		}
	}
	return
}
