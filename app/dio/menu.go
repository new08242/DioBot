package dio

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"
)

var (
	Menu []string
	message = `manual order format:
เลขเมนู ธรรมดา0พิเศษ1 ไม่ไข่0ไข่ดาว1ไข่เจียว2

ตัวอย่าง
- ผัดมาม่าหมู พิเศษ ไข่เจียว
ให้สั่ง "#order512"
- หมูทอดคั่วพริกเกลือ ธรรมดา
ให้สั่ง "#order300"
`
	Special = make(map[int]string)

	Egg = make(map[int]string)

	MenuMessage = linebot.NewTextMessage(message)
)

func init() {
	Menu = append(Menu,"หมูทอดคั่วพริกเกลือ+กะหล่ำทอดน้ำปลา")
	Menu = append(Menu,"ไก่กรอบคั่วพริกเกลือ+กะหล่ำทอดน้ำปลา")
	Menu = append(Menu,"หมูทอดคั่วพริกเกลือ")
	Menu = append(Menu,"ไก่กรอบคั่วพริกเกลือ")
	Menu = append(Menu,"ผัดมาม่าหมู")
	Menu = append(Menu,"สุกี้แห้ง")
	Menu = append(Menu,"กระเพราหมู")

	Special[0] = "ธรรมดา"
	Special[1] = "พิเศษ"

	Egg[0] = ""
	Egg[1] = "ไข่ดาว"
	Egg[2] = "ไข่เจียว"

	menuList := ""
	//gen menu message
	for i, _ := range Menu {
		menuList = fmt.Sprintf("%s\n%d.%s", menuList, i+1, Menu[i])
	}

	message = fmt.Sprintf("%s\n\n%s", menuList, message)

	MenuMessage = linebot.NewTextMessage(message)
}
