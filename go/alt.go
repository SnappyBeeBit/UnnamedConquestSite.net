package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8001", http.FileServer(http.Dir("./static"))))
}
