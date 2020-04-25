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

	// Find and visit all links
	collect.OnHTML(".itinerary-controls", func(e *colly.HTMLElement) {
		e.ForEach(".trip-column", func(i int, element *colly.HTMLElement) {
			//fmt.Printf("index : %d ,vale %s \n", i, element.Text)
			//fmt.Println(element.ChildText(".train-number"))
			data := element.ChildTexts("td")
			for _, v := range data {
				fmt.Printf("idx: %s ", v)
			}
			fmt.Println()
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

	return &SearchTrainResponse{Message: "value"}, collect.Post(SEARCHURL, body)
}

func NewSearchService() *SearchService {
	//s := &routeGuideServer{routeNotes: make(map[string][]*pb.RouteNote)}
	//s.loadFeatures(*jsonDBFile)
	return &SearchService{}
}
