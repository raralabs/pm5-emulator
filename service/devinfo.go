package service

import (
	"pm5-emulator/config"

	"github.com/bettercap/gatt"
)

//PM5 Device Info Server UUIDs
var (
	attrDeviceInfoUUID, _       = gatt.ParseUUID(getFullUUID("0010"))
	attrModelNumberUUID, _      = gatt.ParseUUID(getFullUUID("0011"))
	attrSerialNumberUUID, _     = gatt.ParseUUID(getFullUUID("0012"))
	attrHardwareRevisionUUID, _ = gatt.ParseUUID(getFullUUID("0013"))
	attrFirmwareRevisionUUID, _ = gatt.ParseUUID(getFullUUID("0014"))
	attrManufacturerNameUUID, _ = gatt.ParseUUID(getFullUUID("0015"))
	attrErgMachineTypeUUID, _   = gatt.ParseUUID(getFullUUID("0016"))
)

var (
	firmwareRevisionId = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
)

// NewDevInfoService registers a new Device Information service as per PM5 specs
func NewDevInfoService() *gatt.Service {
	s := gatt.NewService(attrDeviceInfoUUID)

	s.AddCharacteristic(attrModelNumberUUID).SetValue([]byte(config.MODEL_NO))               //x0011
	s.AddCharacteristic(attrSerialNumberUUID).SetValue([]byte(config.SERIAL_NO))             //x0012
	s.AddCharacteristic(attrHardwareRevisionUUID).SetValue([]byte{0x00, 0x00, 0x00})         //x0013
	s.AddCharacteristic(attrFirmwareRevisionUUID).SetValue(firmwareRevisionId)               //x0014
	s.AddCharacteristic(attrManufacturerNameUUID).SetValue([]byte(config.MANUFACTURER_NAME)) //x0015
	s.AddCharacteristic(attrErgMachineTypeUUID).SetValue([]byte{0x00})                       //x0016
	return s
}
