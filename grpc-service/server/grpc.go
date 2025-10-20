package server

import (
	"fmt"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	grpchealthv1 "google.golang.org/grpc/health/grpc_health_v1"
	"grpc-service/api"
	"grpc-service/internal/handler"
	"log"
	"net"
	"time"
)

type Server struct {
	host   string
	port   string
	server *grpc.Server
	health *health.Server
}

func NewServer(host, port string) *Server {
	s := Server{host: host, port: port}
	s.server = grpc.NewServer(grpcServerOption()...)
	s.health = health.NewServer()

	return &s
}

func (s *Server) RegisterGrpCServices() {
	grpchealthv1.RegisterHealthServer(s.server, s.health)
	api.RegisterUploadServer(s.server, &handler.UploadHandler{})
}

func (s *Server) Start() {
	address := fmt.Sprintf("%s:%s", s.host, s.port)

	listen, netListenerErr := net.Listen("tcp", address)
	if netListenerErr != nil {
		err := netListenerErr
		log.Fatalf("[grpc] init failure: %s", err)
	}

	s.health.SetServingStatus("", grpchealthv1.HealthCheckResponse_SERVING)

	if err := s.server.Serve(listen); err != nil {
		log.Fatalf("[grpc] start failure: %s", err)
	}
}

func (s *Server) Shutdown() {
	// mark as not_serving to stop new traffic
	s.health.SetServingStatus("", grpchealthv1.HealthCheckResponse_NOT_SERVING)

	// wait for existing requests to complete
	time.Sleep(10 * time.Second)

	fmt.Println("[grpc] shutdown!")
	s.server.GracefulStop()
}

// HELPERS

func grpcServerOption() (opt []grpc.ServerOption) {
	opt = append(opt, grpc.ChainUnaryInterceptor(
		grpcrecovery.UnaryServerInterceptor(),
	))

	opt = append(opt, grpc.ChainStreamInterceptor(
		grpcrecovery.StreamServerInterceptor(),
	))

	opt = append(opt,
		grpc.MaxRecvMsgSize(1*1024*1024), // Set maximum receive message size up to 1MB
		grpc.MaxSendMsgSize(1*1024*1024), // Set maximum send message size up to 1MB
	)

	return
}
