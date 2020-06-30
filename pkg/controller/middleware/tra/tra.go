package tra

import (
	"context"
	"fmt"
	"k8s.io/klog/v2"
	tra "traffic-bot/pkg/apis/tra/v1alpha1"
)

//user send a req to bot and proxy to backend.
func GetCity(c context.Context, client tra.SearchClient) {
	response, err := client.GetCity(c, &tra.Empty{})
	if err != nil {
		klog.Error("rpc call GetCity error :", client, err)
	}
	for _, value := range response.Results {
		fmt.Printf("name:%s code:%s \n", value.Name, value.Code)
	}
	//for _, result := range stationData.Results {
	//	cityCode := result.CityCode
	//	for _, station := range result.StationList {
	//		err := insertStation(db, &Station{
	//			CityCode:    cityCode,
	//			Name:        station.Name,
	//			StationCode: station.Code,
	//		})
	//		if err != nil {
	//			fmt.Errorf("error:%s", err)
	//		}
	//	}
	//}
}
