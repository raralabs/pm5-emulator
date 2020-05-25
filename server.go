package main

import (
	"fmt"
	"log"
	"pm5-emulator/service"

	"github.com/bettercap/gatt"
	"github.com/bettercap/gatt/linux/cmd"
)

func main() {
	d, err := gatt.NewDevice(DefaultServerOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s", err)
	}

	// Register optional handlers.
	d.Handle(
		gatt.CentralConnected(func(c gatt.Central) { fmt.Println("Connect: ", c.ID()) }),
		gatt.CentralDisconnected(func(c gatt.Central) { fmt.Println("Disconnect: ", c.ID()) }),
	)

	// handler for monitoring device state.
	onStateChanged := func(d gatt.Device, s gatt.State) {
		fmt.Printf("State: %s\n", s)
		switch s {
		case gatt.StatePoweredOn:
			// Setup GAP and GATT services for PM5
			_ = d.AddService(service.NewGapService("PM5 430000000"))
			_ = d.AddService(service.NewGattService())

			// Setup Device info service for PM5
			_ = d.AddService(service.NewDevInfoService())

			// A fake battery service for demo.
			s1 := service.NewBatteryService()
			d.AddService(s1)

			// Advertise device name and service's UUIDs.
			d.AdvertiseNameAndServices("PM5 430000000", []gatt.UUID{s1.UUID()})

			// Advertise as an OpenBeacon iBeacon
			d.AdvertiseIBeacon(gatt.MustParseUUID("CE061800-43E5-11E4-916C-0800200C9A66"), 1, 2, -59)

		default:
		}
	}

	d.Init(onStateChanged)
	select {}
}

//BLE Server Options
var DefaultServerOptions = []gatt.Option{
	gatt.LnxMaxConnections(1),
	gatt.LnxDeviceID(-1, true),
	gatt.LnxSetAdvertisingParameters(&cmd.LESetAdvertisingParameters{
		AdvertisingIntervalMin: 0x00f4,
		AdvertisingIntervalMax: 0x00f4,
		AdvertisingChannelMap:  0x7,
	}),
}
