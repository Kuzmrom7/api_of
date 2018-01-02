package main

import (
	"github.com/orderfood/api_of/pkg/api"
	"github.com/orderfood/api_of/pkg/api/config"
	"github.com/jawher/mow.cli"
	"fmt"
	"os"
	"github.com/orderfood/api_of/pkg/log"
)

const (
	host = "localhost"
	port = 8080
)

func main() {

	var (
		cfg config.Config
	)

	app := cli.App("", "Manager api")

	app.Version("v version", "0.1.0")

	app.Spec = "[OPTIONS]"

	cfg.APIServer.Host = app.String(cli.StringOpt{
		Name:   "http-server-host", Desc: "Http server host",
		EnvVar: "HTTP_SERVER_HOST", Value: host, HideValue: true,
	})

	cfg.APIServer.Port = app.Int(cli.IntOpt{
		Name:   "http-server-port", Desc: "Http server port",
		EnvVar: "HTTP_SERVER_PORT", Value: port, HideValue: true,
	})

	cfg.Database.Connection = app.String(cli.StringOpt{
		Name:   "pgsql-connection", Desc: "Set postgres connection string",
		EnvVar: "PGSQL_CONNECTION", Value: "host= 0.0.0.0 port=5432 user=orderfood  password=orderfood dbname=orderfood sslmode=disable", HideValue: true,
	})

	var help = app.Bool(cli.BoolOpt{
		Name:      "h help",
		Value:     false,
		Desc:      "Show the help info and exit",
		HideValue: true,
	})

	app.Before = func() {
		if *help {
			app.PrintLongHelp()
		}
		log.New(true)
	}

	app.Action = func() {
		api.Daemon(&cfg)
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Errorf("Error: run application: %s", err.Error())
		return
	}

}
