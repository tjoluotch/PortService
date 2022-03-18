package serverutils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const (
	NETWORK  = "tcp"
	PORT     = 50051
	HOSTNAME = "localhost"
)

func GetAddress() string {
	return fmt.Sprintf("%s:%d", HOSTNAME, PORT)
}

func InitChannel() chan os.Signal {
	chanExit := make(chan os.Signal, 1)
	signal.Notify(chanExit, os.Interrupt, syscall.SIGTERM)
	return chanExit
}
