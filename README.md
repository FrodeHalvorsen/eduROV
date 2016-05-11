# eduROV
Open Source ROV for educational use.

This is the first iteration of this ROV.

It`s currently in developement, and is still not testet in the water.

To compile and run go program:

- open terminal
$ cd ~/go/src/github.com/FrodeHalvorsen/eduROV/rovcontrol
$ go run main.go (to run locally)

To compile for raspberry pi:

$ export GOOS=linux
$ export GOARCH=arm
$ go build main.go

copy html, dist and rovcontrol binary file to raspberry pi and run.

$ scp rovcontrol pi@IP-TO-PI:~/rovcontrol/
$ scp -r html pi@IP-TO-PI:~/rovcontrol/
$ scp -r dist pi@IP-TO-PI:~/rovcontrol/

Log into pi:

$ cd rovcontrol
$ ./rovcontrol
