package main

import (
	"fmt"
	"net/http"
	"log"
)
func main(){
	http.HandleFunc("/",getRequest)
	http.HandleFunc("/hellow",helloRequest)
	http.HandleFunc("/headers",headers)
	//
	log.Println("Listening on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080",nil))
	
}
//folder printing
func getRequest(w http.ResponseWriter, r *http.Request){
	file_requested := "./" + r.URL.Path
	http.ServeFile(w, r, file_requested)
	return
}
// Basic handler for /hello requests
func helloRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
	return
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}