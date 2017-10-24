package main

import "net/rpc"
import (
	"flag"
	"log"
	"net"

	"github.com/raven4752/gtimer"
)

func main() {
	log.Print("server start")

	token := flag.String("token", "gouliguojiashengsiyi", "token used to authentication")
	port := flag.String("port", "12345", "token used to authentication")

	flag.Parse()
	t := new(gtimer.RemoteTimer)
	t.NewTimer(*token)
	rpc.Register(t)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", "0.0.0.0:"+*port)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("error accept", err.Error())
		}
		go rpc.ServeConn(conn)
	}
}
