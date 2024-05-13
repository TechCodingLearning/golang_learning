package jsonSyntax

import (
	"encoding/json"
	"fmt"
	"log"
)

// Contact 结构代表我们的JSON字符串
type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

// JSON 包含用于反序列化的演示字符串
var JSON1 = `{
	"name": "Gopher",
	"title": "programmer",
    "contact": {
		"home": "415.333.3333",
		"cell": "415.555.5555"
	}
}`

func TestJsonString() {
	fmt.Printf("%T\n", JSON1)

	// 将JSON字符串反序列化到变量
	var c Contact
	err := json.Unmarshal([]byte(JSON1), &c)
	if err != nil {
		log.Println("EROOR:", err)
		return
	}
	fmt.Println(c)
}
