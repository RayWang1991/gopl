package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
)

func main() {
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSize := make(chan int64)
	total := int64(0)
	go func() {
		for _, dir := range roots {
			walkDir(dir, fileSize)
		}
		close(fileSize)
	}()

	fileNums := 0
	for size := range fileSize {
		fileNums ++
		total += size
	}
	fmt.Printf("%d files, total sizes are %.1f MB\n", fileNums, float64(total)/1e6)
}

// walkDir walk through the directory tree recursively at the root `dir`
// sends the size of file when it encounters one
func walkDir(dir string, fileSize chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

// dirents read dir, returns the entries of directory dir
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1:%v\n", err)
		return nil
	}
	return entries
}
