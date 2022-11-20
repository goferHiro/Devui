package main

import (
	"fmt"
	"github.com/goferHiro/DevEUI/factory"
)

func main() {
	fmt.Println("Hey")

	/*	factory := "78111FFFE452555B"
		lorowanServices := lorowan.NewServices()
		lorowanServices.RegisterDEVEUI(factory)*/

	factoryServices := factory.NewServices()
	devuis := factoryServices.BatchOf100()
	factoryServices.ProduceBatch100(devuis)
}
