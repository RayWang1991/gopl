package thumbnail6

import (
	"sync"
	"gopl.io/ch8/thumbnail"
	"os"
)

// make thumbnails receive file names from a channel, make thumbnails of the file,
// then return the bytes it creates
func makeThumbnails(filenames chan string) int64 {
	sizes := make(chan int64)
	var wg = sync.WaitGroup{}
	for _, file := range filenames {
		wg.Add(1)
		// worker
		go func(f string) { // attention!
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(file)
	}
	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
