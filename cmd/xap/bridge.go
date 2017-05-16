package main

import (
	"net"

	"github.com/urfave/cli"
	"rbg.re/robertgzr/xapper/pkg/socketbridge"
)

func BridgeSubcommand() *cli.Command {
	return &cli.Command{
		Name:      "bridge",
		Aliases: []string{"b"},
		Usage:     "create a socket bridge to connect to the mpv unix socket via tcp",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "port, p",
				Usage: "tcp port for the bridge",
				Value: "8914",
			},
		},
		Action: spawnBridge,
	}
}

func spawnBridge(ctx *cli.Context) error {

	mpvSock, err := net.Dial("unix", ctx.String("socket"))
	if err != nil {
		return err
	}
	defer mpvSock.Close()

	tcpLn, err := net.Listen("tcp", ":"+ctx.String("port"))
	if err != nil {
		return err
	}
	defer tcpLn.Close()

	for {
		tcpConn, err := tcpLn.Accept()
		if err != nil {
			panic(err)
		}
		go socketbridge.BidirectionalBridge(mpvSock, tcpConn)
	}
}
