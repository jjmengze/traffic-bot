package main

import (
	"flag"
	"fmt"
	"k8s.io/klog/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
	"traffic-bot/pkg/bot"
	"traffic-bot/pkg/controller"
	"traffic-bot/pkg/controller/middleware"
)

func init() {
	klog.InitFlags(flag.CommandLine)
}

var onlyOneSignalHandler = make(chan struct{})
var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}
var shutdownHandler chan os.Signal
var shutdownDelayDuration = time.Second * 5

func main() {
	klog.Info("Start telegram bot..")

	//建立stop chan 捕捉關機訊號
	stop := setupSignalHandler()
	//startHandler
	_, err := func(stopCh <-chan struct{}) (*middleware.HandlerContext, error) {
		handlerCtx, err := middleware.CreateHandlerContext(stopCh)
		err = controller.StartHandlers(handlerCtx, controller.NewHandlerInitializers())
		return handlerCtx, err
	}(stop)
	b := bot.NewBot(os.Getenv("BOTTOKEN"), handlerCtx)
	b.Register()
	go b.Start()

	if err != nil {
		klog.Fatalf("error starting handler: %v", err)
	}
	//c := controller.NewNewControllerWithOptions()
	//c.Start(stop)
	//收到關機訊號之後等待shutdownDelayDuration在結束main
	Run(stop)
}

func setupSignalHandler() <-chan struct{} {
	close(onlyOneSignalHandler) // panics when called twice

	shutdownHandler = make(chan os.Signal, 2)

	stop := make(chan struct{})
	signal.Notify(shutdownHandler, shutdownSignals...)
	go func() {
		<-shutdownHandler
		close(stop)
		<-shutdownHandler
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}

func Run(stopCh <-chan struct{}) {
	delayedStopCh := make(chan struct{})

	go func() {
		defer close(delayedStopCh)
		<-stopCh
		fmt.Println("hello")
		time.Sleep(shutdownDelayDuration)
	}()
	<-delayedStopCh
}
