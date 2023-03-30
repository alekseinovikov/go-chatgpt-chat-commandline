package main

import (
	"go-chatgpt-chat-commandline/tui"
	"log"
)

func main() {
	model, err := tui.SelectGptModel()
	if err != nil {
		log.Fatalln(err)
	}

	apiKey, err := tui.GetApiKey()
	if err != nil {
		log.Fatal(err)
	}

	err = tui.RunChat(model, apiKey)
	if err != nil {
		log.Fatal(err)
	}
}
