package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func ReceiveMessageHandler(w http.ResponseWriter, r *http.Request) {
	rawM, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Errorf("read message from request error:", err)
	}
	defer r.Body.Close()

	fmt.Println(fmt.Sprintf("message received: %+v", string(rawM)))
}
