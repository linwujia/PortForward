package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var listen = &cli.Command{
	Name:  "listen",
	Usage: "Forward data from the server to the listen address(转发服务端的数据到目标地址)",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "address",
			Usage:   "The IP address or hostname of the listen, such as 127.0.0.1 or www.baidu.com(目标的IP地址或者主机名， 例如127.0.0.1或者www.baidu.com)",
			Value:   "0.0.0.0",
			Aliases: []string{"addr"},
		},

		&cli.StringFlag{
			Name:    "port",
			Usage:   "Target port address(目标的端口地址)",
			Value:   "8080",
			Aliases: []string{"p"},
		},

		&cli.StringFlag{
			Name:    "forwardPort",
			Usage:   "Forwarded to local port(转发到本地的端口)",
			Aliases: []string{"f_Port"},
			Value:   "443",
		},
	},

	Action: func(context *cli.Context) error {
		address := context.String("address")
		port := context.String("port")
		forwardPort := context.String("forwardPort")

		client := NewClient(address, port, forwardPort)
		client.Accept()
		fmt.Printf("listen %s:%s", address, port)
		return nil
	},
}
