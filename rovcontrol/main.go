package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/desertbit/glue"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/jacobsa/go-serial/serial"
)

//Cmd ...
type Cmd struct {
	Action string `json:"action"`
}

var port io.ReadWriteCloser
var readChan chan string

func init() {
	// Set up options.
	options := serial.OpenOptions{
		PortName:        "/dev/cu.usbmodem1421",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 10,
	}

	var err error
	// Open the port.
	port, err = serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}
}

func main() {

	// Make sure to close it later.
	defer port.Close()

	readChan = make(chan string)

	go serialReaderHandler(readChan)

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./html", true)))
	r.Use(static.Serve("/dist/", static.LocalFile("./dist", true)))

	r.GET("/api", testHandler)

	// Create a new glue server.
	server := glue.NewServer(glue.Options{
		HTTPSocketType: glue.HTTPSocketTypeNone,
	})

	// Release the glue server on defer.
	// This will block new incoming connections
	// and close all current active sockets.
	defer server.Release()

	r.Use(gin.WrapH(server))

	// Set the glue event function to handle new incoming socket connections.
	server.OnNewSocket(onNewSocket)

	log.Println("Listening on :8080")

	r.Run(":8080")
}

func testHandler(c *gin.Context) {

	c.String(300, "Works")

}

func onNewSocket(s *glue.Socket) {

	// We won't read any data from the socket itself.
	// Discard received data!
	s.DiscardRead()

	// Set a function which is triggered as soon as the socket is closed.
	s.OnClose(func() {
		log.Printf("socket closed with remote address: %s", s.RemoteAddr())
	})

	// Set a function which is triggered during each received message.
	s.OnRead(func(data string) {
		// Echo the received data back to the client.
		s.Write(data)
	})

	// Create a channel.
	c := s.Channel("control")

	// Set the channel on read event function.
	c.OnRead(func(data string) {
		var cmd Cmd
		err := json.Unmarshal([]byte(data), &cmd)
		if err != nil {
			log.Println("Error decoding ctrl msg", err)
			return
		}
		fmt.Println("Got action", cmd.Action)
		port.Write([]byte(cmd.Action + "\r"))

		// Echo the received data back to the client.
		c.Write("ack: " + cmd.Action)
	})

	// Send a welcome string to the client.
	s.Write("Hello Client")
}

func serialReaderHandler(r chan string) {
	reader := bufio.NewReader(port)
	for {
		reply, err := reader.ReadBytes('\n')

		if err != nil {
			log.Printf("Could not read from serial: %s\n", err)
		}

		s := string(reply)
		fmt.Printf("Got from serial: %s", s)
		//r <- s

	}

}
