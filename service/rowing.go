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


	/*
		C2 rowing additional status 1 characteristic
	*/
	additionalStatus1Char := s.AddCharacteristic(attrAdditionalStatus1CharacteristicsUUID)


	/*
		C2 rowing additional status 2 characteristic
	*/
	additionalStatus2Char := s.AddCharacteristic(attrAdditionalStatus2CharacteristicsUUID)

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


	/*
		C2 rowing additional stroke data characteristic 0x0036
	*/
	additionalStrokeDataChar := s.AddCharacteristic(attrAdditionalStrokeDataCharacteristicsUUID)


	/*
		C2 rowing split/interval data characteristic
	*/
	splitIntervalDataChar := s.AddCharacteristic(attrSplitIntervalDataCharacteristicsUUID)


	/*
		C2 rowing additional split/interval data characteristic
	*/
	additionalSplitIntervalDataChar := s.AddCharacteristic(attrAdditionalSplitIntervalDataCharacteristicsUUID)

	/*
		C2 rowing end of workout summary data characteristic
	*/
	endOfWorkoutSummaryDataChar := s.AddCharacteristic(attrEndOfWorkoutSummaryDataCharacteristicsUUID)

	/*
		C2 rowing end of workout additional summary data characteristic
	*/
	additionalEndOfWorkoutSummaryDataChar := s.AddCharacteristic(attrAdditionalEndOfWorkoutSummaryDataCharacteristicsUUID)


	/*
		C2 rowing heart rate belt information characteristic
	*/
	heartRateBeltInfoChar := s.AddCharacteristic(attrHeartRateBeltInfoCharacteristicsUUID)

	/*
		C2 force curve data characteristic
	*/
	forceCurveDataChar := s.AddCharacteristic(attrForceCurveDataCharacteristicsUUID)

	/*
		C2 multiplexed information 	characteristic

		0x0080 | Up to 20 bytes | READ Permission
	*/
	multiplexedInfoChar := s.AddCharacteristic(attrMultiplexedInfoCharacteristicsUUID)

	multiplexedInfoChar.HandleNotifyFunc(func(r gatt.Request, n gatt.Notifier) {
		logrus.Info("Multiplex Info Char Notify Func")
		//generate a rowing general status payload here
		m:=mux.Multiplexer{}
		n.Write(m.HandleC2RowingGeneralStatus([]byte{}))
	})


	return s
}
