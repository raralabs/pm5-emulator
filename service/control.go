package service

import (
	"fmt"
	"log"
	"pm5-emulator/protocol/csafe"

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
	csafeDec := csafe.Decoder{}

	receiveChar := s.AddCharacteristic(attrReceiveCharacteristicsUUID)
	receiveChar.HandleWriteFunc(func(r gatt.Request, data []byte) (status byte) {
		pck, err := csafeDec.Decode(data)

		str := fmt.Sprintf("[[Control]] Decoded Command: 0x%x Data: [ ", pck.Cmds[0])
		for i := 0; i < len(pck.Data); i++ {
			str = fmt.Sprintf("%s0x%x ", str, pck.Data[i])
		}
		str = fmt.Sprintf("%s] Error: %v", str, err)

		log.Println(str)

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
