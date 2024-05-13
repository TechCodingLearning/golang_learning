package unitTestSyntax

import (
	"net/http"
	"testing"
)

//const checkMark = "\u2713" // ✅
//const ballotX = "\u2717"   // ❎

// TestDownloadByTableTest 确认http包的Get函数可以下载内容并正确处理不同的状态
func TestDownloadByTableTest(t *testing.T) {
	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			"https://www.goinggo.net/feeds/posts/default?alt=rss",
			http.StatusOK,
		},
		{
			"https://rss.cnn.com/rss/cnn_topstbadurl.rss",
			http.StatusNotFound,
		},
	}

	t.Log("Given the need to test downloading different content.")
	{
		for _, u := range urls {
			t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
				u.url, u.statusCode)
			{
				resp, err := http.Get(u.url)
				if err != nil {
					t.Fatal("\t\tShould be able to Get the url.",
						ballotX, err)
				}
				t.Log("\t\tShould be able to Get the url",
					checkMark)

				defer resp.Body.Close()

				if resp.StatusCode == u.statusCode {
					t.Logf("\t\tShould have a \"%d\" status. %v",
						u.statusCode, checkMark)
				} else {
					t.Errorf("\t\tShould have a \"%d\" status %v %v",
						u.statusCode, ballotX, resp.StatusCode)
				}
			}
		}
	}
}
