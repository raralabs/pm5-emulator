package main

import (
	"pm5-emulator/emulator"
	_ "pm5-emulator/log"
)

func main() {
	em := emulator.NewEmulator() //factory method
	em.RunEmulator()
	select {}
}
