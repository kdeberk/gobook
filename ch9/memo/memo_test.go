package memo

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMemoGet1(t *testing.T) {
	m := New(httpGetBody)
	testMemoGetFn(t, m.Get1)
}

func TestMemoGet2(t *testing.T) {
	m := New(httpGetBody)
	testMemoGetFn(t, m.Get2)
}

func TestMemoGet3(t *testing.T) {
	m := New(httpGetBody)
	testMemoGetFn(t, m.Get3)
}

func TestMemoGet4(t *testing.T) {
	m := New(httpGetBody)
	testMemoGetFn(t, m.Get4)
}

func testMemoGetFn(t *testing.T, fn func(string) (interface{}, error)) {
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()

			start := time.Now()
			value, err := fn(url)
			if err != nil {
				t.Error(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}

func incomingURLs() <-chan string {
	urls := []string{"https://golang.org", "https://godoc.org", "https://play.golang.org", "http://gopl.io"}

	urlC := make(chan string)
	go func() {
		for i := 0; i < 2; i++ {
			for _, url := range urls {
				urlC <- url
			}
		}
		close(urlC)
	}()
	return urlC
}
