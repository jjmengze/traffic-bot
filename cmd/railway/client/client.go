package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	tra "traffic-bot/pkg/apis/tra/v1alpha1"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile     = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	timeout    = flag.Int64("timeout", 2, "The client request server timeout xtime,default 2 second")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := tra.NewSearchClient(conn)
	printScheduleTime(client,
		&tra.SearchTrainRequest{
			Csrf:           "d6cecf75-c5c9-4027-8ab9e-36e8f1dc89a1",
			StartStation:   "7380-四腳亭",
			EndStation:     "1000-臺北",
			Transfer:       "ONE",
			RideDate:       "2020/06/27",
			StartOrEndTime: "true",
			StartTime:      "00:00",
			EndTime:        "23:59",
			TrainTypeList:  "ALL",
			Query:          "",
		})
	printCityCode(client, &tra.Empty{})
}

// printFeature gets the feature for the given point.
func printScheduleTime(client tra.SearchClient, req *tra.SearchTrainRequest) {
	//log.Printf("Getting feature for point (%d, %d)", point.Latitude, point.Longitude)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	response, err := client.SearchTrain(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
	}
	log.Println(response.Message)
}

func printCityCode(client tra.SearchClient, req *tra.Empty) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	response, err := client.GetCity(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
	}
	for _, value := range response.Results {
		fmt.Printf("name:%s code:%s \n", value.Name, value.Code)
	}
}
