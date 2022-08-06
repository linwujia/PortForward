package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "server",
		Usage: "Port Forward Server,used to map the intranet port to the public network, provided that you must deploy the server on the public network(端口映射服务端, 用于将内网端口映射到公网上, 将该服务部署到公网上)",
		Commands: []*cli.Command{
			listen,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
