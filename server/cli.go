package main

import (
	"github.com/urfave/cli/v2"
)

var listen = &cli.Command{
	Name:    "listen",
	Usage:   "(Listening port for connecting to clients) 监听端口，用于给客户端连接",
	Aliases: []string{"l"},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "address",
			Usage:   "The listening IP address or hostname, such as 127.0.0.1 or www.baidu.com(监听的IP地址或者主机名， 例如127.0.0.1或者www.baidu.com)",
			Aliases: []string{"addr"},
			Value:   "0.0.0.0",
		},

		&cli.StringFlag{
			Name:    "port",
			Usage:   "Listening port address(监听的端口地址)",
			Value:   "8080",
			Aliases: []string{"p"},
		},

		&cli.StringFlag{
			Name:    "targetAddress",
			Usage:   "The IP address or hostname of the target, such as 127.0.0.1 or www.baidu.com(目标的IP地址或者主机名， 例如127.0.0.1或者www.baidu.com)",
			Value:   "0.0.0.0",
			Aliases: []string{"tAddr"},
		},
		&cli.StringFlag{
			Name:    "targetPort",
			Usage:   "Target port address(目标的端口地址)",
			Value:   "8080",
			Aliases: []string{"tPort"},
		},
	},

	Action: func(context *cli.Context) error {
		address := context.String("address")
		port := context.String("port")
		targetAddress := context.String("targetAddress")
		targetPort := context.String("targetPort")
		listener := NewListener(address, port, targetAddress, targetPort)
		listener.Accept()
		return nil
	},
}
