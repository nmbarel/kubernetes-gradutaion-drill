package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func mainpage(w http.ResponseWriter, req *http.Request) {
	backendurl := os.Getenv("BACKENDURL")
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(backendurl)
	fmt.Fprintf(w, "%v", "<script type='text/javascript'> setInterval(function() {window.location.reload();},1000);</script>")

	if err != nil {
		fmt.Fprintf(w, "%v", "<div style='width:500px;position:relative;display: block;margin: 20px auto;'><h1>Welcome to K8s Graduation drill</h1></div>")
		fmt.Fprintf(w, "%v", "<div style='animation: blinker 1s linear infinite;width:500px;text-align:center;position:relative;display: block;color: red;margin: 20px auto;'>Backend is not working</div>")
	}
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Fprintf(w, "%v", "<div style='width:500px;position:relative;display: block;margin: 20px auto;'><h1>Welcome to K8s Graduation drill</h1></div>")
	fmt.Println(string(body))
	if string(body) == "OK" {
		fmt.Fprintf(w, "%v", "<div style='width:500px;text-align:center;position:relative;display: block;color: green;margin: 20px auto;'>Backend is operational</div>")
	} else {
		fmt.Fprintf(w, "%v", "<div style='animation: blinker 1s linear infinite;width:500px;text-align:center;position:relative;display: block;color: red;margin: 20px auto;'>Backend is not working</div>")
	}
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
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":"+port, nil)
}
