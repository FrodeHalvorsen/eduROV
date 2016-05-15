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

	motor1Rear, err = hwio.GetPin("gpio17") // pin  Relay 1 OK GPIO
	err = hwio.PinMode(motor1Rear, hwio.OUTPUT)
	if err != nil {
		log.Println("Could not initalize motor1Rear")
	}
	motor2Rear, err = hwio.GetPin("gpio4") // pin  Relay 2
	err = hwio.PinMode(motor2Rear, hwio.OUTPUT)
	if err != nil {
		log.Println("Could not initalize motor2Rear")
	}
	motor3Rear, err = hwio.GetPin("gpio22") // pin  Relay 3 OK GPIO
	err = hwio.PinMode(motor3Rear, hwio.OUTPUT)
	if err != nil {
		log.Println("Could not initalize motor3Rear")
	}
	motor1Front, err = hwio.GetPin("gpio18") // pin  Relay 39 OK GPIO
	err = hwio.PinMode(motor1Front, hwio.OUTPUT)
	if err != nil {
		log.Println("Could not initalize motor1Front")
	}

	motor2Front, err = hwio.GetPin("gpio23") // pin  Relay 5 OK GPIO
	err = hwio.PinMode(motor2Front, hwio.OUTPUT)
	if err != nil {
		log.Println("Could not initalize motor2Front")
	}
	motor3Front, err = hwio.GetPin("gpio24") // pin  Relay 6 OK GPIO
	err = hwio.PinMode(motor3Front, hwio.OUTPUT)
	if err != nil {
		log.Println("Could not initalize motor3Front")
	}
	ledLight, err = hwio.GetPin("gpio10")
	err = hwio.PinMode(ledLight, hwio.OUTPUT)
	if err != nil {
		log.Println("Could not initalize motor1Front")
	}
}

func main() {

	hwio.DigitalWrite(ledLight, hwio.HIGH)

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
		case "up":
			up()
			break
		case "down":
			down()
			break
		case "stop":
			stop()
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
	log.Println("ROV going up")
	hwio.DigitalWrite(motor1Rear, hwio.LOW)
	hwio.DigitalWrite(motor2Rear, hwio.HIGH)
	hwio.DigitalWrite(motor3Rear, hwio.HIGH)
	hwio.DigitalWrite(motor1Front, hwio.LOW)
	hwio.DigitalWrite(motor2Front, hwio.HIGH)
	hwio.DigitalWrite(motor3Front, hwio.HIGH)
}

func down() {
	log.Println("ROV going down")
	hwio.DigitalWrite(motor1Rear, hwio.HIGH)
	hwio.DigitalWrite(motor2Rear, hwio.LOW)
	hwio.DigitalWrite(motor3Rear, hwio.LOW)
	hwio.DigitalWrite(motor1Front, hwio.HIGH)
	hwio.DigitalWrite(motor2Front, hwio.LOW)
	hwio.DigitalWrite(motor3Front, hwio.LOW)
}
func stop() {
	hwio.DigitalWrite(motor1Rear, hwio.LOW)
	hwio.DigitalWrite(motor2Rear, hwio.LOW)
	hwio.DigitalWrite(motor3Rear, hwio.LOW)
	hwio.DigitalWrite(motor1Front, hwio.LOW)
	hwio.DigitalWrite(motor2Front, hwio.LOW)
	hwio.DigitalWrite(motor3Front, hwio.LOW)
}
