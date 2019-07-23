package xxhashdir

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/cespare/xxhash"
)

// Entry of out chan
type Entry struct {
	Path   string
	Xxhash uint64
}

func hashFile(path string) (uint64, error) {
	var (
		hash      = xxhash.New()
		file, err = os.Open(path)
	)

	if file != nil {
		defer file.Close()
		_, err = io.Copy(hash, file)
	}

	return hash.Sum64(), err
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

func consume(in chan string, out chan Entry, wg *sync.WaitGroup) {
	defer wg.Done()

	for path := range in {
		hash, err := hashFile(path)
		if err == nil {
			out <- Entry{Path: path, Xxhash: hash}
		}
	}
}

func stop(out chan Entry, wg *sync.WaitGroup) {
	wg.Wait()
	close(out)
}

// Hashdir prints all directory contents with xxhash sums
func Hashdir(root string, out chan Entry) {
	in := make(chan string)
	wg := &sync.WaitGroup{}
	go produce(root, in)

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go consume(in, out, wg)
	}
	go stop(out, wg)
}
