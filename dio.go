package main

import (
	"fmt"
	"errors"

	"github.com/line/line-bot-sdk-go/linebot"
)

type Dio struct {
	BotClient *linebot.Client
	LIFFURL   *linebot.TextMessage
	DefaultMessage *linebot.TextMessage
}

func NewDio() Dio {
	channelSecret := "3c11fac7b55b109d58ebc6ccd0307ac0"
	channelAccesstoken := "Lnh5vmM/RQCHPYTQ5kC2n6UxZcuHx+6h095FC+XLoFxEmypvGIJvGibTnohk+bPvlfCzHIoqzeiJWsThEgG+2mWiOrnMsinBpWxmdkXnCB33FTbwPV6whCvkL9GasSPouf13WT1PYk/wVsbUZHIvvFGUYhWQfeY8sLGRXgo3xvw="
	bot, err := linebot.New(channelSecret, channelAccesstoken)
	if err != nil {
		return Dio{}
	}

	return Dio{
		BotClient: bot,
		LIFFURL:   linebot.NewTextMessage("line://app/1622344082-zYW59LY3"),
		DefaultMessage: linebot.NewTextMessage("MUDA MUDA!!!"),
	}
}

func (d Dio) MUDA(token string) error {
	// MUDAAAAA!!! for unhandle case
	if _, err := d.BotClient.ReplyMessage(token, d.DefaultMessage).Do(); err != nil {
		return err
	}
	return errors.New("muda muda")
}

func (d Dio) HandleEvent(events []*linebot.Event) error {
	if len(events) < 1 {
		return errors.New("no event for Dio")
	}

	for _, event := range events {
		switch event.Type {
		case linebot.EventTypeMessage:
			if err := d.EventTypeMessageHandler(event); err != nil {
				return err
			}
		default:
			if err := d.MUDA(event.ReplyToken); err != nil {
				return err
			}
		}
	}

	return nil
}

func (d Dio) EventTypeMessageHandler(event *linebot.Event) error {
	//source := *event.Source
	//switch source.Type {
	//case linebot.EventSourceTypeUser:
	//	return nil
	//case linebot.EventSourceTypeGroup:
	//	return nil
	//case linebot.EventSourceTypeRoom:
	//	return nil
	//default:
	//	return errors.New(fmt.Sprintf("no handler for event type: %+v", source.Type))
	//}

	fmt.Printf("[EventTypeMessageHandler] event: %+v \n", *event)
	fmt.Printf("[EventTypeMessageHandler] event source: %+v \n", *event.Source)
	fmt.Printf("[EventTypeMessageHandler] event message: %+v \n", event.Message)

	switch m := event.Message.(type) {
	case *linebot.TextMessage:
		if err := d.MessageTypeTextHandler(event.ReplyToken, m.Text); err != nil {
			return err
		}
	default:
		if err := d.MUDA(event.ReplyToken); err != nil {
			return err
		}
	}

	return nil
}

func (d Dio) MessageTypeTextHandler(token, message string) error {
	switch message {
	case "#hiew":
		fmt.Println("Wryyyy!!!! Zawarudo The World!!!!!")
		if _, err := d.BotClient.ReplyMessage(token, d.LIFFURL).Do(); err != nil {
			return err
		}
	default:
		if err := d.MUDA(token); err != nil {
			return err
		}
	}

	return nil
}
