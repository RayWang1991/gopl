package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := readFile(os.Args[1])
	if err != nil {
		fmt.Println("Failed due to %s", err)
		return
	}
	fmt.Println("Succeed!")
	fmt.Println(string(bs))
}

func readFile(fileName string) ([]byte, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer fmt.Println("file closed")
	defer f.Close()
	return ioutil.ReadAll(f)
}
