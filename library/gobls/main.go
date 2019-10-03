package main

import (
	"fmt"
	"github.com/karrick/gobls"
	"os"
)

func main() {
	var lines, characters int
	ls := gobls.NewScanner(os.Stdin)
	for ls.Scan() {
		lines++
		characters += len(ls.Bytes())
	}
	if err := ls.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "cannot scan:", err)
	}
	fmt.Println("Counted", lines, "lines and", characters, "characters.")
}
