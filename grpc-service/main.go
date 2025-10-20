package main

import (
	"fmt"
	"grpc-service/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s := server.NewServer("0.0.0.0", "8082")
	s.RegisterGrpCServices()

	go s.Start()

	fmt.Println("[grpc] started on port 8082")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	s.Shutdown()
}
