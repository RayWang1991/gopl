package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"bob":   12,
		"alice": 13,
	}
	sortAges(ages)
}

func sortAges(ages map[string]int) {
	names := make([]string, 0, len(ages))
	// allocate cap is len(ages), more space efficient
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("name:%s, age:%d\n", name, ages[name])
	}
}
