package main

import (
    "fmt"
    "log"
    "os"
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

    if len(os.Args) < 2 {
        log.Fatalf("Error. Username and password not provided.")
    }
    username := os.Args[1]
    password := os.Args[2]

    // Contact the server and print out its response.
    t1 := time.Now()
    log.Printf("get time from server at %s", timefmt.FmtNano(t1))
    res, err := client.GetTime(context.Background(), &pb.Auth{Username: username, Password: password})
    if err != nil {
        log.Fatalf("Failed to connect to the server: %v", err)
    }
    t2 := time.Now()
    serverTime := time.Unix(0, res.Timestamp)
    log.Printf("server time: %s", timefmt.FmtNano(serverTime))
    log.Printf("got server time at %s", timefmt.FmtNano(t2))
    rtt := t2.Sub(t1)
    log.Printf("RTT: %d milliseconds", rtt.Nanoseconds() / int64(math.Pow10(6)))

    realServerTime := serverTime.Add(rtt / 2)
    log.Printf("real server time: %s", timefmt.FmtNano(realServerTime))

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


