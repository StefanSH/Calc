package main

import (
	"fmt"
	"go/token"
	"go/types"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("expr")
	if len(query) == 0 {
		w.Write([]byte("nil expression"))
		return
	}
	expr := url.QueryEscape(query)
	expr = strings.ReplaceAll(expr, "%28", "(")
	expr = strings.ReplaceAll(expr, "%29", ")")
	expr = strings.ReplaceAll(expr, "%2A", "*")
	expr = strings.ReplaceAll(expr, "%2F", "/")
	log.Printf("get expression: %s", expr)

 	fs := token.NewFileSet()
	tv, err := types.Eval(fs, nil, 0, expr)
	if err != nil {
		log.Println(err)
		w.Write([]byte("invalid expression"))
		return
	}
	result, _ := strconv.ParseFloat(tv.Value.String(), 64)
	w.Write([]byte(fmt.Sprintf("Result of expression: %v\n", result)))
}
