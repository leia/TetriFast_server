package network

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/leia/TetriFast_server/commands"
)

func StartServer(address string) {
	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server started on port " + address)
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
	rw := bufio.NewReadWriter(bufio.NewReader(c), bufio.NewWriter(c))
	//command := ""
	//for {
	command, err := rw.ReadString('\n')
	if err != nil {
		if err != io.EOF {
			fmt.Println("read error:", err, command)
		}
		//break
	}

	//}

	fmt.Println("Getting command " + command)
	s := commands.ProcessCommand(command)

	if s != "" {
		SendResponse(s, rw)
	}
}

func SendResponse(response string, rw *bufio.ReadWriter) {
	fmt.Println("Sending response " + response)
	n, err := rw.WriteString(response)
	rw.Flush()

	if err != nil {
		if err != io.EOF {
			fmt.Println("read error:", err, n)
		}
	}

}
