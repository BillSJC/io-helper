package io_helper

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

/*
@Helper: main func of helper
*/
type Helper struct {
	MaxByteLength uint
}

/*
@Name:	GetBytes
@Desc:	Get data in reader without throw the buffer bu create a new Reader
@param:
	r	*oi.Reader	input data
@return:
	data	[]byte	data in reader
	err		error

*/
func GetBytes(r *io.Reader) (data []byte, err error) {
	b, err := ioutil.ReadAll(*r)
	if err != nil {
		return nil, err
	}
	r2 := bytes.NewReader(b)
	*r = r2
	return b, nil
}

/*
@Name:	GetBytesCloser
@Desc:	Get data in readCloser without throw the buffer bu create a new readCloser
@param:
	r	*oi.Reader	input data
@return:
	data	[]byte	data in reader
	err		error

*/
func GetBytesCloser(r *io.ReadCloser) (data []byte, err error) {
	b, err := ioutil.ReadAll(*r)
	if err != nil {
		return nil, err
	}
	r2 := ioutil.NopCloser(bytes.NewReader(b))
	*r = r2
	return b, nil
}

/*
@Name: 	GetRequestPayload
@Desc:	Get data in request payload
@param:
	r	*oi.Reader	input data
@return:
	data	[]byte	data in reader
	err		error

*/
func GetRequestPayload(r *http.Request) (data []byte, err error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	r2 := ioutil.NopCloser(bytes.NewReader(b))
	r.Body = r2
	return b, nil
}
