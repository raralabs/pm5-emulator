package service

import (
	"fmt"
	"log"

	"github.com/bettercap/gatt"
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
		fmt.Print("[[Control]:0021] received char : [ ")
		for _, d := range data {
			fmt.Printf("%x ", d)
		}
		fmt.Println("]")
		return gatt.StatusSuccess
	})

	/*
		C2 PM transmit characteristic
	*/
	transmitChar := s.AddCharacteristic(attrTransmitCharacteristicsUUID)
	transmitChar.HandleReadFunc(func(resp gatt.ResponseWriter, req *gatt.ReadRequest) {
		data := make([]byte, 20)
		resp.Write(data)
		log.Println("[[Control]:0022] Transmitting Data")
	})

	return s
}
