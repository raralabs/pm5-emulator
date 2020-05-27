package service

import (
	"github.com/bettercap/gatt"
	"log"
)

/*
C2 PM control
primary service
*/

var (
	attrControlServiceUUID, _          = gatt.ParseUUID(getFullUUID("0020"))
	attrReceiveCharacteristicsUUID, _  = gatt.ParseUUID(getFullUUID("0021"))
	attrTransmitCharacteristicsUUID, _ = gatt.ParseUUID(getFullUUID("0022"))
)

func NewControlService() *gatt.Service {
	s := gatt.NewService(attrControlServiceUUID)

	/*
		C2 PM receive characteristic
	*/
	receiveChar := s.AddCharacteristic(attrReceiveCharacteristicsUUID)
	receiveChar.HandleWriteFunc(func(r gatt.Request, data []byte) (status byte) {
		log.Println("[[Control]] received char : ", string(data))
		return gatt.StatusSuccess
	})

	/*
		C2 PM transmit characteristic
	*/
	transmitChar := s.AddCharacteristic(attrTransmitCharacteristicsUUID)
	transmitChar.HandleReadFunc(func(resp gatt.ResponseWriter, req *gatt.ReadRequest) {
		data := make([]byte, 20)
		resp.Write(data)
		log.Println("[[Control]] Transmitting Data")
	})

	return s
}
