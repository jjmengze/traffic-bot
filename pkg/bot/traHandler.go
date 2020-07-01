package bot

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"traffic-bot/pkg/controller/handler/tra/actiontype"
	"traffic-bot/pkg/controller/middleware"
	"traffic-bot/pkg/utils/xtime"
)

type TRA struct{}

var replySchBtn = tb.ReplyButton{Text: "查詢時刻表"}
var reservedBtn = tb.ReplyButton{Text: "非對號"}
var noReservedBtn = tb.ReplyButton{Text: "對號"}

func ComRegister() *tb.ReplyMarkup {
	replyKeys := [][]tb.ReplyButton{{replySchBtn}}
	return &tb.ReplyMarkup{
		ReplyKeyboard: replyKeys,
	}
}

func typeRegister() *tb.ReplyMarkup {
	replyKeys := [][]tb.ReplyButton{{reservedBtn, noReservedBtn}}
	return &tb.ReplyMarkup{
		ReplyKeyboard: replyKeys,
	}
}

func traHandle(b *Bot) error {
	var err error

	b.Handle(&replyBtn, func(m *tb.Message) {
		msg, err := b.Send(m.Sender, "好的，請問您要查詢的車種？", typeRegister())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(msg)
	})
	//menu := &tb.ReplyMarkup{}
	//menu
	//btn := menu.Data("基隆市", "apple", "apple")
	//btn2 := menu.Data("基隆市", "基隆市")
	//btn := tb.Btn{Text: "基隆市", Data: "apple"}
	//btn2 := tb.Btn{Text: "基隆市", Data: "基隆市"}
	//rows[idx] =
	//b.Handle(&btn, func(c *tb.Callback) {
	//	fmt.Println(c.Data)
	//
	//	// ...
	//	// Always respond!
	//	b.Respond(c, &tb.CallbackResponse{
	//		Text: "apple",
	//	})
	//})
	//b.Handle(&btn2, func(c *tb.Callback) {
	//	fmt.Println(c.Data)
	//	// ...
	//	// Always respond!
	//	b.Respond(c, &tb.CallbackResponse{
	//		Text: c.Data,
	//	})
	//})
	//menu.Inline(menu.Row(btn, btn2))

	b.Handle(&reservedBtn, func(m *tb.Message) {
		// on reply button pressed
		_, err := b.Send(m.Sender, "您選擇對號列車 請選擇您要查詢的車站")
		if err != nil {
			log.Fatal(err)
		}
		b.handlerContext.GetMiddle(middleware.TRA).EventCh <- actiontype.EventInfo{
			Action: actiontype.QUERY,
			Type:   actiontype.City}

	})
	//eventCh := make(chan int)

	b.Handle(&noReservedBtn, func(m *tb.Message) {
		// on reply button pressed
		msg, err := b.Send(m.Sender, "您選擇一般列車! 請輸入查詢日期。")
		if err != nil {
			log.Fatal(err)
		}
		times := xtime.GetTme()
		msg, err = b.Send(m.Sender, fmt.Sprintf("輸入格式為：%v %v %v ", times.Year(), int(times.Month()), times.Day()))
		msg, err = b.Reply(msg, "輸入現在，代表當天日期。")
		if err != nil {
			log.Fatal(err)
		}
		//eventCh <- m.Sender.ID
		//b.Send(m.Sender, "請選擇欲搭乘地點", menu)

	})

	//menu.Inline(menu.Row(btn))
	//go eventHandler(b.Bot, eventCh)
	return err
}

func eventHandler(b *tb.Bot, ch chan int) {
	menu := &tb.ReplyMarkup{ReplyKeyboardRemove: true}

	//menu.InlineKeyboard = append(menu.InlineKeyboard, cityName...)
	//rows := make([]tb.Row, len(cityMapBtn))
	rows := make([]tb.Row, len(cityMapBtn))
	rp := &tb.ReplyMarkup{ReplyKeyboardRemove: true}
	for i := 0; i < len(cityMapBtn); i++ {
		rows[i] = menu.Row(cityMapBtn[i])
		b.Handle(&cityMapBtn[i], func(c *tb.Callback) {
			fmt.Println(c.Data)
			eeee := make([]tb.Row, len(cityMapping[c.Data]))

			for index, item := range cityMapping[c.Data] {
				eeee[index] = menu.Row(tb.Btn{Text: item})
			}
			rp.Inline(eeee...)

			b.Send(&tb.User{ID: c.Sender.ID}, "請選擇欲搭乘地點", rp)
		})
	}
	//for idx, btn := range cityMapBtn {
	//	rows[idx] = menu.Row(btn)
	//	b.Handle(&btn, func(c *tb.Callback) {
	//		fmt.Println(c.Data)
	//	})
	//}
	//
	//}
	//btn := tb.Btn{Text: "apple"}
	//btn := menu.Data("ℹ Help")
	//rows[idx] =
	//b.Handle(&btn, func(c *tb.Callback) {
	//	// ...
	//	// Always respond!
	//	b.Respond(c, &tb.CallbackResponse{
	//		Text: "apple",
	//	})
	//})
	//
	//menu.Inline(menu.Row(btn))
	//
	menu.Inline(rows...)

	fmt.Println(menu.InlineKeyboard)
	for {
		id := <-ch
		b.Send(&tb.User{ID: id}, "請選擇欲搭乘地點", menu)
	}
}

//type btn [][]tb.InlineButton
//
//func
//btnToRow(btn
//[][]tb.Btn) {
//}

var cityMapBtn = []tb.Btn{
	{Text: "基隆市", Data: "基隆市", Unique: "apple"},
	//{Text: "新北市", Data: "新北市"},
	//{Text: "臺北市", Data: "臺北市"},
	//{Text: "桃園市", Data: "桃園市"},
	//{Text: "新竹縣", Data: "新竹縣"},
	//{Text: "新竹市", Data: "新竹市"},
	//{Text: "苗栗縣", Data: "苗栗縣"},
	//{Text: "臺中市", Data: "臺中市"},
	//{Text: "彰化縣", Data: "彰化縣"},
	//{Text: "南投縣", Data: "南投縣"},
	//{Text: "雲林縣", Data: "雲林縣"},
	//{Text: "嘉義縣", Data: "嘉義縣"},
	//{Text: "嘉義市", Data: "嘉義市"},
	//{Text: "臺南市", Data: "臺南市"},
	//{Text: "高雄市", Data: "高雄市"},
	//{Text: "屏東縣", Data: "屏東縣"},
	//{Text: "臺東縣", Data: "臺東縣"},
	//{Text: "花蓮縣", Data: "花蓮縣"},
	//{Text: "宜蘭縣", Data: "宜蘭縣"},
}

type cityStation struct {
	cityBtn    [][]tb.InlineButton
	stationBtn [][]tb.InlineButton
}

var cityMapping = map[string][]string{
	"基隆市": {
		"三坑",
		"八堵",
		"七堵",
		"百福",
		"海科館",
		"暖暖",
	},
	"新北市": {
		"五堵,",
		"汐止",
		"汐科",
		"板橋",
		"浮洲",
		"樹林",
		"南樹林",
		"山佳",
		"鶯歌",
		"福隆",
		"貢寮",
		"雙溪",
		"牡丹",
		"三貂嶺,",
		"大華",
		"十分",
		"望古",
		"嶺腳",
		"平溪",
		"菁桐",
		"猴硐",
		"瑞芳",
		"八斗子",
		"四腳亭",
	},
	"臺北市": {
		"南港",
		"松山",
		"臺北",
		"臺北-環島",
		"萬華",
	},
	//"桃園市": {},
	//"新竹縣": {},
	//"新竹市": {},
	//"苗栗縣": {},
	//"臺中市": {},
	//"彰化縣": {},
	//"南投縣": {},
	//"雲林縣": {},
	//"嘉義縣": {},
	//"嘉義市": {},
	//"臺南市": {},
	//"高雄市": {},
	//"屏東縣": {},
	//"臺東縣": {},
	//"花蓮縣": {},
	//"宜蘭縣": {},
}

//var cityScationList = []cityStation
//
//func initStation() {
//	//cityStation := cityStation{}
//	for cityName, station := range cityMapping {
//		stationBtn := make([][][]tb.InlineButton, len(station))
//		for index, stationName := range station {
//			stationBtn[index] = [][]tb.InlineButton{{{Text: stationName}}}
//		}
//		cityStation := cityStation{
//			cityBtn:    [][]tb.InlineButton{{{Text: cityName}}},
//			stationBtn: stationBtn,
//		}
//
//		//cityStation.cityBtn
//	}
//
//}
