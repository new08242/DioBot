package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ReceiveMessageHandler(w http.ResponseWriter, r *http.Request) {
	m := RawMessage{}

	rawM, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Errorf("read message from request error:", err)
	}
	defer r.Body.Close()

	fmt.Println(fmt.Sprintf("raw message received: %+v \n", string(rawM)))

	if err = json.Unmarshal(rawM, &m); err != nil {
		fmt.Errorf("json unmarshal message error:", err)
	}

	//if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
	//	fmt.Errorf("json unmarshal message error: %s", err)
	//}
	//

	fmt.Println(fmt.Sprintf("message received: %+v \n", m))

	if m.Events[0].Message.Type == MessageTypeText {
		if m.Events[0].Message.Text == "#hiew" {
			fmt.Println("Wryyyy!!!! Zawarudo The World!!!!!")
		}
	}
}
