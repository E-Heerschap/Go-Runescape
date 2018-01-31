package Go_Runescape

import (
	"net/http"
	"errors"
)

//HttpClientWrap is made so the dev can use their own
//method of performing the http request.
type HttpClientWrap interface {
	Get(url string) (*http.Response, error)
}

//NotNilHttpClient is used in the testing files to test for
//not nil errors.
type NotNilHttpClient struct {
	HttpClientWrap
}

func (nnhc NotNilHttpClient) Get(url string) (*http.Response, error){
	err := errors.New("TEST ERROR")
	return &http.Response{}, err
}