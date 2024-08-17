package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/teclegacy/ms/oms/common"
)

func HelloAir() {
	log.Println("air HMR")
}

var value = common.EnvString("CONST_PORT", "1010")

func main() {

	HelloAir()

	log.Printf("value %v", value)
}
