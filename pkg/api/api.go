package api

import (
	"github.com/orderfood/api_of/pkg/api/http"
	"github.com/orderfood/api_of/pkg/api/config"
	"github.com/orderfood/api_of/pkg/api/context"
	"github.com/orderfood/api_of/pkg/storage"
	"os/signal"
	"syscall"
	"os"
	"github.com/orderfood/api_of/pkg/log"
)

func Daemon(_cfg *config.Config) {
	var (
		ctx  = context.Get()
		cfg  = config.Set(_cfg)
		done = make(chan bool, 1)
		sigs = make(chan os.Signal)
	)

	ctx.SetConfig(cfg)
	stg, err := storage.Get(cfg.GetPGDB())
	if err != nil {
		log.Errorf("Cannot initialize storage: %v",err)
	}

	ctx.SetStorage(stg)

	go func() {
		http.Listen(*cfg.APIServer.Host, *cfg.APIServer.Port)
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
