## simple makefile
.PHONY: all clean dockerfile local build install
RPI_IP = 192.168.1.2
SSH_USER = pi
PROJECT_NAME = edurov
FILENAME_LINUX = $(PROJECT_NAME)_linux_arm
FILENAME_OSX = $(PROJECT_NAME)_darwin_amd64
pi:
	GOOS=linux GOARCH=arm go build -o $(FILENAME_LINUX)
	ssh $(SSH_USER)@$(RPI_IP) "sudo service $(PROJECT_NAME) stop || true"
	scp $(FILENAME_LINUX) $(SSH_USER)@$(RPI_IP):/home/$(SSH_USER)/rovcontrol/rovcontrol
	scp -r html $(SSH_USER)@$(RPI_IP):/home/$(SSH_USER)/rovcontrol/
	scp -r dist $(SSH_USER)@$(RPI_IP):/home/$(SSH_USER)/rovcontrol/
	ssh $(SSH_USER)@$(RPI_IP) "sudo service $(PROJECT_NAME) start || true"
pi-go:
	GOOS=linux GOARCH=arm go build -o $(FILENAME_LINUX)
	ssh $(SSH_USER)@$(RPI_IP) "sudo service $(PROJECT_NAME) stop || true"
	scp $(FILENAME_LINUX) $(SSH_USER)@$(RPI_IP):/home/$(SSH_USER)/rovcontrol/rovcontrol
	ssh $(SSH_USER)@$(RPI_IP) "sudo service $(PROJECT_NAME) start || true"
local:
	go build -o $(FILENAME_OSX)
	./$(FILENAME_OSX)
clean:
	@go clean $(GOFLAGS) -i ./...
## EOF
