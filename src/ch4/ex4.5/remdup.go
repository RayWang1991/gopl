package main

import "fmt"

func main() {
	strs := []string{"1", "1", "1", "1"}
	//strs = make([]string, 3, 5)
	fmt.Println(strs)
	strs = rmdump(strs)
	fmt.Println(strs)
}

func rmdump(strs []string) []string {
	if len(strs) == 0 {
		return strs
	}
	ind := 1
	for i, str := range strs[1:] {
		if str != strs[i] {
			strs[ind] = str
			ind ++
		}
	}
	return strs[:ind]
}
