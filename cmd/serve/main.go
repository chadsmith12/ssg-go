package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	fmt.Println("Starting server on 8080")
	http.ListenAndServe(":8080", nil)
}
