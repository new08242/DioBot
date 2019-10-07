package main

import (
	"fmt"
	"net/http"

	"dio-bot/app/dio"
	"dio-bot/app/handler"
)

func main() {
	db, err := dio.NewMemory()
	if err != nil {
		fmt.Errorf("dio new memory error:", err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Ping)
	mux.HandleFunc("/dio/", handler.Ping)
	mux.HandleFunc("/dio/receive_message", handler.ReceiveMessageHandler)

	http.ListenAndServe(":9999", mux)
}
