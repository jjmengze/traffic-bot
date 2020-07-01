package bot

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
	"traffic-bot/pkg/controller/middleware"
)

type Bot struct {
	*tb.Bot
	handlerContext *middleware.HandlerContext
}

func NewBot(token string, handlerContext *middleware.HandlerContext) *Bot {
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	return &Bot{
		b,
		handlerContext,
	}
}

func (b *Bot) Register() {
	for _, fun := range AddToManagerFuncs {
		if err := fun(b); err != nil {
			log.Fatal(err)
		}
	}
}
