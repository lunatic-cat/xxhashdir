package main

import (
	"fmt"
	"os"

	"github.com/razum2um/xxhashdir"
)

func main() {
	out := make(chan xxhashdir.Entry)
	xxhashdir.Hashdir(os.Args[1], out)
	for entry := range out {
		fmt.Printf("%-21d %s\n", entry.Xxhash, entry.Path)
	}
}
