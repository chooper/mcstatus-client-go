
package main

import (
    "flag"
    "log"
    "os"
    "github.com/chooper/mcstatus-client-go/poller"
)

func usage() {
    log.Printf("Usage: %s poller")
}

func launch_poller() {
    var server string
    server = os.Getenv("MINECRAFT_SERVER")

    c := make(chan poller.Notification)
    p := &poller.Poller{
        Server:         server,
        NotifyChan:     c,
    }
    go p.Loop()
    for {
        select {
            case n := <- c:
                log.Printf("N: %v", n)
        }
    }
}

func main() {
    flag.Parse()

    if len(os.Args) < 2 {
        usage()
        log.Fatal("See usage")
    }

    if os.Args[1] == "poller" {
        launch_poller()
    }
}

