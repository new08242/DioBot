package handler

import (
	diegoBrando "dio-bot/app/dio"

	"fmt"
	"net/http"
	"io/ioutil"
)

func Ping(w http.ResponseWriter, r *http.Request) {}

func ReceiveMessageHandler(w http.ResponseWriter, r *http.Request) {
	dio := diegoBrando.NewDio()
	dioBot := dio.BotClient

	rawReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("[ReceiveMessageHandler] read request:", err)
		return
	}
	fmt.Println("[ReceiveMessageHandler] dio get raw request:", string(rawReq))

	events, err := dioBot.ParseRequest(r)
	if err != nil {
		fmt.Println("[ReceiveMessageHandler] dio parse request error:", err)
		return
	}

	if err := dio.HandleEvent(events); err != nil {
		fmt.Println("[ReceiveMessageHandler] dio can not handle these events, error:", err)
		return
	}
}
