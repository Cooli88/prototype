package dto

import (
	"device-listener-go/parse"
	"fmt"
)

const (
	Acc         = 1
	Door        = 2
	Analog      = 4
	Gsm         = 5
	Speed       = 6
	Voltage     = 7
	GPSPower    = 8
	Temperature = 9
	Odometer    = 16
	Stop        = 20
	Trip        = 28
	Immobilizer = 29
	Authorized  = 30
	Greedriving = 31
	Overspeed   = 33
	Pdop        = 181
	Hdop        = 182
)

type DataToRabbit struct {
	Imei        string  `json:"imei"`
	Priority    int64   `json:"priority"`
	UtcDateTime int64   `json:"utcDateTime"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	Direction   int64   `json:"direction"`
	Ignition    int64   `json:"ignition"`
	Altitude    int64   `json:"altitude"`
	Satellite   int64   `json:"satellite"`
	Heading     int64   `json:"heading"`
	Speed       int64   `json:"speed"`
	Acc         int64   `json:"acc"`
	Door        int64   `json:"door"`
	Gsm         int64   `json:"gsm"`
	IsStop      int64   `json:"isStop"`
	Mileage     int64   `json:"mileage"`
	IoCount     int64   `json:"ioCount"`
	IoEvent     int64   `json:"ioEvent"`
	Pdop        int64   `json:"pdop"`
	Hdop        int64   `json:"hdop"`
	CustomIds   string  `json:"customIds"`
}

type DataToRabbitArr []DataToRabbit

func CreateDataToRabbitFromRecords(records parse.Data, deviceImei string) DataToRabbitArr {
	dataToRabbitArr := []DataToRabbit{}
	for _, record := range records.Records {
		dataToRabbit := CreateDataToRabbitFromRecord(deviceImei, record)
		dataToRabbitArr = append(dataToRabbitArr, dataToRabbit)
		fmt.Printf("record: %+v\n", record)
		fmt.Printf("dataToRabbit: %+v\n", dataToRabbit)
	}
	return dataToRabbitArr
}

func CreateDataToRabbitFromRecord(imei string, record parse.Record) DataToRabbit {
	dataToRabbit := DataToRabbit{
		Imei:        imei,
		Priority:    record.Priority,
		UtcDateTime: record.TimestampMilliseconds,
		Longitude:   float32(record.GpsData.Longitude) / 10000000.0,
		Latitude:    float32(record.GpsData.Latitude) / 10000000.0,
		Altitude:    record.GpsData.Altitude,
		Satellite:   record.GpsData.VisibleSattelites,
		Speed:       record.GpsData.Speed,
		IoEvent:     record.IOElementIDEventGenerated,
		IoCount:     record.IOElementCount,
	}

	dataToRabbit.setDirectionFromAngle(record.GpsData.Angle)
	dataToRabbit.setIoElements(record)

	return dataToRabbit
}

func (dataToRabbit *DataToRabbit) setDirectionFromAngle(gpsAngle int64) {
	var direction int64

	if gpsAngle < 90 {
		direction = 1
	} else if gpsAngle == 90 {
		direction = 2
	} else if gpsAngle < 180 {
		direction = 3
	} else if gpsAngle == 180 {
		direction = 4
	} else if gpsAngle < 270 {
		direction = 5
	} else if gpsAngle == 270 {
		direction = 6
	} else if gpsAngle > 270 {
		direction = 7
	}

	dataToRabbit.Direction = direction
}

func (dataToRabbit *DataToRabbit) setIoElements(record parse.Record) {
	dataToRabbit.setAcc(record)
	dataToRabbit.setDoor(record)
	dataToRabbit.setGsm(record)
	dataToRabbit.setStop(record)
	dataToRabbit.setPdop(record)
	dataToRabbit.setHdop(record)
	dataToRabbit.setMileage(record)
	dataToRabbit.setCustomIds(record)
}

func (dataToRabbit *DataToRabbit) setAcc(record parse.Record) {
	if _, ok := record.IOElements[Acc]; ok {
		switch record.IOElements[Acc] {
		case 1:
			dataToRabbit.Acc = 0
		default:
			dataToRabbit.Acc = 1
		}
	}
}

func (dataToRabbit *DataToRabbit) setDoor(record parse.Record) {
	if _, ok := record.IOElements[Door]; ok {
		switch record.IOElements[Door] {
		case 1:
			dataToRabbit.Door = 0
		default:
			dataToRabbit.Door = 1
		}
	}
}

func (dataToRabbit *DataToRabbit) setGsm(record parse.Record) {
	if _, ok := record.IOElements[Gsm]; ok {
		dataToRabbit.Gsm = record.IOElements[Gsm]
	}
}

func (dataToRabbit *DataToRabbit) setStop(record parse.Record) {
	if _, ok := record.IOElements[Stop]; ok {
		switch record.IOElements[Stop] {
		case 1:
			dataToRabbit.IsStop = 0
		default:
			dataToRabbit.IsStop = 1
		}
	}
}

func (dataToRabbit *DataToRabbit) setPdop(record parse.Record) {
	if _, ok := record.IOElements[Pdop]; ok {
		dataToRabbit.Pdop = record.IOElements[Pdop]
	}
}

func (dataToRabbit *DataToRabbit) setHdop(record parse.Record) {
	if _, ok := record.IOElements[Hdop]; ok {
		dataToRabbit.Hdop = record.IOElements[Hdop]
	}
}

func (dataToRabbit *DataToRabbit) setMileage(record parse.Record) {
	if _, ok := record.IOElements[Odometer]; ok {
		dataToRabbit.Mileage = record.IOElements[Odometer]
	}
}

func (dataToRabbit *DataToRabbit) setCustomIds(record parse.Record) {
	if len(record.IOElements) > 0 {
		//todo пока что оно всегда было пустым
		dataToRabbit.CustomIds = ""
	}
}
