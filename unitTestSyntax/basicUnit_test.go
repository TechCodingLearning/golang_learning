package unitTestSyntax

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713" // ✅
const ballotX = "\u2717"   // ❎

// TestDownload 确认http包的Get函数可以下载内容
func TestDownloadByUnitTest(t *testing.T) { // 测试函数必须以Test单词开头，而且必须接受一个指向testing.T类型的指针，并且不返回任何值。否则，测试框架就不会认为这个函数是一个测试函数。
	url := "https://www.goinggo.net/feeds/posts/default?alt=rss"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
			url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.",
					ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get Call.",
				checkMark)
			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\"  status. %v",
					statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v",
					statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}
