package main

import (
	"http-service/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	grpcConn := server.GrpcConn()

	client := server.NewServer(grpcConn, "0.0.0.0", "8080")
	client.SetRoutes()
	client.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	server.GrpcClose(grpcConn)
	client.Shutdown()
}
