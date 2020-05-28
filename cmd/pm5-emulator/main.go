package main

import "pm5-emulator/emulator"

func main() {
	em:=emulator.NewEmulator()  //factory method
	em.RunEmulator()
	select {}
}

