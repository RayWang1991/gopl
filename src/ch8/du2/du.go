package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
	"flag"
	"time"
)

// du2 receives a list of dir names
// walk through the directory tree recursively
// reports the total file numbers
// and the total file sizes
// if flag -v is opened, show the progress

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	if *verbose {
		inVerbose(roots)
	} else {
		normal(roots)
	}
}

func inVerbose(roots []string) {
	tick := time.NewTicker(500 * time.Millisecond)
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
loop:
	for {
		select {
		case <-tick.C:
			fmt.Printf("\r%d/%.1fMB", fileNums, float64(total)/1e6)
		case size, ok := <-fileSize:
			if !ok {
				// fileSize is closed
				tick.Stop()
				break loop
			} else {
				fileNums ++
				total += size
			}
		}
	}
	fmt.Printf("%d files, total sizes are %.1f MB\n", fileNums, float64(total)/1e6)
}

func normal(roots []string) {
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
