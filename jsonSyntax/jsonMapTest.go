package jsonSyntax

import (
	"encoding/json"
	"fmt"
	"log"
)

// JSON包含要反序列化的样例字符串
var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "415.333.3333",
		"cell": "415.555.5555"
	}
}`

func TestJsonMap() {
	var c map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println("Name:", c["name"])
	fmt.Println("Title:", c["title"])
	fmt.Println("Contact")
	fmt.Println("H:", c["contact"].(map[string]interface{})["home"]) // 需要将c["contact"]转换成合适的类型，因为原本c["contact"]是interface{}
	fmt.Println("C:", c["contact"].(map[string]interface{})["cell"])
}
