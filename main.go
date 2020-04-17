package main

import (
	"flag"
	"fmt"
	"net/http"
	"encoding/json"	
	"bytes"
)

type Message struct {
	Username string `json:"username"`
    Content string `json:"content"`
}

func main() {

	//Discord webhook api url flag
	var api string
	flag.StringVar(&api, "a", "",  "Discord webhook URL of server to send message to")
	
	var message string
	flag.StringVar(&message, "m", "", "Message to send")
	
	var user string
	flag.StringVar(&user, "u", "Captain Hook", "Discord Username")
	
	//debug flags
	//var debug bool
	//flag.BoolVar(&debug, "debug", false, "Display debug logs with -debug option")
	
	flag.Parse()
	//fmt.Println("username", user)
	//fmt.Println("message", message)
	
	myMsg := Message{Username:user,Content:message}
	b, err := json.Marshal(myMsg)
	
	if err != nil {
		fmt.Println("error:", err)	
	}
	//os.Stdout.Write(b)
	SendMessage(api, string(b))
}

func SendMessage(api string, jsonData string) {

	var jsonStr = []byte(jsonData)

	req, err := http.NewRequest("POST", api, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}