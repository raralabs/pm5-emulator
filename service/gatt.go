package service

import (
	"log"

	"github.com/bettercap/gatt"
)

//PM5 GATT Server UUIDs
var (
	attrGATTUUID, _             = gatt.ParseUUID(getFullUUID("1801"))
	attrServiceChangedUUID, _   = gatt.ParseUUID(getFullUUID("2A05"))
	attrGATTClientConfigChar, _ = gatt.ParseUUID(getFullUUID("2902"))
)

// NewGattService registers a new GATT service as per PM5 specs
func NewGattService() *gatt.Service {
	s := gatt.NewService(attrGATTUUID)
	c:=s.AddCharacteristic(attrServiceChangedUUID)
	c.HandleNotifyFunc(
		func(r gatt.Request, n gatt.Notifier) {
			go func() {
				log.Printf("TODO: indicate client when the services are changed")
			}()
		})

	c.AddDescriptor(attrGATTClientConfigChar).HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest){
		log.Println("[[GATT]] Handle Read")
		d:=make([]byte,2)
		rsp.Write(d)
	})


	//c := s.AddCharacteristic(attrGATTClientConfigChar)
	//lv := byte(100)
	//
	//c.HandleReadFunc(
	//	func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
	//		log.Println("[[GATT]] Client Configuration Read")
	//		rsp.Write([]byte{lv})
	//		lv--
	//	})
	//c.HandleWriteFunc(func(r gatt.Request, data []byte) (status byte) {
	//	log.Println("[[GATT]] Client Configuration Write")
	//	fmt.Println(data)
	//	return 0x00
	//})
	return s
}
