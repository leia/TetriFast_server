package network

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/leia/TetriFast_server/commands"
)

func StartServer(port string) {
	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server started on port " + port)
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Echo all incoming data.
			fmt.Println("Connection accepted")
			//io.Copy(c, c)
			// Shut down the connection.
			ProcessRequest(c)

			c.Close()
		}(conn)
	}
}

func ProcessRequest(c net.Conn) {
	rBytes := make([]byte, 256) // using small tmo buffer for demonstrating
	for {
		n, err := c.Read(rBytes)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err, n)
			}
			break
		}

	}
	fmt.Println(string(rBytes[:]))
	commands.ProcessCommand(string(rBytes[:]))
}
