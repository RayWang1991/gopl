package main

import (
	"fmt"
	"net/http"
	"log"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", float32(d))
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	switch path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s:%s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		if p, ok := db[item]; ok {
			fmt.Fprintf(w, "%s\n", p)
		} else {
			fmt.Fprint(w, "no such item\n")
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page %s\n", req.URL)
	}
}

func main() {
	db := database{
		"shoes":  50.0,
		"socks": 15.0,
	}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
