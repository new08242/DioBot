package handler

import (
	diegoBrando "dio-bot/app/dio"

	"fmt"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Hello")
}

func ReceiveMessageHandler(w http.ResponseWriter, r *http.Request) {
	dio := diegoBrando.NewDio()
	dioBot := dio.BotClient

	events, err := dioBot.ParseRequest(r)
	if err != nil {
		fmt.Errorf("dio parse request error:", err)
	}
	fmt.Println(fmt.Sprintf("len: %v, message received: %+v \n", len(events), events[0]))

	if err := dio.HandleEvent(events); err != nil {
		fmt.Errorf("dio can't handle these events, error: %+v", err)
	}
}
