## Network programming
Go comes with a formidable set of functionalities for network communication using the `net` package.  We will only scratch the surface with a couple of examples to give you an idea what is possible with a few lines of code.

### A simple socket server 
To get started, let us create a simple echo server and client.  The server is implemented with a few lines of code, but is capable of scaling to handle large number of connections using a goroutine as shown below:
```go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("echo server running on 4040...")

	// connection loop
	for {
		// wait and accept new client connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// handle connection
		// read text and echo it
		go func(c net.Conn) {
			defer conn.Close()
			reader := bufio.NewReader(c)
			for {
				line, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println(err)
					return
				}
				if _, err := fmt.Fprint(conn, line); err != nil {
					fmt.Println(err)
					return
				}
			}
		}(conn)
	}
}
```
When the echo server program is started, a telnet session can be used to test it as shown below:
```sh
> telnet localhost 4040
Trying ::1...
Connected to localhost.
Escape character is '^]'.
hello
hello
I like your hat
I like your hat
```
###A simple socket client
Writing a socket client in Go is just as easy.  The following is a simple client to the echo server above.  It captures the keyboard input using buffered IO reader `cmdReader`, then sends the data to the server, and lastly, prints the response from the server captured with `conReader`.
```go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var host, port = "127.0.0.1", "4040"
var addr = net.JoinHostPort(host, port)

func main() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connected to echo server at:", addr)

	for {
		cmdReader := bufio.NewReader(os.Stdin)
		conReader := bufio.NewReader(conn)

		// read command-line input
		data, err := cmdReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		// send data to server
		if _, err := fmt.Fprint(conn, data); err != nil {
			fmt.Println(err)
			continue
		}

		// read server response
		resp, err := conReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Print(resp)
	}
}
```
### A simple web server
Another well-appointed capability of Go is its support for creating HTTP-based servers and clients using the `http` package.  The following shows a program that implements a time API server which returns the current date/time of the server.  
```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprint(writer, "/now\n/date\n/time")
	})

	http.HandleFunc("/now", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "%s", time.Now())
	})

	http.HandleFunc("/date", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "%s", time.Now().Format("2006-01-02"))
	})

	http.HandleFunc("/time", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "%s", time.Now().Format("03:04:05 MST"))
	})

	fmt.Println("time api server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
```
When we run the program, it immediately blocks to listen and serve incoming HTTP requests on port `8080`.  Using `cURL` we can test the program as shown below:
```sh
$ curl localhost:8080
/now
/date
/time

$ curl localhost:8080/now
2017-08-02 08:32:14.854542532 -0400 EDT

$ curl localhost:8080/time
08:33:34 ED

$ curl localhost:8080/date
2017-08-02
```
The previous program uses a default package-provided server object with default configurations.  The URL routes are handled by a default packaged-provider handler using function `http.HandleFunc` which takes a path and the function used to handle HTTP request.  

However, the program can be re-written using an explicit declarations of an `http.Server` value along with an `http.Handler` object used for route multiplexing called `http.ServeMux`.  This explicit declaration of these objects provides deeper control and configuration of the server as shown below:
```go
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprint(writer, "/now\n/date\n/time")
	})

	mux.HandleFunc("/now", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "%s", time.Now())
	})

	mux.HandleFunc("/date", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "%s", time.Now().Format("2006-01-02"))
	})

	mux.HandleFunc("/time", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "%s", time.Now().Format("03:04:05 MST"))
	})

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 3,
	}

	fmt.Println("time api server running on :8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
```
### Writing HTTP clients
The Go `http` package comes with everything you need to write robust HTTP client programs.  The following shows a simple client, using default configurations, to the time api server above. 
```go
func main() {
	resp, err := http.Get("http://127.0.0.1:8080/now")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		fmt.Println(err)
	}
}
```
The previous code uses a default client object provided by the `http` package.  If you want want more control over client configuration, you can explicitly declare your own client object as is shown below:
```go
func main() {
	client := &http.Client{
		Timeout: 21 * time.Second,
	}
	resp, err := client.Get("http://127.0.0.1:8080/date")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

### Next
Continue to the next section to learn about [testing](./testing.md) in Go.