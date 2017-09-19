package main

import (
	"io"
	"image"
	"fmt"
	"os"
	"image/jpeg"
	_ "image/png"
	"bufio"
	"log"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("usage: input an path for converting image")
		return
	}
	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("open file %s\n", err)
	}

	err = toJPEG(bufio.NewReader(file), os.Stdout)
	if err != nil {
		fmt.Printf("to jpeg %s\n", err)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "kind = %s\n", kind)
	return jpeg.Encode(out, img, &jpeg.Options{95})
}
