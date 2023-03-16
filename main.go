package main

import (
	tgClient "NoteNukeBot/clients/telegram"
	config "NoteNukeBot/config"
	event_consumer "NoteNukeBot/consumer/event-consumer"
	"NoteNukeBot/events/telegram"
	"NoteNukeBot/storage/files"
	"flag"
	"log"
)

const (
	batchSize = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(config.TgBotHost, mustToken()),
		files.New(config.StoragePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("bot has been stopped")
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		config.Token,
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
