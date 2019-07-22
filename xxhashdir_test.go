package xxhashdir

import (
	"testing"
)

func checkOut(t *testing.T, out chan Entry) {
	got := make([]Entry, 0)
	for path := range out {
		got = append(got, path)
	}

	if len(got) < 1 {
		t.Fatalf("no entries to hash, expected one")
	}

	if len(got) > 1 {
		t.Fatalf("extra entry to hash: %v", got[1])
	}

	expectedXxhash := uint64(6467850080536788703)
	if got[0].Xxhash != expectedXxhash {
		t.Fatalf("got: %v; expected: %v", got[0].Xxhash, expectedXxhash)
	}

	expectedPath := "bin/xxhashdir.go"
	if got[0].Path != expectedPath {
		t.Fatalf("got: %v; expected: %v", got[0].Path, expectedPath)
	}
}

func TestAll(t *testing.T) {
	t.Run("Test directory", func(t *testing.T) {
		out := make(chan Entry)
		Hashdir("bin", out)
		checkOut(t, out)
	})
}
