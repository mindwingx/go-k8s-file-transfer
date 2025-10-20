package server
	
import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	grpchealthv1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"log"
	"os"
	"time"
)

func GrpcConn() *grpc.ClientConn {
	grpcHost := os.Getenv("GRPC_HOST")
	grpcPort := os.Getenv("GRPC_PORT")

	if grpcPort == "" {
		log.Fatal("[grpc] invalid port")
	}

	target := fmt.Sprintf("%s:%s", grpcHost, grpcPort)

	client, err := grpc.NewClient(
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`), // Force round-robin
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			PermitWithoutStream: true,
			Time:                10 * time.Minute,
		}),
	)

	if err != nil {
		log.Fatalf("[grpc] connection failure : %v", err)
	}

	grpcHealthCheck(client)

	fmt.Println("[grpc] connected!")

	return client
}

func GrpcClose(conn *grpc.ClientConn) {
	if err := conn.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("[grpc] closed!")
}

// HELPERS

func grpcHealthCheck(conn *grpc.ClientConn) {
	client := grpchealthv1.NewHealthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.Check(ctx, &grpchealthv1.HealthCheckRequest{
		Service: "", // Leave empty for overall server health
	})

	if err != nil {
		log.Fatalf("[grpc] health check failed: %v", err.Error())
	}

	if response.Status != grpchealthv1.HealthCheckResponse_SERVING {
		log.Fatal("[grpc] health check status is not serving")
	}
}
