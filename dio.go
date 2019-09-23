package main

import "github.com/line/line-bot-sdk-go/linebot"

type Dio struct {
	BotClient *linebot.Client
	LIFFURL *linebot.TextMessage
}

func(d Dio) NewDio() Dio {
	channelSecret := "3c11fac7b55b109d58ebc6ccd0307ac0"
	channelAccesstoken := "Lnh5vmM/RQCHPYTQ5kC2n6UxZcuHx+6h095FC+XLoFxEmypvGIJvGibTnohk+bPvlfCzHIoqzeiJWsThEgG+2mWiOrnMsinBpWxmdkXnCB33FTbwPV6whCvkL9GasSPouf13WT1PYk/wVsbUZHIvvFGUYhWQfeY8sLGRXgo3xvw="
	bot, err := linebot.New(channelSecret, channelAccesstoken)
	if err != nil {
		return Dio{}
	}

	return Dio{
		BotClient: bot,
		LIFFURL: linebot.NewTextMessage("line://app/1622344082-zYW59LY3"),
	}
}
