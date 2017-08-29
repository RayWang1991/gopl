package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"bob":   12,
		"alice": 13,
		`r
		aw`: 11,
	}

	ages2 := map[string]int{
		"bob":   12,
		"alice": 13,
		`r
		aw`: 11,
	}
	sortAges(ages)
	fmt.Println(equal(ages, ages2))
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

func equal(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; !ok || v2 != v1 {
			return false
		}
	}
	return true
}
