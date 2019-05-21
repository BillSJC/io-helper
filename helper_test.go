package io_helper

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetRequestPayload(t *testing.T) {
	testStr := "testSt"
	r := new(http.Request)
	r.Body = ioutil.NopCloser(bytes.NewReader([]byte(testStr)))

	b1, err := GetRequestPayload(r)
	if err != nil {
		panic(err)
	}
	b2, err := GetRequestPayload(r)
	if err != nil {
		panic(err)
	}
	if string(b1) == string(b2) && string(b2) == testStr {
		fmt.Println("success")
		t.Log("success")

	} else {
		fmt.Println("failed")
		t.Errorf("faild")
	}

}
