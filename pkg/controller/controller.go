package controller

import (
	"k8s.io/klog/v2"
	"net/http"
	"traffic-bot/pkg/controller/handler/tra"
	"traffic-bot/pkg/controller/middleware"
)

//type ControllerFactory interface {
//	Start(stopCh <-chan struct{})
//}

type InitFunc func(handlerContext *middleware.HandlerContext, handlerName middleware.EventType) (debuggingHandler http.Handler, enabled bool, err error)

//type Controller struct {
//	lock           sync.Mutex
//	handler        map[string]handler.Handler
//	startedHandler map[string]bool
//}
//
//type ControllerOption func(*Controller) *Controller
//
//func NewController() *Controller {
//	return NewNewControllerWithOptions()
//}
//
//func NewNewControllerWithOptions(options ...ControllerOption) *Controller {
//	controller := &Controller{
//		handler:        make(map[string]handler.Handler),
//		startedHandler: make(map[string]bool),
//	}
//
//	for _, opt := range options {
//		controller = opt(controller)
//	}
//
//	return controller
//}
//
//func (c *Controller) Start(stopCh <-chan struct{}) {
//	c.lock.Lock()
//	defer c.lock.Unlock()
//	for name, h := range c.handler {
//		if !c.startedHandler[name] {
//			go h.Run(stopCh)
//			c.startedHandler[name] = true
//		}
//	}
//}

func NewHandlerInitializers() map[middleware.EventType]InitFunc {
	handlers := map[middleware.EventType]InitFunc{}
	handlers[middleware.TRA] = startTRAHandlers
	return handlers
}

func StartHandlers(ctx *middleware.HandlerContext, handlers map[middleware.EventType]InitFunc) error {
	for handlerName, initFn := range handlers {
		klog.V(1).Infof("Starting %q", handlerName)
		_, started, err := initFn(ctx, handlerName)
		if err != nil {
			klog.Errorf("Error starting %q", handlerName)
			return err
		}
		if !started {
			klog.Warningf("Skipping %q", handlerName)
			continue
		}
		klog.Infof("Started %q", handlerName)
	}
	return nil
}

func startTRAHandlers(ctx *middleware.HandlerContext, handlerName middleware.EventType) (http.Handler, bool, error) {
	go tra.NewTRAHandler(
		ctx,
		ctx.GetMiddle(handlerName),
	).Run(ctx.Stop)
	return nil, true, nil
}
