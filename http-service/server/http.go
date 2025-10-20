package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"http-service/internal/handler"
	"log"
	"net/http"
	"time"
)

type Server struct {
	host   string
	port   string
	router *mux.Router
	grpc   *grpc.ClientConn
	client *http.Server
}

func NewServer(client *grpc.ClientConn, host, port string) *Server {
	s := Server{grpc: client, host: host, port: port}
	s.router = mux.NewRouter()

	return &s
}

func (s *Server) SetRoutes() {
	s.router.HandleFunc("/handshake", handler.Handshake).Methods(http.MethodGet)
	s.router.HandleFunc("/upload", handler.Upload(s.grpc)).Methods(http.MethodPost)
}

func (s *Server) Start() {
	s.client = &http.Server{
		Handler:      s.router,
		Addr:         fmt.Sprintf("%s:%s", s.host, s.port),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	go func() {
		err := s.client.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("[http] error: %v", err)
		}
	}()

	fmt.Println("[http] started!")
}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.client.Shutdown(ctx); err != nil {
		log.Fatalf("[http] shutdown error: %v", err)
	}

	fmt.Println("[http] shutdown!")
}
