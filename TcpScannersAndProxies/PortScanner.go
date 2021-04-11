package main

import (
    "fmt"
    "net"
    "sync"
)

func main() {
/////////////////////////////////////////////////////////////////////////////
    // The first argument states the kind of connection to initiate.
    // Other options involve: UDP, Unix Sockets, and custom non standard 
    // layer 4 protocols.
    // The second argument is the host and port to which to connect.
    // In ipv4/tcp connections, this will be something like 'host:port'
    //_, err := net.Dial("tcp", "scanme.nmap.org:80")

    // We can know if the connection has been established via error checking.
    // if err == nil {
    //     fmt.Println("Connection successful") 
    // }
//////////////////////////////////////////////////////////////////////////////
    // Nonconcurrent Scanning
    //for i:=1; i <= 1024; i++ {
    //    // There are two ways to convert integers to strings:
    //    //  a. strconv
    //    //  b. Sprintf(format strin, a ...interface{})
    //    address := fmt.Sprintf("scanme.nmap.org:%d", i)
    //    conn, err := net.Dial("tcp", address)
    //    if err != nil {
    //        // port is closed or filtered
    //        continue
    //    }
    //    conn.Close()
    //    fmt.Printf("%d open\n", i)
    //}
/////////////////////////////////////////////////////////////////////////////
    // Performing Concurrent Scanning
    // Naive implementation will exit inmediately.
    // We need a WaitGroup from the sync package to control the concurrency.
    // This acts as a synchronization counter that is incremented each time a 
    // go routine is going to be created.
    // However, this version of the program is still too limited, if we run it
    // against multiple hosts, we will get inconsistent results.
    // var waitGroup sync.WaitGroup
    // for i := 1; i <= 1024; i++ {
    //     waitGroup.Add(1)
    //     go func(j int) {
    //         defer waitGroup.Done()
    //         address := fmt.Sprintf("scanme.nmap.org:%d", j)
    //         conn, err := net.Dial("tcp", address)
    //         if err != nil {
    //             return
    //         }
    //         conn.Close()
    //         fmt.Printf("%d open\n", j)
    //     }(i)
    // }
    // waitGroup.Wait()
///////////////////////////////////////////////////////////////////////////////
    // Port Scanning Using a Worker Pool : TODO
}
