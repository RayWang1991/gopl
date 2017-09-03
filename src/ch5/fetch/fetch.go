package fetch

import (
	"net/http"
	"path"
	"os"
	"io"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from copy, if any
	// (For NFS like file systems report write errors till the file is closed)
	/*
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	*/
	// ex5_18
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	return local, n, err
}