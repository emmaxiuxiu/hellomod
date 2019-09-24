package main

import (
	"fmt"
	"github.com/donvito/hellomod"
	"github.com/emmaxiuxiu/hellomod/src/utils"
)

//SayHello function
func main() {
	fmt.Println(utils.AddHola("Hello World"))
	hellomod.SayHello()
}
