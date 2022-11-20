package main

import (
	"fmt"
	"github.com/goferHiro/DevEUI/factory"
	"os"
	"os/signal"
	"syscall"
)

func getKillSignal() chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGKILL,
	)
	return c
}

func main() {

	exit := getKillSignal()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("exiting graacefully")
		} else {
			fmt.Println("exiting with exit code 0")
		}

	}()

	go func() {
		<-exit
		fmt.Println("got the kill signal")
	}()

	fmt.Println("starting the cli")
	factoryServices := factory.NewServices()
	devuis := factoryServices.BatchOf100()
	factoryServices.ProduceBatch100(devuis)

}
