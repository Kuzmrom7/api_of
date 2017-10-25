package api

import (
	"github.com/orderfood/api_of/pkg/api/http"
	"os/signal"
	"syscall"
	"os"
)


func Daemon(host string,port int){
	var ( sigs = make(chan os.Signal)
		  done = make(chan bool, 1)
	)

	go func() {
		http.Listen(host,port)
	}()

	// Handle SIGINT and SIGTERM.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-sigs:
				done <- true
				return
			}
		}
	}()

	<-done
}
