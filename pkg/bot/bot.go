package bot

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
	"traffic-bot/pkg/bot/handler"
)

var AddToManagerFuncs []func(bot *tb.Bot) error

type Bot struct {
	*tb.Bot
}

func NewBot(token string) *Bot {
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}
	return &Bot{b}
}

func (b *Bot) Register() {
	for _, fun := range handler.AddToManagerFuncs {
		if err := fun(b.Bot); err != nil {
			log.Fatal(err)
		}
	}
}


