package app

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/receive_message", ReceiveMessageHandler)

	http.ListenAndServe(":9999", mux)
}
