package mutexSyntex

import (
	"errors"
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch 返回URL的body内容，并且将这个页面上找到的URL放到一个slice中
	Fetch(url string) (body string, urls []string, err error)
}

var vis = struct {
	m map[string]error
	sync.Mutex
}{m: make(map[string]error)}

var loading = errors.New("url load in progress")

// Crawl 使用fetcher从某个URL开始递归地爬取页面，直到达到最大深度
func Crawl(url string, depth int, fetcher Fetcher) {

	if depth <= 0 {

		return
	}

	vis.Lock()

	if _, ok := vis.m[url]; ok {
		vis.Unlock()
		return
	}
	vis.m[url] = loading
	vis.Unlock()

	body, urls, err := fetcher.Fetch(url)

	vis.Lock()
	vis.m[url] = err
	vis.Unlock()

	if err != nil {
		fmt.Printf("Error: url{%v} doesn't exist!\n", url)
		return
	}

	fmt.Printf("url{%s} -> body: %q\n", url, body)

	ch := make(chan bool) // 方便阻塞的时候用
	for _, u := range urls {
		go func(url string) {
			Crawl(url, depth-1, fetcher)
			ch <- true
		}(u)
	}

	for _ = range urls {
		<-ch // 阻塞到协程结束
	}
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的fakeFetcher
var fetcher = &fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

func TestMutexCase() {
	Crawl("https://golang.org/", 4, fetcher)
	// time.Sleep(100000)  // 主线程必须等到协程运行完结束，否则看不到结果，因为主线程结束，协程就不用运行了----已经通过channel优化
}
