package main

import (
	"reflect"
	"fmt"
	"os"
)

// Print prints the methods set of the value x
func Print(x interface{}){
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n",t)

	for i:= 0; i< v.NumMethod(); i++{
		methType := v.Method(i).Type()
		fmt.Printf("%s | %s\n",t.Method(i).Name,methType.String())
	}
}

func main(){
	var i interface{} = os.Stdin
	Print(i)
}
