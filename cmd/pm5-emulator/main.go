package main

import "pm5-emulator/emulator"

func main() {
	em:=emulator.Emulator{}
	em.RunEmulator()
	select {}
}

