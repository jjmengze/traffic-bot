package tra

import (
	"fmt"
	"k8s.io/klog/v2"
	"traffic-bot/pkg/controller/handler/tra/actiontype"
	"traffic-bot/pkg/controller/middleware"
)

type TRAHandler struct {
	handlerContext *middleware.HandlerContext
	middle         *middleware.Middle
}

func NewTRAHandler(handlerContext *middleware.HandlerContext, middle *middleware.Middle) *TRAHandler {
	return &TRAHandler{
		handlerContext: handlerContext,
		middle:         middle,
	}
}

func (e *TRAHandler) Run(stopCh <-chan struct{}) {
	innerCh := make(chan struct{})
	klog.Infof("Starting TRA handler ")
	defer klog.Infof("Shutting down TRA handler ")
	defer close(innerCh)

	go func() {
	out:
		for {
			select {
			case event := <-e.middle.EventCh:
				e.eventRoute(&event)
				fmt.Printf("recice %s", event)
			case <-innerCh:
				klog.Info("received a Interrupt signal ")
				break out
			}
		}
	}()
	<-stopCh
}

func (e *TRAHandler) eventRoute(event *actiontype.EventInfo) {
	if event.Action == actiontype.QUERY {

		switch event.Type {
		case actiontype.City:
		}
	}
}
