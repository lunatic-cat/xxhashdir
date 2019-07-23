# xxhashdir

[![Build Status](https://travis-ci.com/razum2um/xxhashdir.svg?branch=master)](https://travis-ci.org/razum2um/xxhashdir)

## Usage

this package does fast filysystem fingerprinting using [xxHash](http://cyan4973.github.io/xxHash/)

```sh
# instead of "find . -type f -exec shasum -b {} \;"
$ ./xxhashdir .
...
880788507839261490    README.md
11541949788444589007  .travis.yml
6467850080536788703   bin/xxhashdir.go
...
```

typical CLI use:

```sh
./xxhashdir dir > before
# modify fs
./xxhashdir dir > after
diff <(sort before) <(sort after) | sort -nk3
```

## Speed

Blazingly fast, order of magnitude :) Digesting xcode-10.2 with >250K files:

| Time | Cmd |
| --- | --- |
| 656 sec | time find /Applications/Xcode.app -type f -exec xxhsum {} \; > xxhsum.txt |
| 45 sec | time ./xxhashdir /Applications/Xcode.app > xxhsumdir.txt |

## Golang api

```go
func Hashdir(root string, out chan Entry)
```

where

```go
type Entry struct {
    Path   string
    Xxhash uint64
}
```
