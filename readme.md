# io-helper

An IO util to help use to reuse `io.Reader`

[![Go Report Card](https://goreportcard.com/badge/github.com/BillSJC/io-helper)](https://goreportcard.com/report/github.com/BillSJC/io-helper)
[![Open Source Helpers](https://www.codetriage.com/billsjc/io-helper/badges/users.svg)](https://www.codetriage.com/billsjc/io-helper)

## Install

```bash
go get -u github.com/BillSJC/io-helper
```

## Sample

By using func `io_helper.GetRequestPayload(r)`, u can read `Request.Body` without remove it

```go

package main

import (
	"fmt"
	io_helper "io-helper"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover() ; err != nil {
			fmt.Println("recover")
			fmt.Println(err)
		}
	}()
	//first time get Data
	data,err := io_helper.GetRequestPayload(r)
	if err != nil {
		panic(err)
	}
	//print
	fmt.Printf("First time data : %s \n",data)

	//second time get Data
	data2,err := io_helper.GetRequestPayload(r)
	if err != nil {
		panic(err)
	}
	//print
	fmt.Printf("First time data : %s \n",data2)

	//call foo()
	foo(r)

}

func foo(r *http.Request){

	//second time get Data
	data3,err := io_helper.GetRequestPayload(r)
	if err != nil {
		panic(err)
	}
	//print
	fmt.Printf("Third time data : %s \n",data3)

}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe("127.0.0.0:8888", nil)
}
```
