package parse

import (
        "errors"
        "fmt"
)

type Data struct {
        Records      []Record `json:"records"`
        NumberOfData int64    `json:"number_of_data"`
        Error        error
}

type Record struct {
        TimestampMilliseconds     int64 `json:"timestamp_milliseconds"`
        Priority                  int64 `json:"priority"`
        GpsData                   GPSData `json:"gps"`
        IOElementIDEventGenerated int64 `json:"io_element_id_event_generated"`
        IOElementCount            int64 `json:"io_element_count"`
        IOElements                map[int64]int64 `json:"io_elements"`
}

type GPSData struct {
        Latitude          int64 `json:"latitude"`
        Longitude         int64 `json:"longitude"`
        Altitude          int64 `json:"altitude"`
        Angle             int64 `json:"angle"`
        VisibleSattelites int64 `json:"visible_sattelites"`
        Speed             int64 `json:"speed"`
}

func newRecord() Record {
        return Record{IOElements: make(map[int64]int64), GpsData: GPSData{}}
}

//CreateData !!!НЕЛЬЗЯ МЕНЯТЬ ПОРЯДОК ВЫЗОВА ФУНКЦИЙ DataBuffer, каждую функцию необходимо вызвать в любом случае, необходимо для правильного сдвига счетчика байт
func CreateData(data []byte) Data {
        dataBuffer := newDataBuffer(data)
        dataBuffer.GetCodecID()
        numberOfData := dataBuffer.GetNumberOfData()

        allRecords := Data{}
        allRecords.NumberOfData = numberOfData

        if dataBuffer.Error != nil {
                allRecords.Error = dataBuffer.Error
                return allRecords
        }

        fmt.Printf("number of data: %+v\n", numberOfData)

        for countRecord := 0; countRecord < int(numberOfData); countRecord++ {
                record := newRecord()
                record.TimestampMilliseconds = dataBuffer.GetTimestampMilliseconds()
                record.Priority = dataBuffer.GetPriority()
                record.GpsData.Longitude = dataBuffer.GetGpsLongitude()
                record.GpsData.Latitude = dataBuffer.GetGpsLatitude()
                record.GpsData.Altitude = dataBuffer.GetGpsAltitude()
                record.GpsData.Angle = dataBuffer.GetGpsAngle()
                record.GpsData.VisibleSattelites = dataBuffer.GetGpsVisibleSattelites()
                record.GpsData.Speed = dataBuffer.GetGpsSpeed()
                record.IOElementIDEventGenerated = dataBuffer.GetIOElementIDofEventGenerated()

                record.IOElementCount = dataBuffer.GetCountAllIOElementsForRecord()

                loadAllIoElementsData(&dataBuffer, &record)

                allRecords.Records = append(allRecords.Records, record)
        }

        numberOfDataConfirm := dataBuffer.GetNumberOfData()

        if (dataBuffer.Error != nil) {
                allRecords.Error = dataBuffer.Error
                return allRecords
        }

        if (numberOfData != numberOfDataConfirm) {
                fmt.Printf("numberOfDataConfirm: %+v\n", numberOfDataConfirm)
                allRecords.Error = errors.New("count record not equal")
        }

        return allRecords
}

func loadAllIoElementsData(dataBuffer *DataBuffer, record *Record) {
        for countPartIO := 0; countPartIO < iOElementsPartsCount; countPartIO++ {

                var sizeByte int

                switch countPartIO {
                case 0:
                        sizeByte = iOElementsPartFirstSizeByte
                case 1:
                        sizeByte = iOElementsPartSecondSizeByte
                case 2:
                        sizeByte = iOElementsPartThirdSizeByte
                case 3:
                        sizeByte = iOElementsPartFourthSizeByte
                }

                countIoElementsInPart := dataBuffer.GetCountIOElementsNext()
                if (countIoElementsInPart > 0) {
                        for IoElements := 0; IoElements < int(countIoElementsInPart); IoElements++ {
                                record.IOElements[dataBuffer.GetIOElementId()] = dataBuffer.GetIOElementValueWithSize(sizeByte)
                        }
                }
        }
}
