package main

import (
	"fmt"
	"net/http"
	"os"
)

func mainpage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%v", "OK")

	fmt.Println("Main page accessed!")
}
func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", mainpage)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":"+port, nil)
}
