package main

import (
	"fmt"
	//"httpclient"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type jsonMessage struct {
	Id   string `json:"id"`
	Time int64  `json:"time"`
	Body string `json:"body"`
}

func main() {

	resp, err := http.Get("https://jch.irif.fr:8082/chat/messages.json?count=4")

	if err != nil {
		log.Fatal(err)
	}

	head := resp.Header.Get("eTag")

	for true {

		resp, err = http.Get("https://jch.irif.fr:8082/chat/messages.json?count=4")

		if err != nil {
			log.Fatal(err)
		}

		new := resp.Header.Get("eTag")
		if head != new {
			head = new

			var messages []jsonMessage

			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				log.Fatal(err)
			}

			json.Unmarshal(body, &messages)

			for i := 0; i < len(messages); i++ {
				fmt.Println("------------------------------------------")
				fmt.Println("===> ", resp.Header.Get("If-None-Match"))
				fmt.Println("id=" + messages[i].Id)
				fmt.Println("time=", messages[i].Time)
				fmt.Println("body=" + messages[i].Body)
			}
		}
	}

}

func exo1() {
	//resp, err := http.Get("https://jch.irif.fr:8082/chat/messages.json?count=4")
	resp, err := http.Get("https://jch.irif.fr:8082/chat/messages.json")

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var messages []jsonMessage

	json.Unmarshal(body, &messages)

	for i := 0; i < len(messages); i++ {
		fmt.Println("------------------------------------------")
		fmt.Println("id=" + messages[i].Id)
		fmt.Println("time=", messages[i].Time)
		fmt.Println("body=" + messages[i].Body)
	}
}
