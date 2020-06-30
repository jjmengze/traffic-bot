package handler

type Handler interface {
	Run(stopCh <-chan struct{})
}
