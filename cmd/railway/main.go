package main

import (
	"bytes"
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"mime/multipart"
	"net/http"
	"traffic-bot/pkg/utils/xhttp"
)

func main() {

	body := map[string]string{
		"_csrf":          "d6cecf75-c5c9-4027-8ab9e-36e8f1dc89a1",
		"startStation":   "7380-四腳亭",
		"endStation":     "1000-臺北",
		"transfer":       "ONE",
		"rideDate":       "2020/04/25",
		"startOrEndTime": "true",
		"startTime":      "00:00",
		"endTime":        "23:59",
		"trainTypeList":  "ALL",
		"query":          "查詢",
	}
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	for key, value := range body {
		if err := writer.WriteField(key, value); err != nil {
			log.Fatal(err)
		}
	}
	if err := writer.Close(); err != nil {
		log.Fatal(err)
	}

	//jsonStr := []byte(`{ "username": "auto", "password": "auto123123" }`)
	url := "https://www.railway.gov.tw/tra-tip-web/tip/tip001/tip112/querybytime"
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	_, err = xhttp.Do(req)
	//fmt.Println(string(resp))

	//data := bytes.NewReader(resp)
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML(".itinerary-controls", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))
		e.ForEach(".trip-column", func(i int, element *colly.HTMLElement) {
			//fmt.Printf("index : %d ,vale %s \n", i, element.Text)
			//fmt.Println(element.ChildText(".train-number"))
			data:=element.ChildTexts("td")
			for _, v := range data {
				fmt.Printf("idx: %s ", v)
			}
			fmt.Println()
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Post(url, body)
}
