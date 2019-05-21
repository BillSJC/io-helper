package io_helper

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

// Helper main struct of helper=
type Helper struct {
	MaxByteLength uint
}

// GetBytes Get data in reader without throw the buffer bu create a new Reader
func GetBytes(r *io.Reader) (data []byte, err error) {
	b, err := ioutil.ReadAll(*r)
	if err != nil {
		return nil, err
	}
	r2 := bytes.NewReader(b)
	*r = r2
	return b, nil
}

// GetBytesCloser Get data in readCloser without throw the buffer bu create a new readCloser
func GetBytesCloser(r *io.ReadCloser) (data []byte, err error) {
	b, err := ioutil.ReadAll(*r)
	if err != nil {
		return nil, err
	}
	r2 := ioutil.NopCloser(bytes.NewReader(b))
	*r = r2
	return b, nil
}

// GetRequestPayload Get data in request payload
func GetRequestPayload(r *http.Request) (data []byte, err error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	r2 := ioutil.NopCloser(bytes.NewReader(b))
	r.Body = r2
	return b, nil
}
