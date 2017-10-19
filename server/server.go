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
    "math"
    "fmt"
)

const (
    port = ":50051"
)

// server is used to implement pb.ClockProviderServer
type clockProviderServer struct {

}

func (server *clockProviderServer) GetTime(ctx context.Context, in *pb.Auth) (*pb.Time, error) {
    ticker.Wait(100 * time.Millisecond)
    t := time.Now()
    log.Printf("get time: %s", timefmt.FmtNano(t))
    ticker.Wait(100 * time.Millisecond)
    return &pb.Time{Timestamp: t.UnixNano()}, nil
}

func main() {


    go func () {
        toWait := int(math.Pow10(9)) - time.Now().Nanosecond()
        ticker.Wait(time.Duration(toWait))

        ticker := time.NewTicker(time.Second)
        for t := range ticker.C {
            fmt.Println(timefmt.FmtNano(t))
        }
    } ()

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
