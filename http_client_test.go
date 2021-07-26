package myhttp

import (
	"testing"
)

var (
	url = "http://google.com"
	c = NewClient(url)
)

func TestGetURL(t *testing.T) {
	got := c.GetURL()

	if got != url {
		t.Errorf("get url returned wrong string, got: %s ,want: %s",got , url)
	}
}

func TestClientGetRequest_success(t *testing.T) {
	resp, err := c.GetRequest()

	if err != nil {
		t.Errorf("get request should not return error; got: %s, want: nil",err.Error())
	}

	if resp == nil {
		t.Errorf("get request should return resp; got: nil")
	}
}

func TestClientGetRequest_error(t *testing.T) {
	c = NewClient("goo")
	resp, err := c.GetRequest()

	if err == nil {
		t.Errorf("get request should not return error; got: nil, want: error")
	}

	if resp != nil {
		t.Errorf("get request should not return resp; got: value, want: nil")
	}
}

func TestIsContainProtocol_true(t *testing.T) {
	ok := isContainProtocol(url)

	if !ok {
		t.Errorf("url does not contains http protocol, got: %t ,want: %t",ok , true )
	}
}

func TestIsContainProtocol_false(t *testing.T) {
	ok := isContainProtocol(url)

	if !ok {
		t.Errorf("url does contains http protocol, got: %t ,want: %t", ok , false )
	}
}
