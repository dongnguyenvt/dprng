package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
)

const format = `
package dprng

const seedSize = %d

// generates by crypto/rand
var seedData = [seedSize]byte{
%s
}
`

const seedSize = 16 * 1024

func main() {
	const lineSize = 8
	var buf = make([]byte, lineSize)
	var wb bytes.Buffer
	for i := 0; i < seedSize/lineSize; i++ {
		_, _ = rand.Read(buf)
		wb.WriteString("\t")
		for j := 0; j < lineSize; j++ {
			wb.WriteString(fmt.Sprintf("0x%02x,", buf[j]))
			if j != lineSize-1 {
				wb.WriteString(" ")
			}
		}
		if i != seedSize/lineSize-1 {
			wb.WriteString("\n")
		}
	}
	fmt.Printf(format, seedSize, wb.String())
}
