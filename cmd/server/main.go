package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"os"

	"github.com/toddsifleet/godo/index"
	"github.com/toddsifleet/godo/server"
	"github.com/toddsifleet/godo/tail"
)

// TODO: MOVE TO CONFIG
const HOST = ":1234"

func main() {
	// CREATE STRUCTURE TO STORE DATA
	idx := index.New()

	// START MONITORING LOG FILE
	t, err := tail.New(idx, os.Args[1])
	if err != nil {
		panic(err)
	}
	go t.Run()

	// CREATE SERVER TO RESPOND TO REQUESTS
	s, err := server.New(idx)
	if err != nil {
		panic(err)
	}

	// START LISTENING FOR REQUESTS
	rpc.Register(s)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", HOST)
	if err != nil {
		panic(err)
	}
	fmt.Println("LISTENING")
	defer listener.Close()

	// RUN FOREVER
	http.Serve(listener, nil)
}
