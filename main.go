package main

import (
	"fmt"
	"github.com/goferHiro/DevEUI/factory"
	"time"
)

func main() {
	fmt.Println("Hey")

	/*	factory := "78111FFFE452555B"
		lorowanServices := lorowan.NewServices()
		lorowanServices.RegisterDEVEUI(factory)*/

	factoryServices := factory.NewServices()

	devuis := factoryServices.BatchOf100()
	fmt.Println("devuis are ", devuis)
	factoryServices.BatchOf100()
	time.Sleep(10 * time.Second)
}
