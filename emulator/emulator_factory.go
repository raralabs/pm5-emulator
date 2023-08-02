package emulator

import (
	"log"
	"pm5-emulator/sm"

	"github.com/bettercap/gatt"
)

//NewEmulator factory methods initializes emulator
func NewEmulator() *Emulator {
	d, err := gatt.NewDevice(defaultServerOptions...)
	if err != nil {
		log.Fatalf("Failed to open config, err: %s", err)
	}
	return &Emulator{
		device:       d,
		stateMachine: sm.NewStateMachine(),
	}
}

//BLE Server Options
var defaultServerOptions = []gatt.Option{
	gatt.MacDeviceRole(gatt.PeripheralManager),
}
