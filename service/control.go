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
	attrTransmitDescriptorUUID, _      = gatt.ParseUUID(getFullUUID("2902"))
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
		return 1
		//return gatt.StatusSuccess
	})

	receiveChar.HandleNotifyFunc(func(r gatt.Request, n gatt.Notifier) {
		str := fmt.Sprintf("[[control]] notify called by device ID: %v", r.Central.ID())
		log.Println(str)
		l := n.Cap()
		data := make([]byte, l)
		n.Write(data)
	})

	/*
		C2 PM transmit characteristic
	*/
	transmitChar := s.AddCharacteristic(attrTransmitCharacteristicsUUID)
	transmitChar.HandleNotifyFunc(func(r gatt.Request, n gatt.Notifier) {
		log.Println("[[Control Transmit]] Notify Signal")

		//enc:=csafe.Encoder{}
		//pkt:=csafe.Packet{
		//	Data: []byte{0x1}, //sending connection set = true
		//}
		//buf:=enc.Encode(pkt)
		n.Write([]byte{0x76,0x77,0x7E,0x7F})
		//n.Write([]byte{0x1})
	})

	transmitChar.HandleReadFunc(func(resp gatt.ResponseWriter, req *gatt.ReadRequest) {
		data := make([]byte, 20)
		resp.Write(data)
		log.Println("[[Control]:0022] Transmitting Data")
	})

	transmitChar.AddDescriptor(attrTransmitDescriptorUUID).SetValue([]byte{})

	return s
}
