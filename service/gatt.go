package service

import (
	"log"
	"github.com/bettercap/gatt"
)

//PM5 GATT Server UUIDs
var (
	attrGATTUUID,_=gatt.ParseUUID(getFullUUID("1801"))
	attrServiceChangedUUID,_=gatt.ParseUUID(getFullUUID("2A05"))
	attrGATTClientConfigChar,_=gatt.ParseUUID(getFullUUID("2902"))
)

// NewGattService registers a new GATT service as per PM5 specs
func NewGattService() *gatt.Service {
	s := gatt.NewService(attrGATTUUID)
	s.AddCharacteristic(attrServiceChangedUUID).HandleNotifyFunc(
		func(r gatt.Request, n gatt.Notifier) {
			go func() {
				log.Printf("TODO: indicate client when the services are changed")
			}()
		})
	s.AddCharacteristic(attrGATTClientConfigChar).SetValue([]byte{0x00,0x00})
	return s
}
