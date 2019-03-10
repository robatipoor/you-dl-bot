package main

import (
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func handleBot() *tb.Bot {
	bot, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatalln(err)
	}

	bot.Handle(tb.OnText, func(m *tb.Message) {
		log.Println("Get Text Message!")
		if m.Text == "/start" {
			bot.Send(m.Sender, "please send me url youtube !")
			return
		}
		if validYoutubeUrl(m.Text) {
			url,err := getDownloadURL(m.Text)
			if err != nil {
				log.Println(err)
				bot.Send(m.Sender, "Error !")
				return
			}
			bot.Send(m.Sender, url)
		}else{
			bot.Send(m.Sender, "url not valid !")
		}

	})

	return bot
}
