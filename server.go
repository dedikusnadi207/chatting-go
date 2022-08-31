package main

import (
	"belajar-chatting-grpc/chatserver"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
)

func runServer() {
	port := *flag.String("p", "8899", "")
	flag.Parse()
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Cannot listen @%v::%v", port, err)
	}
	log.Println("Listening @:", port)

	grpcServer := grpc.NewServer()

	// register ChatService
	cs := chatserver.ChatServer{}
	chatserver.RegisterServicesServer(grpcServer, &cs)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start gRPC Server :: %v", err)
	}
}
