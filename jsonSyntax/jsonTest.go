package jsonSyntax

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	// gResult 映射到从搜索拿到的结果文档
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}
	// gResponse包含顶级的文档
	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

func TestJson() {
	//uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"
	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"

	// 向Google发起搜索
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("aaa")
	// 将JSON响应解码到结构类型
	var gr gResponse
	//var gr map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(gr)
}
