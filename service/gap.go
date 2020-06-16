package service

import (
	"github.com/sirupsen/logrus"
	"github.com/bettercap/gatt"
)

//PM5 GAP Server UUIDs
var (
	attrGAPUUID, _               = gatt.ParseUUID(getFullUUID("1800"))
	attrDeviceNameUUID, _        = gatt.ParseUUID(getFullUUID("2A00"))
	attrAppearanceUUID, _        = gatt.ParseUUID(getFullUUID("2A01"))
	attrPeripheralPrivacyUUID, _ = gatt.ParseUUID(getFullUUID("2A02"))
	attrReconnectionAddrUUID, _  = gatt.ParseUUID(getFullUUID("2A03"))
	attrPeferredParamsUUID, _    = gatt.ParseUUID(getFullUUID("2A04"))
)

var gapCharAppearanceGenericComputer = []byte{0x00, 0x00}

//NewGapService registers a new GAP service as per PM5 specs
func NewGapService(name string) *gatt.Service {
	s := gatt.NewService(attrGAPUUID)

	/*
		GAP: Device Name characteristic
	*/
	devNameChar := s.AddCharacteristic(attrDeviceNameUUID)
	devNameChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Device Name Read")
		data := []byte(name)
		rsp.Write(data)
	})

	/*
		GAP: Device Name characteristic
	*/
	appearanceChar := s.AddCharacteristic(attrAppearanceUUID)
	appearanceChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Appearance Read")
		data := []byte{0x00, 0x00}
		rsp.Write(data)
	})

	/*
		GAP: Device Name characteristic
	*/
	ppChar := s.AddCharacteristic(attrPeripheralPrivacyUUID)
	ppChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Peripheral Privacy Read")
		data := []byte{0x00}
		rsp.Write(data)
	})

	/*
		GAP: Reconnect Address characteristic
	*/
	reconAddrChar := s.AddCharacteristic(attrReconnectionAddrUUID)
	reconAddrChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Reconnect Address Read")
		data := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
		rsp.Write(data)
	})

	/*
		GAP: Peripheral Preferred Connection Parameters characteristic
	*/
	prefParamChar := s.AddCharacteristic(attrPeferredParamsUUID)
	prefParamChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Preferred Connection Read")
		data := []byte{0x00, 0x18, 0x00, 0x18, 0x00, 0x00, 0x03, 0xE8}
		rsp.Write(data)
	})

	return s
}
