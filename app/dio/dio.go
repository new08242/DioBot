package dio

import (
	"errors"
	"fmt"
	"strings"
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
)

type Dio struct {
	BotClient      *linebot.Client
	DefaultMessage *linebot.TextMessage
	OrderMemory		FoodOrderPersistance
}

func NewDio() Dio {
	channelSecret := "3c11fac7b55b109d58ebc6ccd0307ac0"
	channelAccesstoken := "Lnh5vmM/RQCHPYTQ5kC2n6UxZcuHx+6h095FC+XLoFxEmypvGIJvGibTnohk+bPvlfCzHIoqzeiJWsThEgG+2mWiOrnMsinBpWxmdkXnCB33FTbwPV6whCvkL9GasSPouf13WT1PYk/wVsbUZHIvvFGUYhWQfeY8sLGRXgo3xvw="
	bot, err := linebot.New(channelSecret, channelAccesstoken)
	if err != nil {
		return Dio{}
	}

	return Dio{
		BotClient:      bot,
		DefaultMessage: linebot.NewTextMessage("MUDA MUDA!!!"),
		OrderMemory:    FoodOrderPersistance{db: GetMemory()},
	}
}

var (
	wryyyy = linebot.NewTextMessage(`Wryyyyy!!! Confirm!!`)
	LIFFURL = linebot.NewTextMessage("line://app/1622344082-zYW59LY3")
)

func (d Dio) MUDA(token string) error {
	// MUDAAAAA!!! for unhandled case
	if _, err := d.BotClient.ReplyMessage(token, d.DefaultMessage).Do(); err != nil {
		return err
	}
	return errors.New("muda muda")
}

func (d Dio) HandleEvent(events []*linebot.Event) error {
	fmt.Println("Wryy handle event")
	if len(events) < 1 {
		return errors.New("no event for Dio")
	}

	for _, event := range events {
		switch event.Type {
		case linebot.EventTypeMessage:
			if err := d.EventTypeMessageHandler(*event); err != nil {
				return err
			}
		default:
			d.MUDA(event.ReplyToken)
		}
	}

	return nil
}

func (d Dio) EventTypeMessageHandler(event linebot.Event) error {
	fmt.Printf("[EventTypeMessageHandler] event: %+v \n", event)
	fmt.Printf("[EventTypeMessageHandler] event source: %+v \n", *event.Source)
	fmt.Printf("[EventTypeMessageHandler] event message: %+v \n", event.Message)

	switch m := event.Message.(type) {
	case *linebot.TextMessage:
		if err := d.MessageTypeTextHandler(event, m.Text); err != nil {
			return err
		}
	default:
		if err := d.MUDA(event.ReplyToken); err != nil {
			return err
		}
	}

	return nil
}

func (d Dio) MessageTypeTextHandler(event linebot.Event, textMessage string) error {
	token := event.ReplyToken

	orderID := ""
	if event.Source.GroupID != "" {
		orderID = event.Source.GroupID

	} else if event.Source.RoomID != "" {
		orderID = event.Source.RoomID

	} else if event.Source.UserID != "" {
		orderID = event.Source.UserID
	}


	switch message {
	case "#hiew":
		fmt.Println("Wryyyy!!!! Zawarudo The World!!!!!")
		if _, err := d.BotClient.ReplyMessage(token, LIFFURL).Do(); err != nil {
			return err
		}
	case "#menu":
		fmt.Println("WAMU!!!!!")
		if _, err := d.BotClient.ReplyMessage(token, MenuMessage).Do(); err != nil {
			return err
		}
	case "#confirm":
		fmt.Println("ROAD ROLLER DA!!!!!")
		// get order if exists
		order, err := d.OrderMemory.FindByID(orderID)
		if err != nil {
			d.MUDA(token)
			return err
		}

		// confirm message
		orderMenuMessage := linebot.NewTextMessage(order.Menu)
		if _, err := d.BotClient.ReplyMessage(token, orderMenuMessage).Do(); err != nil {
			return err
		}

		//delete order
		if err := d.OrderMemory.DeleteByID(orderID); err != nil {
			d.MUDA(token)
			return err
		}

	case "#order":
		//check is order
		menu, err := d.DecodeMenu(message)
		if err != nil {
			d.MUDA(token)
			return err
		}

		// get order if exists
		order, err := d.OrderMemory.FindByID(orderID)
		if err != nil {
			d.MUDA(token)
			return err
		}

		// is order update database
		updateOrder := order
		updateOrder.Menu = fmt.Sprintf("%s\n%s", order.Menu, menu)
		updateOrder, err = d.OrderMemory.Upsert(updateOrder)
		if err != nil {
			d.MUDA(token)
			return err
		}

		//reply updated menu
		if _, err := d.BotClient.ReplyMessage(token, linebot.NewTextMessage(updateOrder.Menu)).Do(); err != nil {
			return err
		}

	default:

	}

	return nil
}

func (d Dio) DecodeMenu(message string) (string, error) {
	order := ""

	message = strings.TrimSpace(message)
	messages := strings.Split(message, "#order")
	if len(messages) < 2 {
		return "", errors.New("invalid message")
	}
	message = messages[1]

	orderNums := strings.Split(message, " ")
	var nums []int

	// validate format must be number
	if len(orderNums) == 3 {
		for _, num := range orderNums {
			n, err := strconv.Atoi(num)
			if err != nil {
				return "", err
			}
			nums = append(nums, n)
		}
	}

	order = fmt.Sprintf("%s %s %s", Menu[nums[0]], Menu[nums[1]], Menu[nums[2]])

	return order, nil
}
