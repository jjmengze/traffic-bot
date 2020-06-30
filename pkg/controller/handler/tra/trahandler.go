package tra

import (
	"fmt"
	"k8s.io/klog/v2"
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
		for {
			select {
			case event := <-e.middle.EventCh:
				fmt.Printf("recice %s", event)
			case <-innerCh:
				klog.Info("received a Interrupt signal ")
				break
			}
		}
	}()

	<-stopCh
}
