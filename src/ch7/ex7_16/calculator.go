package main

import (
	"net/http"
	"gopl/src/ch7/eval"
	"fmt"
	"io"
)

func main() {
	http.HandleFunc("/cal", calculator)
	http.ListenAndServe("localhost:8080", nil)
}

func calculator(w http.ResponseWriter, req *http.Request) {
	z := eval.Expr(req)
	z1 := w.(io.Writer)
	fmt.Printf("%T%T", z, z1)
	req.ParseForm()
	exprStr := req.Form.Get("expr")
	expr, err := eval.Parse(exprStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Expression:%s,%s\n", exprStr, err)
		return
	}
	res := expr.Eval(nil)
	fmt.Fprintf(w, "%s = %.6g\n", exprStr, res)
}
