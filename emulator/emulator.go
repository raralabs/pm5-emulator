package emulator

import (
	"fmt"
	"github.com/bettercap/gatt"
	"github.com/bettercap/gatt/linux/cmd"
	"log"
	"pm5-emulator/config"
	"pm5-emulator/service"
)



type Emulator struct{
	
}


func (em *Emulator) RunEmulator(){
	d, err := gatt.NewDevice(defaultServerOptions...)
	if err != nil {
		log.Fatalf("Failed to open config, err: %s", err)
	}

	// Register optional handlers.
	d.Handle(
		gatt.PeripheralConnected(func(p gatt.Peripheral, err error){log.Println("|Connect|: ") }),
		gatt.CentralConnected(func(c gatt.Central) { log.Println("|Connect|: ", c.ID()) }),
		gatt.CentralDisconnected(func(c gatt.Central) { log.Println("|Disconnect|: ", c.ID()) }),
	)

	// handler for monitoring config state.
	onStateChanged := func(d gatt.Device, s gatt.State) {
		fmt.Printf("State: %s\n", s)
		switch s {
		case gatt.StatePoweredOn:
			// Setup GAP and GATT services for PM5
			_ = d.AddService(service.NewGapService(config.NAME))
			_ = d.AddService(service.NewGattService())

			// Setup Device info service for PM5
			s1:=service.NewDevInfoService()
			d.AddService(s1)

			s2:=service.NewControlService()
			d.AddService(s2)

			s3:=service.NewRowingService()
			d.AddService(s3)

			// A fake battery service for demo.
			s4 := service.NewBatteryService()
			d.AddService(s4)

			// Advertise config name and service's UUIDs.
			d.AdvertiseNameAndServices(config.NAME, []gatt.UUID{s1.UUID(), s2.UUID(),s3.UUID(),s4.UUID()})

			// Advertise as an OpenBeacon iBeacon
			d.AdvertiseIBeacon(gatt.MustParseUUID("CE061800-43E5-11E4-916C-0800200C9A66"), 1, 2, -59)

		default:
		}
	}

	d.Init(onStateChanged)
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

