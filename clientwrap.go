package Go_Runescape

import (
	"net/http"
	"errors"
	"io"
	"bytes"
)

//IHttpClient is made so the dev can use their own
//method of performing the http request.
type IHttpClient interface {
	Get(url string) (*http.Response, error)
}

/*
	ALL Structs and methods from this point downwards are used for testing.
 */

//NotNilHttpClient is used in the testing files to test for
//not nil errors.
type notNilHttpClient struct {
	IHttpClient
}

//Get returns an error containing "TEST ERROR"
func (nnhc notNilHttpClient) Get(url string) (*http.Response, error) {
	err := errors.New("TEST ERROR")
	return &http.Response{}, err
}

//invalidJsonHttpClient is used to test the functions when obtaining
//invalid json.
type invalidJsonHttpClient struct {
	IHttpClient
}

//invalidJsonIOReader is used by the invalidJsonHttpClient to set the
//body of the response from an http request to something invalid.
type invalidJsonIOReader struct {
	io.ReadCloser
}

//read returns invalid json in bytes. Infact its not json at all.
func (iJson invalidJsonIOReader) read(b []byte) {

	bs := bytes.NewBufferString("Invalid Json Bytes")
	bs.Read(b)

}

func (iJson invalidJsonIOReader) close() {
	//Doing absolutely nothing (:
}

//Get returns a *http.Response with invalid json and a nil error.
//Used for testing functions that are trying to parse json.
func (ijhc invalidJsonHttpClient) Get(url string) (*http.Response, error) {

	invalidBody := invalidJsonIOReader{}

	resp := &http.Response{}
	resp.Body = invalidBody

	return resp, nil
}
