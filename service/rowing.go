package service

import (
	"pm5-emulator/service/mux"
	"github.com/sirupsen/logrus"
	"github.com/bettercap/gatt"
	"time"
	"crypto/rand"
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

	rowingGenStatusChar.HandleNotifyFunc(
		func(r gatt.Request, n gatt.Notifier) {
			logrus.Info("General Status Char Notify Request - launching goroutine")
			go func() {
				for true {
					logrus.Info("Sending General Status Char Notification from goroutine")
					byteArray := make([]byte, 1)
					rand.Read(byteArray)		
					// 19 bytes		
					n.Write([]byte{byteArray[0], 0x5, 0x5, 0x5, 0x5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5, 0x5, 0x5, 0x5})
					time.Sleep(500 * time.Millisecond)
				}
			}()
		})	

	/*
		C2 rowing additional status 1 characteristic
	*/
	additionalStatus1Char := s.AddCharacteristic(attrAdditionalStatus1CharacteristicsUUID)
	additionalStatus1Char.HandleNotifyFunc(func(r gatt.Request, n gatt.Notifier) {
		logrus.Info("Additional Status 1 Char Notify Request - launching goroutine")
		go func() {
			for true {
				logrus.Info("Sending Additional Status 1 Notification from goroutine")				
				n.Write([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xb8, 0xb, 0x0, 0x0, 0x0})
				time.Sleep(500 * time.Millisecond)
			}
		}()	
	})

	/*
		C2 rowing additional status 2 characteristic
	*/
	additionalStatus2Char := s.AddCharacteristic(attrAdditionalStatus2CharacteristicsUUID)
	additionalStatus2Char.HandleNotifyFunc(func(r gatt.Request, n gatt.Notifier) {
		logrus.Info("Additional Status 2 Char Notify Request - launching goroutine")
		go func() {
			for true {
				logrus.Info("Sending Additional Status 2 Notification from goroutine")
				byteArray := make([]byte, 1)
				rand.Read(byteArray)				
				n.Write([]byte{byteArray[0], 0x1, 0x2, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})
				time.Sleep(500 * time.Millisecond)
			}
		}()	
	})

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
	strokeDataChar.HandleNotifyFunc(func(r gatt.Request, n gatt.Notifier) {
		logrus.Info("Stroke Data Char Notify Request - launching goroutine")
		go func() {
			for true {
				logrus.Info("Stroke Data Notification from goroutine")
				n.Write([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})
				time.Sleep(1000 * time.Millisecond)
			}
		}()	
	})

	/*
		C2 rowing additional stroke data characteristic 0x0036
	*/
	additionalStrokeDataChar := s.AddCharacteristic(attrAdditionalStrokeDataCharacteristicsUUID)
	additionalStrokeDataChar.HandleNotifyFunc(func(r gatt.Request, n gatt.Notifier) {
		logrus.Info("Additional Stroke Data Char Notify Request - launching goroutine")
		go func() {
			for true {
				logrus.Info("Stroke Data Notification from goroutine")
				n.Write([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff})
				time.Sleep(1000 * time.Millisecond)
			}
		}()	
	})

	/*
		C2 rowing split/interval data characteristic
	*/
	splitIntervalDataChar := s.AddCharacteristic(attrSplitIntervalDataCharacteristicsUUID)
	splitIntervalDataChar.HandleNotifyFunc(func(r gatt.Request, n gatt.Notifier) {
		logrus.Info("Split/Interval Data Char Notify Request - launching goroutine")
		go func() {
			for true {
				logrus.Info("Split/Interval Data Notification from goroutine")
				n.Write(make([]byte, 18))
				time.Sleep(50000 * time.Millisecond)
			}
		}()	
	})


	/*
		C2 rowing additional split/interval data characteristic
	*/
	additionalSplitIntervalDataChar := s.AddCharacteristic(attrAdditionalSplitIntervalDataCharacteristicsUUID)
	additionalSplitIntervalDataChar.HandleNotifyFunc(func(r gatt.Request, n gatt.Notifier) {
		logrus.Info("Additional Split/Interval Data Char Notify Request - launching goroutine")
		go func() {
			for true {
				logrus.Info("Additional Split/Interval Data Notification from goroutine")
				n.Write(make([]byte, 18))
				time.Sleep(50000 * time.Millisecond)
			}
		}()	
	})

	/*
		C2 rowing end of workout summary data characteristic
	*/
	endOfWorkoutSummaryDataChar := s.AddCharacteristic(attrEndOfWorkoutSummaryDataCharacteristicsUUID)
	endOfWorkoutSummaryDataChar.HandleNotifyFunc(func(r gatt.Request, n gatt.Notifier) {
		logrus.Info("End of workout summary Data Char Notify Request - launching goroutine")
		go func() {
			for true {
				time.Sleep(200000 * time.Millisecond)
				logrus.Info("End of workout summary Data Notification from goroutine")
				n.Write(make([]byte, 20))
			}
		}()	
	})

	/*
		C2 rowing end of workout additional summary data characteristic
	*/
	additionalEndOfWorkoutSummaryDataChar := s.AddCharacteristic(attrAdditionalEndOfWorkoutSummaryDataCharacteristicsUUID)
	additionalEndOfWorkoutSummaryDataChar.HandleNotifyFunc(func(r gatt.Request, n gatt.Notifier) {
		logrus.Info("End of workout Additional summary Data Char Notify Request - launching goroutine")
		go func() {
			for true {
				time.Sleep(200000 * time.Millisecond)
				logrus.Info("End of workout Additional summary Data Notification from goroutine")
				n.Write(make([]byte, 20))
			}
		}()	
	})


	/*
		C2 rowing heart rate belt information characteristic
	*/
	heartRateBeltInfoChar := s.AddCharacteristic(attrHeartRateBeltInfoCharacteristicsUUID)
	heartRateBeltInfoChar.HandleNotifyFunc(func(r gatt.Request, n gatt.Notifier) {
		logrus.Info("Heart Rate Belt Info Char Notify Request - launching goroutine")
		go func() {
			for true {
				logrus.Info("Heart Rate Belt Data Notification from goroutine")
				n.Write(make([]byte, 6))
				time.Sleep(100000 * time.Millisecond)

			}
		}()	
	})

	/*
		C2 force curve data characteristic
	*/
	forceCurveDataChar := s.AddCharacteristic(attrForceCurveDataCharacteristicsUUID)
	forceCurveDataChar.HandleNotifyFunc(func(r gatt.Request, n gatt.Notifier) {
		logrus.Info("Force Curve Data Char Notify Request - launching goroutine")
		go func() {
			for true {
				logrus.Info("Force Curve Data Notification from goroutine")
				n.Write([]byte{0b000101001, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})
				time.Sleep(1000 * time.Millisecond)
			}
		}()	
	})

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
