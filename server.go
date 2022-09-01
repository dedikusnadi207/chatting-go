package main

import (
	"belajar-chatting-grpc/chatserver"
	"belajar-chatting-grpc/utils"
	"log"
	"net"

	"google.golang.org/grpc"
)

func runServer(port string) {
	port = utils.AvailablePort("localhost", port)
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Cannot listen @%v::%v", port, err)
	}
	log.Println("Running server on @:" + port)

	grpcServer := grpc.NewServer()

	// register ChatService
	cs := chatserver.ChatServer{}
	chatserver.RegisterServicesServer(grpcServer, &cs)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start gRPC Server :: %v", err)
	}
}
