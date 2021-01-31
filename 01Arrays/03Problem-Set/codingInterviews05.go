package main

import (
	"bytes"
	"fmt"
)

func main() {
	var s string = "We are happy."
	fmt.Println(replaceSpace(s))
}

func replaceSpace(s string) string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			buf.WriteString("%20")
		} else {
			buf.WriteByte(s[i])
		}
	}
	return buf.String()
}
