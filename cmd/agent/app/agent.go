package app

import (
	"net"
	"net/http"
	"sync/atomic"

	logger "github.com/Zelayan/dts/pkg/log"
)

type Agent struct {
	httpServer *http.Server
	httpAddr atomic.Value
	logger logger.LoggerInterface
}

func NewAgent(log logger.LoggerInterface) *Agent {
	return &Agent{
		httpServer: &http.Server{Addr: "0.0.0.0:4000"},
		logger: log,
	}
}


func (a *Agent) Run() error {
	listener, err := net.Listen("tcp", a.httpServer.Addr)
	if err != nil {
		return err
	}
	a.httpAddr.Store(listener.Addr().String())
	go func ()  {
		a.logger.InfoF("Start dts-agent HTTP Server， Addr: %s", listener.Addr().String())
		if err = a.httpServer.Serve(listener); err != http.ErrServerClosed {
			a.logger.Error("http server failed", err)
		}
	}()
	return nil

}