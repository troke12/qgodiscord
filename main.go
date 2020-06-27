package main

import (
	"flag"
	"fmt"
	"net/http"
	"encoding/json"	
	"bytes"
)

type Message struct {
    Content string `json:"content"`
}

func main() {

	api := "MASUKIN_WEBHOOK_TOKEN"
	
	var message string
	flag.StringVar(&message, "m", "", "Message to send")
		
	flag.Parse()

	myMsg := Message{Content:message}
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
