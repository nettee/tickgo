package main

import (
    "fmt"
    "log"
    "time"

    "golang.org/x/net/context"
    "google.golang.org/grpc"

    pb "github.com/nettee/tickgo/tick"
    "github.com/nettee/tickgo/timefmt"
    "math"
    "github.com/nettee/tickgo/ticker"
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
    t1 := time.Now()
    fmt.Printf("get time from server at %s\n", timefmt.FmtNano(t1))
    res, err := client.GetTime(context.Background(), &pb.Auth{Username: "user", Password: "pass"})
    if err != nil {
        log.Fatalf("could not contact the server: %v", err)
    }
    t2 := time.Now()
    serverTime := time.Unix(0, res.Timestamp)
    fmt.Printf("server time: %s\n", timefmt.FmtNano(serverTime))
    fmt.Printf("got server time at %s\n", timefmt.FmtNano(t2))
    rtt := t2.Sub(t1)
    fmt.Printf("RTT: %d milliseconds\n", rtt.Nanoseconds() / int64(math.Pow10(6)))

    realServerTime := serverTime.Add(rtt / 2)
    fmt.Printf("real server time: %s\n", timefmt.FmtNano(realServerTime))

    tickExpected(realServerTime)
}

func tickExpected(t time.Time) {
    toWait := time.Duration(int(math.Pow10(9)) - t.Nanosecond())
    ticker.Wait(toWait)

    st := t.Add(toWait)

    ti := time.NewTicker(time.Second)
    for _ = range ti.C {
        st = st.Add(time.Second)
        fmt.Println(timefmt.Fmt(st))
    }

}


