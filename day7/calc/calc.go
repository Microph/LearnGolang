package calc

import (
	"fmt"
	"net/http"
	"strconv"
)

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b float64) float64 {
	return a / b
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	a, _ := strconv.Atoi(r.FormValue("a"))
	b, _ := strconv.Atoi(r.FormValue("b"))
	fmt.Fprintf(w, `{"result":%d}`, a+b)
}

func SubHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	a, _ := strconv.Atoi(r.FormValue("a"))
	b, _ := strconv.Atoi(r.FormValue("b"))
	fmt.Fprintf(w, `{"result":%d}`, a-b)
}

/*
func SetAddRouter() {
	http.HandleFunc("/add/", AddHandler)
}

func SetSubRouter() {
	http.HandleFunc("/sub/", AddHandler)
}
*/

//refactor the upper tests
func SetRouter() {
	http.HandleFunc("/add/", AddHandler)
	http.HandleFunc("/sub/", SubHandler)
}
