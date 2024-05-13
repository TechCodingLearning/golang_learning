package handlers_test

import (
	"GolangProjects/handlers"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

func init() {
	handlers.Routes()
}

// ExampleSendJSON 提供了基础示例
func ExampleSendJSON() {
	r, _ := http.NewRequest("GET", "/sendjson", nil)
	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, r)
	var u struct {
		Name  string
		Email string
	}
	log.Println(rw.Body)
	if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
		log.Println("ERROR:", err)
	}

	// 使用fmt将结果写到stdout来检测输出
	fmt.Println(u)
	// Output:
	// {Bill bill@ardanstudios.com}
}
