package emulator

import (
	"log"
	"pm5-emulator/sm"

	"github.com/bettercap/gatt"
	"github.com/bettercap/gatt/linux/cmd"
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
	gatt.LnxMaxConnections(1),
	gatt.LnxDeviceID(-1, true),
	gatt.LnxSetAdvertisingParameters(&cmd.LESetAdvertisingParameters{
		AdvertisingIntervalMin: 0x00f4,
		AdvertisingIntervalMax: 0x00f4,
		AdvertisingChannelMap:  0x7,
	}),
}
