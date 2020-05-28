package service

import (
	"log"

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
	*/
	modelNumChar := s.AddCharacteristic(attrModelNumberUUID)
	modelNumChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		log.Println("[[DevInfo]] Module Number String Read")
		data := make([]byte, 16) //16 bytes
		rsp.Write(data)
	})

	/*
		C2 PM Device Info: Serial Number characteristic
	*/
	serialNumberChar := s.AddCharacteristic(attrSerialNumberUUID)
	serialNumberChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		log.Println("[[DevInfo]] Serial Number String Read")
		data := make([]byte, 9) //9 bytes
		rsp.Write(data)
	})

	/*
		C2 PM Device Info: Hardware Revision characteristic
	*/
	hwRevChar := s.AddCharacteristic(attrHardwareRevisionUUID)
	hwRevChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		log.Println("[[DevInfo]] Hardware Revision String Read")
		data := make([]byte, 3) //3 bytes
		rsp.Write(data)
	})

	/*
		C2 PM Device Info: Firmware Revision characteristic
	*/
	fwRevChar := s.AddCharacteristic(attrFirmwareRevisionUUID)
	fwRevChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		log.Println("[[DevInfo]] Firmware Revision String Read")
		data := make([]byte, 20) //20 bytes
		rsp.Write(data)
	})

	/*
		C2 PM Device Info: Manufacturer Name characteristic
	*/
	manuNameChar := s.AddCharacteristic(attrManufacturerNameUUID)
	manuNameChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		log.Println("[[DevInfo]] Manufacturer Name String Read")
		data := make([]byte, 16) //16 bytes
		rsp.Write(data)
	})

	/*
		C2 PM Device Info: Erg Machine Type characteristic
	*/
	ergMachineTypeChar := s.AddCharacteristic(attrErgMachineTypeUUID)
	ergMachineTypeChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		log.Println("[[DevInfo]] Erg Machine Type Read")
		data := make([]byte, 1) //1 byte
		rsp.Write(data)
	})

	return s
}
