package myhttp

import (
	"reflect"
	"sync"
	"testing"
)

var (
	urls = []string{
		"twitter.com",
		"google.com",
		"facebook.com",
	}

	urlsWrong = []string{
		"google",
		"facebook",
		"twitter.com",
		"google.com",
		"facebook.com",
		"baroquemusiclibrary.com",
	}
)

func TestFillChan(t *testing.T) {
	t.Parallel()

	var gotURLS []string

	ch := fillChan(urls)

	for v := range ch {
		gotURLS = append(gotURLS, v)
	}

	if !reflect.DeepEqual(gotURLS, urls) {
		t.Errorf("fillChan got %v, want %v", gotURLS, urls)
	}
}

func TestPerformTask_success(t *testing.T){
	t.Parallel()

	ch, wg := setupTestPerformTask(urls)

	err := performTask(ch,wg)

	if err != nil {
		t.Errorf("performTaskt should not return error; got: %s, want: nil",err.Error())
	}
}

func TestPerformTas_error(t *testing.T){
	t.Parallel()

	ch, wg := setupTestPerformTask(urlsWrong)

	err := performTask(ch,wg)

	if err == nil {
		t.Errorf("performTaskt should return error; got: %s, want: nil",err.Error())
	}
}

func TestStart_success(t *testing.T){
	t.Parallel()

	err := Start(5, urls)

	if err != nil {
		t.Errorf("TestStart should not return error; got: %s, want: nil",err.Error())
	}
}

func TestStart_error(t *testing.T){
	t.Parallel()

	err := Start(5, urlsWrong)

	if err == nil {
		t.Errorf("TestStart should return error; got: %s, want: nil",err.Error())
	}
}

func setupTestPerformTask(urls []string) (ch <-chan string, wg *sync.WaitGroup) {
	ch = fillChan(urls)

	wg = &sync.WaitGroup{}
	wg.Add(len(urls))

	return
}
