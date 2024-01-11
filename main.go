package main

import (
	"os"
  models "github.com/BekirAktas/telego/models"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
  chatId := os.Getenv("chat_ids")

  message := "new message";
	tlg := models.NewTelegramMessage(message, chatId)
	tlg.SendMessage()
}
