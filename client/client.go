package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/obsilp/rmnp"
)

var debug = log.New(os.Stdout, "\x1b[36m[DEBU]\x1b[0m ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
var sendChannel = make(chan sendmsg, 1000)

type sendmsg struct {
	msgType uint8
	msgData []byte
}

func main() {
	client := rmnp.NewClient("127.0.0.1:10001")
	client.ConnectWithData([]byte{0, 1, 2})
	client.ServerConnect = serverConnect
	client.ServerTimeout = serverTimeout
	client.PacketHandler = handleClientPacket

	// client.ServerDisconnect = serverDisconnect
	sendText()
	select {}
}

func serverConnect(conn *rmnp.Connection, data []byte) {
	debug.Println("connected to server")
	var msg sendmsg
	for {
		msg = <-sendChannel
		switch msg.msgType {
		case 0:
			conn.SendUnreliable(msg.msgData)
		case 1:
		case 2:
		case 3:
			conn.SendReliableOrdered(msg.msgData)
		}
	}
}
func serverDisconnect(conn *rmnp.Connection, data []byte) {
	fmt.Println("disconnected from server:", string(data))
}

func serverTimeout(conn *rmnp.Connection, data []byte) {
	fmt.Println("server timeout")
}

func handleClientPacket(conn *rmnp.Connection, data []byte, channel rmnp.Channel) {
	fmt.Println("'"+string(data)+"'", "on channel", channel)
}

func sendText() {
	var newMsg = &sendmsg{}
	for {
		newMsg.msgType = 0
		newMsg.msgData = []byte("PING")
		sendChannel <- *newMsg
		time.Sleep(60e6)
	}
}
