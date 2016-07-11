export GOOS=linux
export GOARCH=arm
go build
scp rovcontrol pi@192.168.1.38:/home/pi/rovcontrol/
scp -r html pi@192.168.1.38:/home/pi/rovcontrol/
scp -r dist pi@192.168.1.38:/home/pi/rovcontrol/
