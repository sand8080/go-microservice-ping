package main

import (
	"github.com/sand8080/go-microservice-ping/service"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Ping micro-service"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "address, a",
			Usage: "Listen address",
			Value: "0.0.0.0",
		},
		cli.IntFlag{
			Name:  "port, p",
			Usage: "Service port",
			Value: 8080,
		},
	}
	app.Action = service.Serve

	app.Run(os.Args)
}
