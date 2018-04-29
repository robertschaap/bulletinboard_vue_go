package main

import (
	"./controllers"
)

func main() {
	port := ":3000"
	controllers.APIController(port)
}
