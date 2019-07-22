package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/cespare/xxhash"
)

func hashFunc(data []byte) uint64 {
	return xxhash.Sum64(data)
}

func main() {
	root := os.Args[1]
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		fi, err := os.Stat(path)
		if (err == nil) && (fi.Mode().IsRegular()) {
			dat, err := ioutil.ReadFile(path)
			if err == nil {
				fmt.Printf("%-21d %s\n", hashFunc(dat), path)
			}
		}
		return nil
	})
}
