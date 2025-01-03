package main

import (
	"net/http"
	"fmt"
	"html/template"
	"os"
	"github.com/gorilla/mux"
	"strings"
)


func main() {
	r := mux.NewRouter()
	fmt.Println("Server started on port 8001")

	tmpl := template.Must(template.ParseFiles("./static/html/main_template.html"))
	http.Handle("/static", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/static/", serveFiles)
	r.HandleFunc("/{page}", func(w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
		page := vars["page"]

		if strings.HasPrefix(page, "/static/css") {
			serveFiles(w,r)
			return
		}

		if page == "" {
			page = "index"
		}

		path := "./static/html/"
		fullpath := path + page + ".html"
		fmt.Println(fullpath)
		content, err := os.ReadFile(fullpath)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(string(content))
		tmpl.Execute(w, template.HTML(string(content)))

	} )
	http.ListenAndServe(":8001", r)
}
func serveFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	http.ServeFile(w, r, p)
}
