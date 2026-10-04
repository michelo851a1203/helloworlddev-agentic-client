package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(map[string]string{
		"message": "Hello World Dev",
	})
}

func handleSSE(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// TODO : 提示，這裡需要寫個迴圈
	// 跟時間有關的 可以用 time.NewTicker 這樣就可以每秒送
	// flusher, ok := w.(http.Flusher) 需要去判斷是否支援 event-stream
	// 送出去後記得後面要 Flush 提示小精靈就說到這囉 😃加油加油!
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /", handleRoot)
	router.HandleFunc("GET /sse", handleSSE)
	fmt.Println("run on http://localhost:8080")
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server : %v\n", err)
	}
}
