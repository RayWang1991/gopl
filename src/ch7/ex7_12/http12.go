package main

import (
	"fmt"
	"sync"
	"net/http"
	"strconv"
	"html/template"
)

var listHTML = template.Must(template.New("list").Parse(`
<html>
  <body>
    <table>
      <tr>
        <th>item</th>
        <th>price</th>
      </tr>
{{range $k,$v := .}}
      <tr>
        <td>{{$k}}</td>
        <td>{{$v}}</td>
      </tr>
{{end}}
   </table>
 </body>
</html>
`))

type dollar int

func (d dollar) String() string {
	return fmt.Sprintf("$%d", d)
}

type database map[string]dollar

type PriceDB struct {
	mu sync.Mutex
	db database
}

func (pdb *PriceDB) update(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid item value\n")
	}
	price := req.FormValue("price")
	p, err := strconv.Atoi(price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid price value\n")
	}
	pdb.mu.Lock()
	pdb.db[item] = dollar(p)
	pdb.mu.Unlock()
	fmt.Fprintf(w, "Update item = %s price = %d successfully!\n", item, p)
}

func (pdb *PriceDB) read(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid item value\n")
	}
	pdb.mu.Lock()
	p := pdb.db[item]
	pdb.mu.Unlock()
	fmt.Fprintf(w, "item = %s price = %d\n", item, p)
}

func (pdb *PriceDB) delete(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid item value\n")
	}
	pdb.mu.Lock()
	delete(pdb.db, item)
	pdb.mu.Unlock()
	fmt.Fprintf(w, "Delete item = %s successfully!\n", item)
}

func (pdb *PriceDB) list(w http.ResponseWriter, req *http.Request) {
	pdb.mu.Lock()
	if err := listHTML.Execute(w, pdb.db); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Template Err%s\n",err)
	}
	pdb.mu.Unlock()
}

func main() {
	pdb := PriceDB{sync.Mutex{}, database{"a": 12}}
	http.HandleFunc("/read", pdb.read)
	http.HandleFunc("/list", pdb.list)
	http.HandleFunc("/update", pdb.update)
	http.HandleFunc("/delete", pdb.delete)
	http.ListenAndServe("localhost:8000", nil)
}
