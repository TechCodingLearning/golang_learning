package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// Routes set the routes for the web service.
func Routes() {
	http.HandleFunc("/sendjson", SendJSON)
}

// SendJSON returns a simple JSON document.
func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Bill",
		Email: "bill@ardanstudios.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	err := json.NewEncoder(rw).Encode(&u)
	if err != nil {
		log.Fatalln("Encode Error:", err)
	}
}

// TestHandlers 测试Handler
func TestHandlers() {
	Routes()

	port := "10000" // 端口可能被占用
	log.Printf("listener : Started : Listening on :%v", port)
	log.Printf("Please use your browser: localhost:%v/sendjson", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
