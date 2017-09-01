package main

import (
	"sort"
	"fmt"
	"gopl/src/ch5/ex5_11"
)

const (
	ALGO  = "algorithms"
	DATAS = "data structures"
	CALC  = "calculus"
	COMPL = "compilers"
	FORM  = "formal languages"
	COMPO = "computer organization"
	DB    = "databases"
	DM    = "decreate math"
	NET   = "networks"
	PL    = "programming languages"
	LA    = "linear algebra"
	INTRO = "intro to programming"
	OS    = "operating systems"
)

var prereps = map[string][]string{
	ALGO:  {DATAS},
	CALC:  {LA},
	COMPL: {DATAS, FORM, COMPO},
	DATAS: {DM},
	DB:    {DATAS},
	DM:    {INTRO},
	FORM:  {DM},
	NET:   {OS},
	OS:    {DATAS, COMPO},
	PL:    {DATAS, COMPO},
	LA:    {CALC},
}

func main() {
	//for _, item := range topoSort(prereps) {
	//	fmt.Printf("%s\n", item)
	//}
	//fmt.Println()
	// ex5_11
	order, err := ex5_11.TopoSort(prereps)
	if err != nil {
		if err == ex5_11.CircleError {
			fmt.Printf("%v : %v", err, order)
		} else {
			fmt.Println(err)
		}
		return
	}
	for _, item := range order {
		fmt.Printf("%s\n", item)
	}
}

func topoSort(m map[string][]string) []string {
	// m is a DAG
	// post order dfs gives the topo sort
	found := map[string]bool{}
	order := []string{}

	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !found[item] {
				found[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	keys := make([]string, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}
