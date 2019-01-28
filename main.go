package main

import (
	"fly/routers"
	"fmt"
)

func main() {

	fmt.Println("Iniciando servidor")
	routers.InitServer()
}
