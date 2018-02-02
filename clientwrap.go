package Go_Runescape

//Author: Edwin Heerschap
//Defines interfaces, structs and methods used in the package.

import (
	"net/http"
	"errors"
	"io"
	"bytes"
	)

//IHttpUtil is used for all major function calls in the package.
//An implementing struct should define the Get() method to retrieve the data from the passed url and
//store the information as a *http.Response. See the github repository for more information.
type IHttpClient interface {
	Get(url string) (*http.Response, error)
}

/*
	ALL Structs and methods from this point downwards are used for testing.
 */

//Is used for testing a failed get request.
//The Get() function returns an error.
type failGetHttpClient struct {
	IHttpClient
}

//Get returns an error containing "TEST ERROR"
func (nnhc failGetHttpClient) Get(url string) (*http.Response, error) {
	err := errors.New("TEST ERROR")
	return &http.Response{}, err
}

//invalidJsonHttpClient is used to test the function responses to invalid json.
type invalidJsonHttpClient struct {
	IHttpClient
}

//invalidJsonIOReader is used by the invalidJsonHttpClient to set the
//body of the response from an http request to something invalid.
type invalidJsonIOReader struct {
	io.ReadCloser
}

//read returns invalid json in bytes. In fact, its not similar to json at all.
func (iJson invalidJsonIOReader) Read(b []byte) (int, error){

	bs := bytes.NewBufferString("Invalid Json Bytes")
	bs.Read(b)

	return bs.Len(), io.EOF
}

//close is defined to satisfy the io.ReadCloser interface.
func (iJson invalidJsonIOReader) Close() error{
	//Doing absolutely nothing (:
	return nil
}

//Get returns a *http.Response with invalid json and a nil error.
//Used for testing functions that are trying to parse json from a Get request.
func (ijhc invalidJsonHttpClient) Get(url string) (*http.Response, error) {

	resp := http.Response{}
	ij := invalidJsonIOReader{}
	resp.Body = ij

	return &resp, nil
}
