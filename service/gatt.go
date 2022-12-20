package service

import (
	"github.com/bettercap/gatt"
	"github.com/sirupsen/logrus"
)

//PM5 GATT Server UUIDs
var (
	attrGATTUUID, _             = gatt.ParseUUID(getFullUUID("1801"))
	attrServiceChangedUUID, _   = gatt.ParseUUID(getFullUUID("2A05"))
)

// NewGattService registers a new GATT service as per PM5 specs
func NewGattService() *gatt.Service {
	s := gatt.NewService(attrGATTUUID)
	c := s.AddCharacteristic(attrServiceChangedUUID)
	c.HandleNotifyFunc(
		func(r gatt.Request, n gatt.Notifier) {
			go func() {
				logrus.Info("TODO: indicate client when the services are changed")
			}()
		})

	c.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Handle Read")
		rsp.Write([]byte{0x0,0x0}) //2 bytes
	})

	return s
}
