package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {


	srv := grpc.NewServer()
	server := &GRPCServer{}
	RegisterAServer(srv, server)

	l, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal(err)
	}

	if err := srv.Serve(l); err != nil{
		log.Fatal(err)
	}


}