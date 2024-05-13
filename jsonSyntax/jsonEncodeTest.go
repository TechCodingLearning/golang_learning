package jsonSyntax

import (
	"encoding/json"
	"fmt"
	"log"
)

func TestJsonEncode() {
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}

	data, err := json.MarshalIndent(c, "", "	")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Println(string(data))
}
