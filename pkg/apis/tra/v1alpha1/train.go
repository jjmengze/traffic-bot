package v1alpha1

import (
	"context"
	"fmt"
	"github.com/gocolly/colly/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var collect *colly.Collector

const (
	SEARCHURL = "https://www.railway.gov.tw/tra-tip-web/tip/tip001/tip112/querybytime"
)

func init() {
	collect = colly.NewCollector()
}

type SearchService struct{}

func (s *SearchService) SearchTrain(ctx context.Context, r *SearchTrainRequest) (*SearchTrainResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "searchService.Search canceled")
	}
	result := make([]*Result, 0, 0)
	// Find and visit all links
	collect.OnHTML(".itinerary-controls", func(e *colly.HTMLElement) {

		e.ForEach(".trip-column", func(i int, element *colly.HTMLElement) {
			//fmt.Printf("index : %d ,vale %s \n", i, element.Text)
			//fmt.Println(element.ChildText(".train-number"))
			data := element.ChildTexts("td")
			dataKey := []string{"Url", "StartTime", "EndTime", "Spend", "Ticket"}
			mapping := make(map[string]string)

			for idx, v := range data {
				if idx == 6 {
					mapping[dataKey[4]] = v
				} else if idx < len(dataKey) {
					mapping[dataKey[idx]] = v
				}
			}
			r := &Result{
				Url:       mapping["Url"],
				StartTime: mapping["StartTime"],
				EndTime:   mapping["EndTime"],
				Spend:     mapping["Spend"],
				Ticket:    mapping["Ticket"],
			}
			result = append(result, r)
		})
	})

	body := map[string]string{
		"_csrf":          r.Csrf,
		"startStation":   r.StartStation,
		"endStation":     r.EndStation,
		"transfer":       r.Transfer,
		"rideDate":       r.RideDate,
		"startOrEndTime": r.StartOrEndTime,
		"startTime":      r.StartTime,
		"endTime":        r.EndTime,
		"trainTypeList":  r.TrainTypeList,
		"query":          "查詢",
	}
	collect.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	err := collect.Post(SEARCHURL, body)
	return &SearchTrainResponse{Message: result}, err
}

func NewSearchService() *SearchService {
	//s := &routeGuideServer{routeNotes: make(map[string][]*pb.RouteNote)}
	//s.loadFeatures(*jsonDBFile)
	return &SearchService{}
}
