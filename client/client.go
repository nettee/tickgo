package main

import (
    "fmt"
    "log"
    "time"

    "golang.org/x/net/context"
    "google.golang.org/grpc"

    pb "github.com/nettee/tickgo/tick"
    "github.com/nettee/tickgo/ticker"
    "github.com/nettee/tickgo/timefmt"
)

const (
    address = "localhost:50051"
)

func main() {
    // Set up a connection to the server
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    client := pb.NewClockProviderClient(conn)

    // Contact the server and print out its response.
    res, err := client.GetTime(context.Background(), &pb.Auth{Username: "user", Password: "pass"})
    if err != nil {
        log.Fatalf("could not contact the server: %v", err)
    }
    serverTime := time.Unix(0, res.Timestamp)
    fmt.Printf("server time: %s\n", timefmt.Fmt(serverTime))

    ticker.Tick(serverTime.Nanosecond())
}
