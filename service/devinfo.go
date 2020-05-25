package service

import (
	"fmt"

	"github.com/bettercap/gatt"
)

func getFullUUID(uuid string) string {
	return fmt.Sprintf("%s%s%s", "CE06", uuid, "-43E5-11E4-916C-0800200C9A66")
}

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
	serialNumberId     = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
	firmwareRevisionId = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	manufacturerNameId = []byte{0x00, 0x18, 0x00, 0x18, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x18,
		0x00, 0x18, 0x00, 0x00, 0x03, 0xE8}
)

// NewDevInfoService registers a new Device Information service as per PM5 specs
func NewDevInfoService() *gatt.Service {
	s := gatt.NewService(attrDeviceInfoUUID)

	s.AddCharacteristic(attrModelNumberUUID).SetValue([]byte{0x00, 0x00})            //x0011
	s.AddCharacteristic(attrSerialNumberUUID).SetValue(serialNumberId)               //x0012
	s.AddCharacteristic(attrHardwareRevisionUUID).SetValue([]byte{0x00, 0x00, 0x00}) //x0013
	s.AddCharacteristic(attrFirmwareRevisionUUID).SetValue(firmwareRevisionId)       //x0014
	s.AddCharacteristic(attrManufacturerNameUUID).SetValue(manufacturerNameId)       //x0015
	s.AddCharacteristic(attrErgMachineTypeUUID).SetValue([]byte{0x00})               //x0016
	return s
}
