package main

import (
    "log"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
    pb "github.com/nettee/tickgo/tick"
    "time"
    "fmt"
    "github.com/nettee/tickgo/ticker"
)

const (
    address = "localhost:50051"
)

//var timeStandard = time.Date(2017,10,1,0,0,0,0,time.UTC)

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
    fmt.Println(serverTime.Format(time.RFC3339Nano))

    ticker.Tick(serverTime.Nanosecond())
}
