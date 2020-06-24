package main

import (
	"flag"
	"k8s.io/klog/v2"
	"os"
	"os/signal"
	"syscall"
	"traffic-bot/pkg/bot"
)

func init() {
	klog.InitFlags(flag.CommandLine)
}

var onlyOneSignalHandler = make(chan struct{})
var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}
var shutdownHandler chan os.Signal

func main() {
	klog.Info("Start telegram bot..")
	b := bot.NewBot(os.Getenv("BOTTOKEN"))
	b.Register()
	go b.Start()
	Run(setupSignalHandler())
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
	<-stopCh
}
