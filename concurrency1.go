package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)
 func querySite(url string){
	fmt.Println("HTTP Request: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	fmt.Println("Read Body: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Body length: ", len(body))
 }
 func main(){
	go querySite("https://api.kattippara.ml/api/products")
	go querySite("https://stackoverflow.com")


	// Using an anonymous inner function in a goroutine
	go func(x int) {
		fmt.Println("Number: ", x)
	}(100)

	time.Sleep(10 * time.Second)
 }