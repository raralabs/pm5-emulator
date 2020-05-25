package service

import (
	"log"
	"github.com/bettercap/gatt"
)

//PM5 GATT Server UUIDs
var (
	attrGATTUUID,_=gatt.ParseUUID("CE061801-43E5-11E4-916C-0800200C9A66")
	attrServiceChangedUUID,_=gatt.ParseUUID("CE062A05-43E5-11E4-916C-0800200C9A66")
	attrGATTClientConfigChar,_=gatt.ParseUUID("CE062902-43E5-11E4-916C-0800200C9A66")
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
