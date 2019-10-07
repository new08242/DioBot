package dio

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"
)

var (
	Menu = make(map[int]string)
	message =
`%s
1. หมูทอดคั่วพริกเกลือ+กะหล่ำทอดน้ำปลา
2. ไก่กรอบคั่วพริกเกลือ+กะหล่ำทอดน้ำปลา
3. หมูทอดคั่วพริกเกลือ
4. ไก่กรอบคั่วพริกเกลือ
5. ผัดมาม่าหมู
6. สุกี้แห้ง

manual order format:
เลขเมนู ธรรมดา0พิเศษ1 ไม่ไข่0ไข่ดาว1ไข่เจียว2
ตัวอย่าง
- ผัดมาม่าหมู พิเศษ ไข่เจียว
ให้สั่ง "5 1 2"
- หมูทอดคั่วพริกเกลือ ธรรมดา
ให้สั่ง "3 0 0"
`
	Special = make(map[int]string)

	Egg = make(map[int]string)

	MenuMessage = linebot.NewTextMessage(message)
)

func init(){
	Menu[0] = "หมูทอดคั่วพริกเกลือ+กะหล่ำทอดน้ำปลา"
	Menu[1] = "ไก่กรอบคั่วพริกเกลือ+กะหล่ำทอดน้ำปลา"
	Menu[2] = "หมูทอดคั่วพริกเกลือ"
	Menu[3] = "ไก่กรอบคั่วพริกเกลือ"
	Menu[4] = "ผัดมาม่าหมู"
	Menu[5] = "สุกี้แห้ง"

	menuList := ""
	//gen menu message
	for i, m := range Menu {
		menuList = fmt.Sprintf("%s\n%d.%s", menuList, i+1, m)
	}

	message = fmt.Sprintf(message, menuList)

	MenuMessage = linebot.NewTextMessage(message)
}
