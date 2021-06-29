package main


/////////////////////////////////////////////////////////////////////////////
    // The first argument states the kind of connection to initiate.
    // Other options involve: UDP, Unix Sockets, and custom non standard 
    // layer 4 protocols.
    // The second argument is the host and port to which to connect.
    // In ipv4/tcp connections, this will be something like 'host:port'
//func main() {
    //_, err := net.Dial("tcp", "scanme.nmap.org:80")

    // We can know if the connection has been established via error checking.
    //if err == nil {
    //    fmt.Println("Connection successful") 
    //}
// }
//////////////////////////////////////////////////////////////////////////////
    // Nonconcurrent Scanning
//func main() {
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
// }
/////////////////////////////////////////////////////////////////////////////
    // Performing Concurrent Scanning
    // Naive implementation will exit inmediately.
    // We need a WaitGroup from the sync package to control the concurrency.
    // This acts as a synchronization counter that is incremented each time a 
    // go routine is going to be created.
    // However, this version of the program is still too limited, if we run it
    // against multiple hosts, we will get inconsistent results.
//func main() {
    //var waitGroup sync.WaitGroup
    //for i := 1; i <= 65535; i++ {
    //    waitGroup.Add(1)
    //    go func(j int) {
    //        defer waitGroup.Done()
    //        address := fmt.Sprintf("127.0.0.1:%d", j)
    //        conn, err := net.Dial("tcp", address)
    //        if err != nil {
    //            return
    //        }
    //        conn.Close()
    //        fmt.Printf("%d open\n", j)
    //    }(i)
    //}
    //waitGroup.Wait()
// }
///////////////////////////////////////////////////////////////////////////////
    // Port Scanning Using a Worker Pool
    // - Use a pool of goroutines to manage the concurrent work being performed.
    // - Use a channel to provide work from the main thread.

//import (
//    "fmt"
//    "sync"
//)
//func worker(ports chan int, wg *sync.WaitGroup) {
//    // Loop until the channel is closed.
//    // - We use range to continously receive from the ports channel.
//    for p:= range ports {
//        fmt.Println(p)
//        wg.Done()
//    }
//}
//func main() {
//    // Create a channel using make.
//    // - The second parameter allows the channel to be buffered.
//    // - A buffered channel is one on which can be sent an item without waiting for a receiver to read it.
//    // - Buffered channels are used to maintain and track work for multiple producers and consumers.
//    // - Buffered channels are slightly more performant.
//    ports := make(chan int, 100)
//    defer close(ports)
//    var wg sync.WaitGroup
//    // Use a loop to initialzie 100 workers.
//    // - Pass them a copy of the buffered channel and a pointer to the wait group.
//    for i:=0; i<cap(ports); i++ {
//        go worker(ports, &wg)
//    }
//
//    // Iterate over all the ports and pass them to each worker of the pool.
//    for i:=1; i<= 1024; i++ {
//        wg.Add(1)
//        ports <- i
//    }
//    wg.Wait()
//
//    // After all work has been completed, the channel shall be closed.
//}
///////////////////////////////////////////////////////////////////////////////////
// Multichannel communications
// - We want to sort the open ports before printing them to the screen.
import (
    "fmt"
    "net"
    "sort"
    "time"
)

// Worker now accepts two channels, one for input and
// the other for output.
func worker(ports, results chan int) {
    for p:=range ports {
        address := fmt.Sprintf("192.168.1.1:%d", p)
        conn, err := net.DialTimeout("tcp", address, time.Second)
        // Send a 0 if the port is closed,
        // then wait for a new port to scan.
        if err != nil {
            results <- 0
            continue
        }
        conn.Close()
        results <- p
    }
}

func main() {
    // Setup a buffered channel for the orders.
    ports := make(chan int, 100)
    // Setup a channel to receive the results.
    results := make(chan int)
    // Store open ports on a slice.
    var openports []int

    for i:=0; i<cap(ports); i++ {
        go worker(ports, results)
    }

    // Producer thread: Sends ports to analyze.
    go func() {
        for i:=1; i<=1024; i++ {
            ports <- i
        }
    }()

    for i:=0; i<1024; i++ {
        // Receive results
        port := <-results
        // 0 means that the port is closed.
        // If the port is not closed, then append it to
        // the list of results.
        if port != 0 {
            openports = append(openports, port)
        }
        fmt.Print(".")
    }

    close(ports)
    close(results)
    // Print open ports in order.
    sort.Ints(openports)
    for _, port := range openports {
        fmt.Printf("%d open\n", port)
    }

}
