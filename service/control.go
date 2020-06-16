package service

import (
	"fmt"
	"pm5-emulator/protocol/csafe"
	"pm5-emulator/service/decorator"
	"github.com/sirupsen/logrus"
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

//NewControlService advertises Control service offered by PM5
func NewControlService() *gatt.Service {
	controlService := gatt.NewService(attrControlServiceUUID)
	s := decorator.NewServiceSubscriber(controlService)

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

		logrus.Info(str)
		return 1
		//return gatt.StatusSuccess
	})

	receiveChar.HandleNotifyFunc(func(r gatt.Request, n gatt.Notifier) {
		logrus.Info(fmt.Sprintf("[[Control]] notify called by device ID: %v", r.Central.ID()))
		l := n.Cap()
		data := make([]byte, l)
		n.Write(data)
	})

	/*
		C2 PM transmit characteristic
	*/
	transmitChar := s.AddCharacteristic(attrTransmitCharacteristicsUUID)
	transmitChar.HandleNotifyFunc(func(r gatt.Request, n gatt.Notifier) {
		logrus.Info("[[Transmit]] Notify Signal")
		n.Write([]byte{0x76, 0x77, 0x7E, 0x7F})
	})

	transmitChar.HandleReadFunc(func(resp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("[[Transmit]] Transmitting Data")
		data := make([]byte, 20)
		resp.Write(data)
	})

	transmitChar.AddDescriptor(attrTransmitDescriptorUUID).SetValue([]byte{})

	return controlService
}
