package unitTestSyntax

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

//const checkMark = "\u2713" // ✅
//const ballotX = "\u2717"   // ❎

// feed 模仿了我们期望接收的XML文档
var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
	<title>Going Go Programming</title>
	<description>Golang : https://github.com/goinggo</description>
	<link>https://www.goinggo.net</link>
	<item>
		<pubDate>Sun, 15 Mar 2015 15:04:00 ++0000</pubDate>
		<title>Object Oriented Programming Mechanics</title>
		<description>Go is an object oriented language.</description>
		<link>https://www.goinggo.net/2015/03/object-oriented</link>
	</item>
</channel>
</rss>
`

// mockServer 返回用来处理请求的服务器的指针
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}
	// HandlerFunc类型是一个适配器，允许常规函数作为HTTP的处理函数使用；
	// 如果函数f具有合适的签名，HandlerFunc(f)就是一个处理HTTP请求的Handler对象，内部通过调用f处理请求
	return httptest.NewServer(http.HandlerFunc(f))
}

// TestDownloadByMock 确认http包的Get函数可以下载内容并且内容可以被正确地反序列化并关闭
func TestDownloadByMock(t *testing.T) {
	statusCode := http.StatusOK

	server := mockServer()
	defer server.Close()

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
			server.URL, statusCode)
		{
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.",
					ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.",
				checkMark)
			defer resp.Body.Close()

			if resp.StatusCode != statusCode {
				t.Fatalf("\t\tShould receive a \"%d\" status. %v %v",
					statusCode, ballotX, resp.StatusCode)
			}
			t.Logf("\t\tShould receive a \"%d\" status. %v",
				statusCode, checkMark)
		}
	}
}
