package main

import (
    "log"
    "net"
    "time"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"

    pb "github.com/nettee/tickgo/tick"
    "github.com/nettee/tickgo/ticker"
)

const (
    port = ":50051"
)

//var timeStandard = time.Date(2017,10,1,0,0,0,0,time.UTC)

// server is used to implement pb.ClockProviderServer
type clockProviderServer struct {

}

func (server *clockProviderServer) GetTime(ctx context.Context, in *pb.Auth) (*pb.Time, error) {
    timestamp := time.Now().UnixNano();
    log.Printf("Nanoseconds since 1970/1/1 00:00:00 UTC: %d", timestamp)
    return &pb.Time{Timestamp: timestamp}, nil
}

func main() {

    go ticker.Tick(time.Now().Nanosecond())

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
