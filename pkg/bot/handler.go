package bot

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
)

var AddToManagerFuncs []func(bot *Bot) error

var replyBtn = tb.ReplyButton{Text: "台灣鐵路查詢及訂票"}
var replyKeys = [][]tb.ReplyButton{
	[]tb.ReplyButton{replyBtn},
}

func init() {
	AddToManagerFuncs = append(AddToManagerFuncs, handle, traHandle)
}

func handle(b *Bot) error {
	var err error

	b.Handle(&replyBtn, func(m *tb.Message) {
		msg, err := b.Send(m.Sender, "我現在幫您查詢台鐵資料....", ComRegister())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(msg)
	})

	b.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		_, err := b.Send(m.Sender, "Hello!", &tb.ReplyMarkup{
			ReplyKeyboard: replyKeys,
		})
		if err != nil {
			log.Fatal(err)
		}
	})
	return err
}
