package dio

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/line/line-bot-sdk-go/linebot"
)

type Dio struct {
	BotClient      *linebot.Client
	DefaultMessage *linebot.TextMessage
	OrderMemory    FoodOrderPersistance
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
	wryyyy  = linebot.NewTextMessage(`Wryyyyy!!! Confirm!!`)
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

	switch textMessage {
	case "#hiew":
		fmt.Println("Wryyyy!!!! Zawarudo!!!!!")
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
		//delete order
		order, err := d.OrderMemory.DeleteByID(orderID)
		if err != nil {
			d.MUDA(token)
			return err
		}

		// confirm message
		orderMenuMessage := linebot.NewTextMessage(order.Menu)
		if _, err := d.BotClient.ReplyMessage(token, orderMenuMessage).Do(); err != nil {
			d.MUDA(token)
			return err
		}

	default:
		if strings.Contains(textMessage, "#order") {
			//check is order
			menu, err := d.DecodeMenu(textMessage)
			if err != nil {
				d.MUDA(token)
				return err
			}
			fmt.Println("menu:", menu)
			menuCount := make(map[string]int)

			// get order if exists
			order, err := d.OrderMemory.FindByID(orderID)
			if err != nil && err != gorm.ErrRecordNotFound {
				d.MUDA(token)
				return err
			}
			if err == gorm.ErrRecordNotFound {
				order.CreatedAt = time.Now()
			}
			order.ID = orderID
			order.UpdatedAt = time.Now()

			fmt.Println("Menu from memory:", order.Menu)
			// count menu list
			menuList := strings.Split(order.Menu, "\n")
			fmt.Println("Menu list after split:", menuList, "len:", len(menuList))

			updateOrder := order
			menuResult := ""
			isNewMenu := true
			for _, mList := range menuList {
				if !strings.Contains(mList, ": ") {
					continue
				}
				mListSplit := strings.Split(mList, ": ")
				c, _ := strconv.Atoi(mListSplit[1])
				if strings.TrimSpace(mListSplit[0]) == strings.TrimSpace(menu) {
					isNewMenu = false
					c++
				}
				menuCount[mListSplit[0]] = c

				menuResult = fmt.Sprintf("%s\n%s: %d", menuResult, mListSplit[0], c)
			}
			if isNewMenu {
				menuResult = fmt.Sprintf("%s\n%s: %d", menuResult, menu, 1)
			}

			updateOrder.Menu = menuResult

			// is order update database
			updateOrder, err = d.OrderMemory.Upsert(updateOrder)
			if err != nil {
				d.MUDA(token)
				return err
			}

			//reply updated menu
			if _, err := d.BotClient.ReplyMessage(token, linebot.NewTextMessage(updateOrder.Menu)).Do(); err != nil {
				return err
			}
		}
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

	var nums []int
	// validate len
	if len(message) < 3 {
		return "", errors.New("invalid message")
	}

	// validate format must be number
	for _, mesByte := range message {
		n, err := strconv.Atoi(string(mesByte))
		if err != nil {
			return "", err
		}
		nums = append(nums, n)
	}

	order = fmt.Sprintf("%s %s %s", Menu[nums[0]-1], Special[nums[1]], Egg[nums[2]])

	return order, nil
}
