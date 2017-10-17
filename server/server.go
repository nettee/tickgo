package main

import (
    "log"
    "net"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
    pb "github.com/nettee/tickgo/tick"
    "google.golang.org/grpc/reflection"
)

const (
    port = ":50051"
)

// server is used to implement pb.ClockProviderServer
type clockProviderServer struct {

}

func (server *clockProviderServer) GetTime(ctx context.Context, in *pb.Auth) (*pb.Time, error) {
    return &pb.Time{Timestamp: 233}, nil
}

func main() {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    server := grpc.NewServer()
    pb.RegisterClockProviderServer(server, &clockProviderServer{})
    // Register reflection service on gRPC server.
    reflection.Register(server)
    if err := server.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
