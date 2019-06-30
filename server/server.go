package main

import (
	"bytes"
	"log"
	"net"
	"os"

	"github.com/obsilp/rmnp"
)

var debug = log.New(os.Stdout, "\x1b[36m[DEBU]\x1b[0m ", log.Ldate|log.Lmicroseconds|log.Lshortfile)

var (
	PING = []byte("ping")
	PONG = []byte("pong")
)

func main() {
	server := rmnp.NewServer(":10001")

	server.ClientConnect = clientConnect
	// server.ClientDisconnect = clientDisconnect
	server.ClientTimeout = clientTimeout
	server.ClientValidation = validateClient
	server.PacketHandler = handleServerPacket

	debug.Println("server started")
	server.Start()
	select {}
}

func clientConnect(conn *rmnp.Connection, data []byte) {
	debug.Println("client connection with:", data)

	if data[0] != 0 {
		conn.Disconnect([]byte("not allowed"))
	}
}

func clientDisconnect(conn *rmnp.Connection, data []byte) {
	debug.Println("client disconnect")
}

func clientTimeout(conn *rmnp.Connection, data []byte) {
	debug.Println("client timeout")
}

func validateClient(addr *net.UDPAddr, data []byte) bool {
	return len(data) == 3
}

func handleServerPacket(conn *rmnp.Connection, data []byte, channel rmnp.Channel) {
	str := string(data)
	debug.Println("'"+str+"'", "from", conn.Addr.String(), "on channel", channel)
	if bytes.Equal(data, PING) {
		conn.SendReliableOrdered(PONG)
	}
}
