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

func (db database) list(w http.ResponseWriter,req *http.Request){
	for item, price := range db {
		fmt.Fprintf(w, "%s:%s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter,req *http.Request){
	item := req.URL.Query().Get("item")
	if p, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", p)
	} else {
		fmt.Fprint(w, "no such item\n")
	}
}

func main() {
	db := database{
		"shoes":  50.0,
		"socks": 15.0,
	}
	mux := http.NewServeMux()
	fmt.Printf("%T\n",db.list)
	fmt.Printf("%T\n",db.price)
	fmt.Printf("%T\n",http.HandlerFunc(db.list))
	fmt.Printf("%T\n",http.HandlerFunc(db.price))
	mux.Handle("/list",http.HandlerFunc(db.list))
	mux.Handle("/price",http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

