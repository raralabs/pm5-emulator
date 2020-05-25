package service

import "github.com/bettercap/gatt"

//PM5 GAP Server UUIDs
var (
	attrGAPUUID,_ = gatt.ParseUUID("CE061800-43E5-11E4-916C-0800200C9A66")
	attrDeviceNameUUID,_ = gatt.ParseUUID("CE062A00-43E5-11E4-916C-0800200C9A66")
	attrAppearanceUUID,_= gatt.ParseUUID("CE062A01-43E5-11E4-916C-0800200C9A66")
	attrPeripheralPrivacyUUID,_=gatt.ParseUUID("CE062A02-43E5-11E4-916C-0800200C9A66")
	attrReconnectionAddrUUID,_=gatt.ParseUUID("CE062A03-43E5-11E4-916C-0800200C9A66")
	attrPeferredParamsUUID,_=gatt.ParseUUID("CE062A04-43E5-11E4-916C-0800200C9A66")
)

var gapCharAppearanceGenericComputer = []byte{0x00, 0x00}

//NewGapService registers a new GAP service as per PM5 specs
func NewGapService(name string) *gatt.Service {
	s := gatt.NewService(attrGAPUUID)
	s.AddCharacteristic(attrDeviceNameUUID).SetValue([]byte(name)) //x2A00
	s.AddCharacteristic(attrAppearanceUUID).SetValue(gapCharAppearanceGenericComputer)  //x2A01
	s.AddCharacteristic(attrPeripheralPrivacyUUID).SetValue([]byte{0x00}) //x2A02
	s.AddCharacteristic(attrReconnectionAddrUUID).SetValue([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00})  // x2A03
	s.AddCharacteristic(attrPeferredParamsUUID).SetValue([]byte{0x00, 0x18, 0x00, 0x18, 0x00, 0x00, 0x03, 0xE8})  //x2A04
	return s
}
