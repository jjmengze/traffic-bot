package middleware

import (
	"google.golang.org/grpc"
	tb "gopkg.in/tucnak/telebot.v2"
	"k8s.io/klog/v2"
	tra "traffic-bot/pkg/apis/tra/v1alpha1"
)

type EventType int

const TRA EventType = iota

type HandlerContext struct {
	middleMap map[EventType]*Middle
	bot       *tb.Bot
	Stop      <-chan struct{}
}

func (hc *HandlerContext) GetMiddle(name EventType) *Middle {
	if _, ok := hc.middleMap[name]; !ok {
		klog.Warning("Not included the key : %s ", name)
		eventCh := make(chan interface{})
		hc.middleMap[name] = &Middle{EventCh: eventCh}
		return hc.middleMap[name]
	}
	return hc.middleMap[name]
}

func (hc *HandlerContext) AppendEventChan(key EventType, ch chan interface{}) {
	//if _, ok := hc.Middle.EventMap[key]; !ok {
	//	klog.Warning("Not included the key : %s ", key)
	//	hc.Middle.EventMap[key] = ch
	//	return
	//}
}

func (hc *HandlerContext) NewRPCClient(serverAddr string) (tra.SearchClient, *grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		klog.Error("fail to dial:", err)
		return nil, conn, err
	}
	client := tra.NewSearchClient(conn)
	return client, conn, err
}

func CreateHandlerContext(stop <-chan struct{}) (*HandlerContext, error) {
	ctx := &HandlerContext{
		Stop:      stop,
		middleMap: make(map[EventType]*Middle),
	}
	return ctx, nil
}

type Middle struct {
	EventCh chan interface{}
}
