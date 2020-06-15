package emulator

import (
	"fmt"
	"log"
	"pm5-emulator/config"
	"pm5-emulator/service"
	"pm5-emulator/sm"

	"github.com/bettercap/gatt"
)

//Emulator emulates PM5 indoor rower machine
type Emulator struct {
	device       gatt.Device
	stateMachine *sm.StateMachine
}

//RunEmulator registers handlers and starts advertising services
func (em *Emulator) RunEmulator() {

	//register optional handlers
	em.registerHandlers()

	// handler for monitoring config state.
	onStateChanged := func(d gatt.Device, s gatt.State) {
		fmt.Printf("State: %s\n", s)
		switch s {
		case gatt.StatePoweredOn:
			// Setup GAP and GATT services for PM5
			_ = d.AddService(service.NewGapService(config.NAME))
			_ = d.AddService(service.NewGattService())

			// Setup Device info service for PM5
			s1 := service.NewDevInfoService()
			d.AddService(s1)

			s2 := service.NewControlService()
			d.AddService(s2)

			s3 := service.NewRowingService()
			d.AddService(s3)

			// Advertise config name and service's UUIDs.
			d.AdvertiseNameAndServices(config.NAME, []gatt.UUID{s1.UUID(), s2.UUID(), s3.UUID()})

			// Advertise as an OpenBeacon iBeacon
			d.AdvertiseIBeacon(gatt.MustParseUUID("CE061800-43E5-11E4-916C-0800200C9A66"), 1, 2, -59)

		default:
		}
	}

	em.device.Init(onStateChanged)
}

//registerHandlers registers optional handlers for handling device connection and disconnection
func (em *Emulator) registerHandlers() {
	// Register optional handlers.
	em.device.Handle(
		gatt.PeripheralConnected(func(p gatt.Peripheral, err error) {
			log.Println("|Peripheral Connected|: ")
			fmt.Println("ID: ", p.ID())
			fmt.Println("Device: ", p.Device())
			fmt.Println("Name: ", p.Name())
		}),
		gatt.CentralConnected(func(c gatt.Central) {
			log.Println("|Device Connected| ID=> ", c.ID())
			fmt.Println("MTU: ", c.MTU())
		}),
		gatt.CentralDisconnected(func(c gatt.Central) {
			log.Println("|Device Disconnected| ID=> ", c.ID())
			fmt.Println("MTU: ", c.MTU())
		}),
	)
}
