package main

import (
	"os"
	"log"
	"fmt"
)

var cwd string

func init() {
	//var err error good
	//cwd, err = os.Getwd()
	cwd, err := os.Getwd() // cwd is a redeclared local var, this do not update the package level one
	if err != nil {
		log.Fatal("os get working dir failed:%v\n", err)
	} else {
		log.Printf("The working dir is: %v", cwd)
	}
}

func main() {
	fmt.Printf("The global wd is %s", cwd)
}
