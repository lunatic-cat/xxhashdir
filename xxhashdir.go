package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/cespare/xxhash"
)

func hashFunc(data []byte) uint64 {
	return xxhash.Sum64(data)
}

func produce(root string, in chan string) {
	defer close(in)
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		fi, err := os.Stat(path)
		if (err == nil) && (fi.Mode().IsRegular()) {
			in <- path
		}
		return nil
	})
}

func consume(in chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for path := range in {
		dat, err := ioutil.ReadFile(path)
		if err == nil {
			fmt.Printf("%-21d %s\n", hashFunc(dat), path)
		}
	}
}

func main() {
	root := os.Args[1]
	in := make(chan string)
	wg := &sync.WaitGroup{}
	go produce(root, in)

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go consume(in, wg)
	}
	wg.Wait()
}
