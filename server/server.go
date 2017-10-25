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
    "errors"
    "os"
)

var auths = map[string]string {
    "user1": "pass1",
    "user2": "pass2",
    "user3": "pass3",
    "user4": "pass4",
    "user5": "pass5",
    "user6": "pass6",
}

// server is used to implement pb.ClockProviderServer
type clockProviderServer struct {

}

func (server *clockProviderServer) GetTime(ctx context.Context, in *pb.Auth) (*pb.Time, error) {

    p, ok := auths[in.Username]
    if !ok {
        log.Printf("username not exist, username: `%s', password: `%s'", in.Username, in.Password)
        return nil, errors.New("username not exist")
    }
    if p != in.Password {
        log.Printf("wrong password, username: `%s', password: `%s'", in.Username, in.Password)
        return nil, errors.New("wrong password")
    }
    log.Printf("auth passed, username: `%s', password: `%s'", in.Username, in.Password)

    ticker.Wait(100 * time.Millisecond)
    t := time.Now()
    log.Printf("get time: %s", timefmt.FmtNano(t))
    ticker.Wait(100 * time.Millisecond)
    return &pb.Time{Timestamp: t.UnixNano()}, nil
}

func main() {

    if len(os.Args) < 1 {
        log.Fatalf("Error. Listening port not provided.")
    }
    port := os.Args[1]

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
