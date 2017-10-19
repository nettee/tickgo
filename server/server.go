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
    "github.com/nettee/tickgo/timefmt"
)

const (
    port = ":50051"
)

// server is used to implement pb.ClockProviderServer
type clockProviderServer struct {

}

func (server *clockProviderServer) GetTime(ctx context.Context, in *pb.Auth) (*pb.Time, error) {
    t := time.Now()
    log.Printf("get time: %s", timefmt.Fmt(t))
    return &pb.Time{Timestamp: t.UnixNano()}, nil
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
