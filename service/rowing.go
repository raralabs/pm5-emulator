package service

import (
	"pm5-emulator/service/mux"
	"github.com/sirupsen/logrus"
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

//NewRowingService advertises rowing service defined by PM5 device
func NewRowingService() *gatt.Service {
	s := gatt.NewService(attrRowingServiceUUID)

	/*
		C2 rowing general status characteristic
	*/
	rowingGenStatusChar := s.AddCharacteristic(attrGeneralStatusCharacteristicsUUID)
	rowingGenStatusChar.HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			logrus.Info("General Status Char Read Request")
			//19 bytes
			rsp.Write([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80, 0x0})
		})

	rowingGenStatusChar.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})

	/*
		C2 rowing additional status 1 characteristic
	*/
	additionalStatus1Char := s.AddCharacteristic(attrAdditionalStatus1CharacteristicsUUID)
	additionalStatus1Char.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Additional Status 1 Char Read Request")
		rsp.Write([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xb8, 0xb, 0x0, 0x0, 0x0})
	})

	additionalStatus1Char.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})

	/*
		C2 rowing additional status 2 characteristic
	*/
	additionalStatus2Char := s.AddCharacteristic(attrAdditionalStatus2CharacteristicsUUID)
	additionalStatus2Char.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Additional Status 2 Status Char Read Request")
		rsp.Write([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})
	})

	additionalStatus2Char.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})
	/*
		C2 rowing general status and additional status sample rate characteristic 0x0034
	*/
	sampleRateChar := s.AddCharacteristic(attrSampleRateCharacteristicsUUID)
	sampleRateChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Sample Rate Char Read Request")
		data := make([]byte, 1)
		rsp.Write(data)
	})

	/*
		C2 rowing stroke data  characteristic 0x0035
	*/
	strokeDataChar := s.AddCharacteristic(attrStrokeDataCharacteristicsUUID)
	strokeDataChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Stroke Data char Read Request")
		rsp.Write([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})
	})

	strokeDataChar.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})

	/*
		C2 rowing additional stroke data characteristic 0x0036
	*/
	additionalStrokeDataChar := s.AddCharacteristic(attrAdditionalStrokeDataCharacteristicsUUID)
	additionalStrokeDataChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Additional Stroke Data char Read Request")
		rsp.Write([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff})
	})

	additionalStrokeDataChar.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})

	/*
		C2 rowing split/interval data characteristic
	*/
	splitIntervalDataChar := s.AddCharacteristic(attrSplitIntervalDataCharacteristicsUUID)
	splitIntervalDataChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Split/Interval Data char Read Request")
		data := make([]byte, 18)
		rsp.Write(data)
	})

	splitIntervalDataChar.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})

	/*
		C2 rowing additional split/interval data characteristic
	*/
	additionalSplitIntervalDataChar := s.AddCharacteristic(attrAdditionalSplitIntervalDataCharacteristicsUUID)
	additionalSplitIntervalDataChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Additional Split/Interval Data char Read Request")
		data := make([]byte, 18)
		rsp.Write(data)
	})

	additionalSplitIntervalDataChar.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})
	/*
		C2 rowing end of workout summary data characteristic
	*/
	endOfWorkoutSummaryDataChar := s.AddCharacteristic(attrEndOfWorkoutSummaryDataCharacteristicsUUID)
	endOfWorkoutSummaryDataChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("End of workout summary Data char Read Request")
		data := make([]byte, 20)
		rsp.Write(data)
	})

	endOfWorkoutSummaryDataChar.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})
	/*
		C2 rowing end of workout additional summary data characteristic
	*/
	additionalEndOfWorkoutSummaryDataChar := s.AddCharacteristic(attrAdditionalEndOfWorkoutSummaryDataCharacteristicsUUID)
	additionalEndOfWorkoutSummaryDataChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Additional End of workout summary Data char Read Request")
		data := make([]byte, 19)
		rsp.Write(data)
	})

	additionalEndOfWorkoutSummaryDataChar.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})

	/*
		C2 rowing heart rate belt information characteristic
	*/
	heartRateBeltInfoChar := s.AddCharacteristic(attrHeartRateBeltInfoCharacteristicsUUID)
	//handle read
	heartRateBeltInfoChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Heart Rate Belt Info Read Request")
		data := make([]byte, 6)
		rsp.Write(data)
	})

	//handle write
	heartRateBeltInfoChar.HandleWriteFunc(func(req gatt.Request, data []byte) (status byte) {
		logrus.Info("received data at heart rate belt info: ", string(data))
		return gatt.StatusSuccess
	})

	heartRateBeltInfoChar.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})

	/*
		C2 force curve data characteristic
	*/
	forceCurveDataChar := s.AddCharacteristic(attrForceCurveDataCharacteristicsUUID)
	forceCurveDataChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Force Curve Data Read Request")
		data := make([]byte, 288)
		rsp.Write(data)
	})

	forceCurveDataChar.HandleWriteFunc(func(req gatt.Request, data []byte) (status byte) {
		logrus.Info("received data at force curve data: ", string(data))
		return gatt.StatusSuccess
	})

	/*
		C2 multiplexed information 	characteristic

		0x0080 | Up to 20 bytes | READ Permission
	*/
	multiplexedInfoChar := s.AddCharacteristic(attrMultiplexedInfoCharacteristicsUUID)
	multiplexedInfoChar.HandleReadFunc(func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
		logrus.Info("Multiplexed Info Char")
		m:=mux.Multiplexer{}
		rsp.Write(m.HandleC2RowingGeneralStatus([]byte{}))
	})

	multiplexedInfoChar.HandleNotifyFunc(func(r gatt.Request, n gatt.Notifier) {
		logrus.Info("Multiplex Info Char Notify Func")
		//generate a rowing general status payload here
		m:=mux.Multiplexer{}
		n.Write(m.HandleC2RowingGeneralStatus([]byte{}))
	})

	multiplexedInfoChar.AddDescriptor(attrGeneralStatusDescriptorUUID).SetValue([]byte{})

	return s
}
