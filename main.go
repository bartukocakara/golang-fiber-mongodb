package main

import (
	"go-delivery-food/exception"
	"go-delivery-food/provider"
)

func main() {
	
	app := provider.ProvideService() 

	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)


}