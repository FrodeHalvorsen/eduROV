package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/desertbit/glue"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/mrmorphic/hwio"
)

//Cmd ...
type Cmd struct {
	Action string `json:"action"`
}

var readChan chan string

var motor1Rear hwio.Pin
var motor2Rear hwio.Pin
var motor3Rear hwio.Pin
var motor1Front hwio.Pin
var motor2Front hwio.Pin
var motor3Front hwio.Pin
var ledLight hwio.Pin

func init() {

	var err error

	// GPIO numbering (Fysical in comment)

	motor1Rear, err = hwio.GetPinWithMode("gpio2", hwio.OUTPUT) //03 12 oclock
	if err != nil {
		//log.Fatal("Could not initalize motor1Rear")
	}
	motor2Rear, err = hwio.GetPinWithMode("gpio3", hwio.OUTPUT)   //05 4 oclock
	motor3Rear, err = hwio.GetPinWithMode("gpio4", hwio.OUTPUT)   //07 8 oclock
	motor1Front, err = hwio.GetPinWithMode("gpio17", hwio.OUTPUT) //11 12 oclock
	motor2Front, err = hwio.GetPinWithMode("gpio27", hwio.OUTPUT) //13 4 oclock
	motor3Front, err = hwio.GetPinWithMode("gpio22", hwio.OUTPUT) //15 8 oclock
	ledLight, err = hwio.GetPinWithMode("gpio10", hwio.OUTPUT)    //19
	if err != nil {
		log.Println("Could not initalize ledLight")
	}
}

func main() {

	readChan = make(chan string)

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

		switch cmd.Action {
		case "fwd":
			forward()
			break
		case "bck":
			backward()
			break
		case "left":
			left()
			break
		case "right":
			right()
			break
		default:
			log.Println("Got unknown action!")
		}

		// Echo the received data back to the client.
		c.Write("ack: " + cmd.Action)
	})

	// Send a welcome string to the client.
	s.Write("Hello Client")

}

func forward() {
	log.Println("running motor fwd")
	hwio.DigitalWrite(motor1Rear, hwio.HIGH)
	hwio.DigitalWrite(motor2Rear, hwio.HIGH)
	hwio.DigitalWrite(motor3Rear, hwio.HIGH)
	hwio.DigitalWrite(motor1Front, hwio.LOW)
	hwio.DigitalWrite(motor2Front, hwio.LOW)
	hwio.DigitalWrite(motor3Front, hwio.LOW)
}

func backward() {
	log.Println("running motor bck")
	hwio.DigitalWrite(motor1Rear, hwio.LOW)
	hwio.DigitalWrite(motor2Rear, hwio.LOW)
	hwio.DigitalWrite(motor3Rear, hwio.LOW)
	hwio.DigitalWrite(motor1Front, hwio.HIGH)
	hwio.DigitalWrite(motor2Front, hwio.HIGH)
	hwio.DigitalWrite(motor3Front, hwio.HIGH)
}

func left() {
	log.Println("running motor left")
	hwio.DigitalWrite(motor1Rear, hwio.LOW)
	hwio.DigitalWrite(motor2Rear, hwio.HIGH)
	hwio.DigitalWrite(motor3Rear, hwio.LOW)
	hwio.DigitalWrite(motor1Front, hwio.LOW)
	hwio.DigitalWrite(motor2Front, hwio.LOW)
	hwio.DigitalWrite(motor3Front, hwio.HIGH)
}

func right() {
	log.Println("running motor right")
	hwio.DigitalWrite(motor1Rear, hwio.LOW)
	hwio.DigitalWrite(motor2Rear, hwio.LOW)
	hwio.DigitalWrite(motor3Rear, hwio.HIGH)
	hwio.DigitalWrite(motor1Front, hwio.LOW)
	hwio.DigitalWrite(motor2Front, hwio.HIGH)
	hwio.DigitalWrite(motor3Front, hwio.LOW)
}

func up() {
	hwio.DigitalWrite(motor1Rear, hwio.LOW)
	hwio.DigitalWrite(motor2Rear, hwio.HIGH)
	hwio.DigitalWrite(motor3Rear, hwio.HIGH)
	hwio.DigitalWrite(motor1Front, hwio.LOW)
	hwio.DigitalWrite(motor2Front, hwio.HIGH)
	hwio.DigitalWrite(motor3Front, hwio.HIGH)
}

func down() {
	hwio.DigitalWrite(motor1Rear, hwio.HIGH)
	hwio.DigitalWrite(motor2Rear, hwio.LOW)
	hwio.DigitalWrite(motor3Rear, hwio.LOW)
	hwio.DigitalWrite(motor1Front, hwio.HIGH)
	hwio.DigitalWrite(motor2Front, hwio.LOW)
	hwio.DigitalWrite(motor3Front, hwio.LOW)
}
