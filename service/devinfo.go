package service

import (
	"pm5-emulator/config"
	"github.com/sirupsen/logrus"
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

// NewDevInfoService registers a new Device Information service as per PM5 specs
func NewDevInfoService() *gatt.Service {
	s := gatt.NewService(attrDeviceInfoUUID)

	/*
		C2 PM Device Info: Model Number characteristic
		0x0011
	*/
	modelNumChar := s.AddCharacteristic(attrModelNumberUUID)
	modelNumChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Module Number String Read")
		rsp.Write([]byte(config.MODEL_NO)) //upto 16 bytes
	})

	/*
		C2 PM Device Info: Serial Number characteristic
		0x0012
	*/
	serialNumberChar := s.AddCharacteristic(attrSerialNumberUUID)
	serialNumberChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Serial Number String Read")
		rsp.Write([]byte(config.SERIAL_NO)) //write serial number as response
	})

	/*
		C2 PM Device Info: Hardware Revision characteristic
		0x0013
	*/
	hwRevChar := s.AddCharacteristic(attrHardwareRevisionUUID)
	hwRevChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Hardware Revision String Read")
		rsp.Write([]byte(config.HARDWARE_VERSION)) //upto 3 bytes
	})

	/*
		C2 PM Device Info: Firmware Revision characteristic
		0x0014
	*/
	fwRevChar := s.AddCharacteristic(attrFirmwareRevisionUUID)
	fwRevChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Firmware Revision String Read")
		rsp.Write([]byte(config.FIRMWARE_VERSION)) //upto 20bytes
	})

	/*
		C2 PM Device Info: Manufacturer Name characteristic
		0x0015
	*/
	manuNameChar := s.AddCharacteristic(attrManufacturerNameUUID)
	manuNameChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Manufacturer Name String Read")
		rsp.Write([]byte(config.MANUFACTURER_NAME)) //upto 16 bytes
	})

	/*
		C2 PM Device Info: Erg Machine Type characteristic
		0x0016
	*/
	ergMachineTypeChar := s.AddCharacteristic(attrErgMachineTypeUUID)
	ergMachineTypeChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Erg Machine Type Read")
		rsp.Write([]byte(config.ERG_MACHINE_TYPE)) //upto 1 byte
	})

	return s
}
