package myhttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client interface {
	GetURL() string
	GetRequest() (bytes []byte, err error)
}

type client struct {
	URL string
}

func NewClient(URL string) Client {
	return &client{URL: URL}
}

func (c *client) GetURL() string {
	return c.URL
}

func (c *client) GetRequest() (bytes []byte, err error) {
	if !isContainProtocol(c.URL) {
		c.URL = fmt.Sprintf("http://%s", c.URL)
	}

	resp, err := http.Get(c.URL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

// HTTPPrinter prints the request url and the md5 hash of the response
func HTTPPrinter(url, md5 string) {
	fmt.Printf("%s %s \n", url, md5)
}

func isContainProtocol(url string) bool {
	return strings.Contains(url, "http://")
}
