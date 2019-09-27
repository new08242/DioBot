package main

import (
	"net/http"

	"dio-bot/app/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/receive_message", handler.ReceiveMessageHandler)
	mux.HandleFunc("/", handler.PingHandler)

	http.ListenAndServe(":9999", mux)
}
