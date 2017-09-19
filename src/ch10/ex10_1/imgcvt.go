package main

import (
	"flag"
	"log"
	"image"
	"os"
	"bufio"
	"fmt"
	"image/jpeg"
	"image/gif"
	"image/png"
)

var typeFlag = flag.String("t", "jpeg", `type for the output image, default is "jpeg""`)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("usage: input a valid image path\n")
	}
	path := args[0]
	// detect the input format of the img
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("open file %s\n", err)
	}
	input := bufio.NewReader(f)
	img, kind, err := image.Decode(input)
	if err != nil {
		log.Fatalf("decodine error %s\n", err)
	}
	fmt.Fprintf(os.Stderr, "kind = %s\n", kind)
	switch *typeFlag {
	case "jpeg", "jpg":
		jpeg.Encode(os.Stdout, img, &jpeg.Options{Quality: 100})
	case "gif":
		gif.Encode(os.Stdout, img, &gif.Options{NumColors: 256})
	case "png":
		png.Encode(os.Stdout, img)
	default:
		fmt.Fprint(os.Stderr, "unsupported kind\n")
	}
}
