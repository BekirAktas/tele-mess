package models

import (
	"encoding/json"
	"fmt"
    "os"
)

type TelegramMessage struct {
	Text   string `json:"text"`
	ChatID string `json:"chat_id"`
}

func NewTelegramMessage(message, chatId string) TelegramMessage {
	return TelegramMessage{
		Text:   message,
		ChatID: chatId,
	}
}

func (telegramMessage *TelegramMessage) SendMessage() {
    payload, err := json.Marshal(telegramMessage)
    if err != nil {
        fmt.Println("Error marshaling message:", err)
        return
    }

    sendMessageUrl := os.Getenv("base_url") + os.Getenv("api_token") + "/sendMessage"

    _, err = MakeRequest(sendMessageUrl, &payload, "POST")
    if err != nil {
        fmt.Println("Error sending message:", err)
        return
    }
}
