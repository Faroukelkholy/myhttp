package myhttp

import (
	"fmt"
	"runtime"
	"sync"
)

func Start(limit int, urls []string) (err error) {
	// make available the maximum number possible of available cpu to execute go routine simultaneously
	runtime.GOMAXPROCS(runtime.NumCPU())

	// create wait group that waits until go routines finish requesting urls
	wg := &sync.WaitGroup{}
	wg.Add(limit)

	ch := fillChan(urls)

	for i := 0; i < limit; i++ {
		go func() {
			taskErr := performTask(ch, wg)
			if taskErr != nil {
				err = fmt.Errorf("%w", taskErr)
			}
		}()
	}

	wg.Wait()

	return
}

// fillChan sends the urls in a channel
func fillChan(urls []string) <-chan string {
	ch := make(chan string)
	go func(){
		defer close(ch)

		for _, url := range urls {
			ch <- url
		}
	}()

	return ch
}

// performTask holds the logic that handles http get request for an url and hashing the response.
// printing the url and the hash
func performTask(ch <-chan string, wg *sync.WaitGroup) error {
	defer wg.Done()

	for url := range ch {
		c := NewClient(url)

		resp, err := c.GetRequest()
		if err != nil {
			return err
		}

		md5 := HashMD5(resp)

		HTTPPrinter(c.GetURL(), md5)
	}

	return nil
}
