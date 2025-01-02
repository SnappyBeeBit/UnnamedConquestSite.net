package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", serveFiles)
	http.ListenAndServe(":8001", nil)
}
func serveFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	if p == "./" {
		p = "./static/index.html"
	}
	http.ServeFile(w, r, p)
}
