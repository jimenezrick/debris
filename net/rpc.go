package net

import (
	"net"
	"net/rpc"
	"strconv"
	"log"
)

const RpcPort = 10000

func startRpc() net.Listener {
	lis, err := net.Listen("tcp", net.JoinHostPort("", strconv.Itoa(RpcPort)))
	if err != nil {
		log.Fatal("rpc.Start:", err)
	}

	go rpc.Accept(lis)
	return lis
}
