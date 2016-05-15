export GOOS=linux
export GOARCH=arm
go build
scp rovcontrol pi@172.17.9.98:/home/pi/rovcontrol/
scp -r html pi@172.17.9.98:/home/pi/rovcontrol/
scp -r dist pi@172.17.9.98:/home/pi/rovcontrol/
