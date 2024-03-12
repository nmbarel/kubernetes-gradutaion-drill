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

func writeGibberish(w http.ResponseWriter, req *http.Request) {
	gibberish := req.URL.Query().Get("gibberish")
	if gibberish != "" {

		targetPath := "/tmp/gibberish"

		// Check if the path exists
		if _, err := os.Stat(targetPath); os.IsNotExist(err) {
			// Path does not exist, create it
			err := os.MkdirAll(targetPath, os.ModePerm)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				panic(err)
			}
			fmt.Println("Directory created successfully.")
		} else if err != nil {
			// Some other error occurred while checking the path
			fmt.Println("Error checking directory:", err)
			panic(err)
		} else {
			// Path already exists
			fmt.Println("Directory already exists.")
		}
		
		err := os.WriteFile("/tmp/gibberish/gib", []byte(gibberish), 0644)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, "%v", gibberish)

	}
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", mainpage)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/write", writeGibberish)

	http.ListenAndServe(":"+port, nil)
}
