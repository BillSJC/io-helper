package main

import (
	"fmt"
	"io-helper"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover")
			fmt.Println(err)
		}
	}()
	//first time get Data
	data, err := io_helper.GetRequestPayload(r)
	if err != nil {
		panic(err)
	}
	//print
	fmt.Printf("First time data : %s \n", data)

	//second time get Data
	data2, err := io_helper.GetRequestPayload(r)
	if err != nil {
		panic(err)
	}
	//print
	fmt.Printf("First time data : %s \n", data2)

	//call foo()
	foo(r)

}

func foo(r *http.Request) {

	//second time get Data
	data3, err := io_helper.GetRequestPayload(r)
	if err != nil {
		panic(err)
	}
	//print
	fmt.Printf("Third time data : %s \n", data3)

}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe("127.0.0.0:8888", nil)
}
