package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "client",
		Usage: "Port Forward Client,used to map the intranet port to the public network, provided that you must deploy the server on the public network(端口映射客户端, 用于将内网端口映射到公网上, 前提是你必须在公网上部署服务端)",
		Commands: []*cli.Command{
			listen,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
