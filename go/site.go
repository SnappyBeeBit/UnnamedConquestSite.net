package main

import (
	"net/http"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"errors"
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

		path := "./static/"
		fullpath := path + page + ".html"

		result_path, err := verifyPath(fullpath)

		if err != nil {
			http.Redirect(w,r,"/", http.StatusBadRequest)
		}

		fmt.Println(fullpath)
		content, err := os.ReadFile(result_path)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(string(content))
		tmpl.Execute(w, template.HTML(string(content)))

	} )
	http.ListenAndServe(":8001", nil)
}
func verifyPath(path string) (string, error) {
	c := filepath.Clean(path)
	fmt.Println("Cleaned path: " + c)

	r, err := filepath.EvalSymlinks(c)
	if err != nil {
		fmt.Println("Error " + err.Error())
		return c, errors.New("Unsafe or invalid path specified")
	} else {
		fmt.Println("Canonical: " + r)
		return r, nil
	}

}
