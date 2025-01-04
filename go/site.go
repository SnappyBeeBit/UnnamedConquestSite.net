package main

import (
	"net/http"
	"fmt"
	"html/template"
	"os"
	"strings"
)
func main() {
	fmt.Println("Server started on port 8001")

	tmpl := template.Must(template.ParseFiles("./static/main_template.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		page := r.URL.Path[1:]

		fmt.Println(page)
		if page == "" {
			page = "index"
		}
		if strings.HasSuffix(page, "/") {
			page += "index"
		}
		path := "./static/"
		fullpath := path + page + ".html"
		fmt.Println(fullpath)

		content, err := os.ReadFile(fullpath)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(string(content))
		tmpl.Execute(w, template.HTML(string(content)))

	} )
	http.ListenAndServe(":8001", nil)
}

