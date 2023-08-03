package emulator

import (
	"log"
	"pm5-emulator/sm"
	"pm5-emulator/config/option"


	"github.com/bettercap/gatt"
)

//NewEmulator factory methods initializes emulator
func NewEmulator() *Emulator {
	d, err := gatt.NewDevice(option.DefaultServerOptions...)
	if err != nil {
		log.Fatalf("Failed to open config, err: %s", err)
	}
	return &Emulator{
		device:       d,
		stateMachine: sm.NewStateMachine(),
	}
}