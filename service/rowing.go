package service

import (
	"log"

	"github.com/bettercap/gatt"
)

/*
	C2 rowing primary service
*/

//C2 rowing primary service and characteristics UUIDs
var (
	attrRowingServiceUUID, _                                    = gatt.ParseUUID(getFullUUID("0030"))
	attrGeneralStatusCharacteristicsUUID, _                     = gatt.ParseUUID(getFullUUID("0031"))
	attrGeneralStatusDescriptorUUID, _                          = gatt.ParseUUID(getFullUUID("2902"))
	attrAdditionalStatus1CharacteristicsUUID, _                 = gatt.ParseUUID(getFullUUID("0032"))
	attrAdditionalStatus2CharacteristicsUUID, _                 = gatt.ParseUUID(getFullUUID("0033"))
	attrSampleRateCharacteristicsUUID, _                        = gatt.ParseUUID(getFullUUID("0034"))
	attrStrokeDataCharacteristicsUUID, _                        = gatt.ParseUUID(getFullUUID("0035"))
	attrAdditionalStrokeDataCharacteristicsUUID, _              = gatt.ParseUUID(getFullUUID("0036"))
	attrSplitIntervalDataCharacteristicsUUID, _                 = gatt.ParseUUID(getFullUUID("0037"))
	attrAdditionalSplitIntervalDataCharacteristicsUUID, _       = gatt.ParseUUID(getFullUUID("0038"))
	attrEndOfWorkoutSummaryDataCharacteristicsUUID, _           = gatt.ParseUUID(getFullUUID("0039"))
	attrAdditionalEndOfWorkoutSummaryDataCharacteristicsUUID, _ = gatt.ParseUUID(getFullUUID("003A"))
	attrHeartRateBeltInfoCharacteristicsUUID, _                 = gatt.ParseUUID(getFullUUID("003B"))
	attrForceCurveDataCharacteristicsUUID, _                    = gatt.ParseUUID(getFullUUID("003D"))
	attrMultiplexedInfoCharacteristicsUUID, _                   = gatt.ParseUUID(getFullUUID("0080"))
)

func NewRowingService() *gatt.Service {
	s := gatt.NewService(attrRowingServiceUUID)

	/*
		C2 rowing general status characteristic
	*/
	rowingGenStatusChar := s.AddCharacteristic(attrGeneralStatusCharacteristicsUUID)
	rowingGenStatusChar.HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			data := make([]byte, 19) //19 bytes
			rsp.Write(data)
			log.Println("[[Rowing]] General Status Char Read Request")
		})

	rowingGenStatusChar.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})

	/*
		C2 rowing additional status 1 characteristic
	*/
	additionalStatus1Char := s.AddCharacteristic(attrAdditionalStatus1CharacteristicsUUID)
	additionalStatus1Char.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		data := make([]byte, 16)
		rsp.Write(data)
		log.Println("[[Rowing]] Additional Status 1 Char Read Request")
	})

	additionalStatus1Char.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})

	/*
		C2 rowing additional status 2 characteristic
	*/
	additionalStatus2Char := s.AddCharacteristic(attrAdditionalStatus2CharacteristicsUUID)
	additionalStatus2Char.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		data := make([]byte, 16)
		rsp.Write(data)
		log.Println("[[Rowing]] Additional Status 2 Status Char Read Request")
	})

	additionalStatus2Char.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})
	/*
		C2 rowing general status and additional status sample rate characteristic
	*/
	sampleRateChar := s.AddCharacteristic(attrSampleRateCharacteristicsUUID)
	sampleRateChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		data := make([]byte, 1)
		rsp.Write(data)
		log.Println("[[Rowing]] Sample Rate Char Read Request")
	})

	/*
		C2 rowing stroke data  characteristic
	*/
	strokeDataChar := s.AddCharacteristic(attrStrokeDataCharacteristicsUUID)
	strokeDataChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		data := make([]byte, 20)
		rsp.Write(data)
		log.Println("[[Rowing]] Stroke Data char Read Request")
	})

	strokeDataChar.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})

	/*
		C2 rowing additional stroke data characteristic
	*/
	additionalStrokeDataChar := s.AddCharacteristic(attrAdditionalStrokeDataCharacteristicsUUID)
	additionalStrokeDataChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		data := make([]byte, 15)
		rsp.Write(data)
		log.Println("[[Rowing]] Additional Stroke Data char Read Request")
	})

	additionalStrokeDataChar.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})

	/*
		C2 rowing split/interval data characteristic
	*/
	splitIntervalDataChar := s.AddCharacteristic(attrSplitIntervalDataCharacteristicsUUID)
	splitIntervalDataChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		data := make([]byte, 18)
		rsp.Write(data)
		log.Println("[[Rowing]] Split/Interval Data char Read Request")
	})

	splitIntervalDataChar.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})

	/*
		C2 rowing additional split/interval data characteristic
	*/
	additionalSplitIntervalDataChar := s.AddCharacteristic(attrAdditionalSplitIntervalDataCharacteristicsUUID)
	additionalSplitIntervalDataChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		data := make([]byte, 18)
		rsp.Write(data)
		log.Println("[[Rowing]] Additional Split/Interval Data char Read Request")
	})

	additionalSplitIntervalDataChar.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})
	/*
		C2 rowing end of workout summary data characteristic
	*/
	endOfWorkoutSummaryDataChar := s.AddCharacteristic(attrEndOfWorkoutSummaryDataCharacteristicsUUID)
	endOfWorkoutSummaryDataChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		data := make([]byte, 20)
		rsp.Write(data)
		log.Println("[[Rowing]] End of workout summary Data char Read Request")
	})

	endOfWorkoutSummaryDataChar.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})
	/*
		C2 rowing end of workout additional summary data characteristic
	*/
	additionalEndOfWorkoutSummaryDataChar := s.AddCharacteristic(attrAdditionalEndOfWorkoutSummaryDataCharacteristicsUUID)
	additionalEndOfWorkoutSummaryDataChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		data := make([]byte, 19)
		rsp.Write(data)
		log.Println("[[Rowing]] Additional End of workout summary Data char Read Request")
	})

	additionalEndOfWorkoutSummaryDataChar.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})

	/*
		C2 rowing heart rate belt information characteristic
	*/
	heartRateBeltInfoChar := s.AddCharacteristic(attrHeartRateBeltInfoCharacteristicsUUID)
	//handle read
	heartRateBeltInfoChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		data := make([]byte, 6)
		rsp.Write(data)
		log.Println("[[Rowing]] Heart Rate Belt Info Read Request")
	})

	//handle write
	heartRateBeltInfoChar.HandleWriteFunc(func(req gatt.Request, data []byte) (status byte) {
		log.Println("[[Rowing]] received data at heart rate belt info: ", string(data))
		return gatt.StatusSuccess
	})

	heartRateBeltInfoChar.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})

	/*
		C2 force curve data characteristic
	*/
	forceCurveDataChar := s.AddCharacteristic(attrForceCurveDataCharacteristicsUUID)
	forceCurveDataChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		data := make([]byte, 288)
		rsp.Write(data)
		log.Println("[[Rowing]] Force Curve Data Read Request")
	})

	forceCurveDataChar.HandleWriteFunc(func(req gatt.Request, data []byte) (status byte) {
		log.Println("[[Rowing]] received data at force curve data: ", string(data))
		return gatt.StatusSuccess
	})

	/*
		C2 multiplexed information 	characteristic

		0x0080 | Up to 20 bytes | READ Permission
	*/
	multiplexedInfoChar := s.AddCharacteristic(attrMultiplexedInfoCharacteristicsUUID)
	multiplexedInfoChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		data := make([]byte, 20)
		rsp.Write(data)
		log.Println("[[Rowing]] Multiplexed Info Char")
	})

	multiplexedInfoChar.HandleNotifyFunc(func(r gatt.Request, n gatt.Notifier) {
		log.Println("[[Rowing]] Multiplex Info Char Notify Func")
		//generate a workout detail payload here
		data := make([]byte, 20)
		data=append(data,[]byte{0x31,0x01,0x80,0x01,0x01,0x80,0x01,0x01,0x80,0x01}...)
		n.Write(data)
	})

	multiplexedInfoChar.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})

	return s
}
