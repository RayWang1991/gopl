package main

import (
	"fmt"
	"net/http"
	"log"
	"strconv"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", float32(d))
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s:%s\n", item, price)
	}
}

func (db database) addOrUpdate(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Invalid item %s\n", item)
	}
	price := req.URL.Query().Get("price")
	p, err := strconv.ParseFloat(price, len(price))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Invalid price %s\n", price)
	}

	db[item] = dollars(p)
	// list the new price
	fmt.Fprintf(w, "%s %s\n", item, dollars(p))
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Invalid item %s\n", item)
	} else {
		delete(db, item)
		fmt.Fprintf(w, "Delete %s successful!\n", item)
	}

}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if p, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", p)
	} else {
		fmt.Fprint(w, "no such item\n")
	}
}

func main() {
	db := database{
		"shoes": 50.0,
		"socks": 15.0,
	}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.addOrUpdate)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
