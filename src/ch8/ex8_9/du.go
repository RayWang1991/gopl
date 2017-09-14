package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
	"flag"
	"time"
	"sync"
)

// du shows the result in separate table
// du is a concurrent version of du2
// du2 receives a list of dir names
// walk through the directory tree recursively
// reports the total file numbers
// and the total file sizes
// if flag -v is opened, show the progress
// the flag -c sets the max concurrency number, default is 20

type fileSizeInfo struct {
	id   int
	size int64
}

var maxcon = flag.Int("c", 20, "set the max concurrency number, default is 20")
var token chan struct{}
var wg sync.WaitGroup

func main() {
	starttime := time.Now()
	flag.Parse()
	token = make(chan struct{}, *maxcon)
	wg = sync.WaitGroup{}
	roots := flag.Args()
	inVerbose(roots)
	fmt.Printf("%.1fs pased\n", time.Since(starttime).Seconds())
}

func inVerbose(roots []string) {
	tick := time.NewTicker(500 * time.Millisecond)
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan *fileSizeInfo)

	for i, dir := range roots {
		wg.Add(1)
		go func(id int, dir string) {
			walkDir(dir, id, fileSizes)
		}(i, dir)
	}

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	fileNums := make([]int, len(roots))
	totals := make([]int64, len(roots))
loop:
	for {
		select {
		case <-tick.C:
			fmt.Printf("/r")
			for i, root := range roots {
				fmt.Printf("%d. %s, %d files / %.1fMB\n", i, root, fileNums[i], float64(totals[i])/1e6)
			}
		case sizeInfo, ok := <-fileSizes:
			if !ok {
				// fileSize is closed
				tick.Stop()
				break loop
			} else {
				fmt.Printf("id: %d size:%d \n", sizeInfo.id, sizeInfo.size)
				fileNums[sizeInfo.id] ++
				totals[sizeInfo.id] += sizeInfo.size
			}
		}
	}
	fmt.Printf("\rDone\n")
	for i, root := range roots {
		fmt.Printf("%d. %s, %d files / %.1fMB\n", i, root, fileNums[i], float64(totals[i])/1e6)
	}
}

// walkDir walk through the directory tree recursively at the root `dir`
// sends the size of file when it encounters one
func walkDir(dir string, id int, fileSize chan<- *fileSizeInfo) {
	token <- struct{}{}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			go walkDir(subdir, id, fileSize)
		} else {
			fileSize <- &fileSizeInfo{id, entry.Size()}
		}
	}
	<-token
	wg.Done()
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
