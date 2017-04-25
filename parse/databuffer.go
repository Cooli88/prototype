package parse

import (
        "errors"
        "fmt"
        "encoding/binary"
)

const (
        countSymbolsInByte = 2

        iOElementsPartsCount = 4

        iOElementsPartFirstSizeByte  = 1
        iOElementsPartSecondSizeByte = 2
        iOElementsPartThirdSizeByte  = 4
        iOElementsPartFourthSizeByte = 8

        oneByte   = 1
        twoByte   = 2
        fourByte  = 4
        eightByte = 8
)

type DataBuffer struct {
        Data       []byte
        cursorLast int
        Error      error
}

func newDataBuffer(data []byte) DataBuffer {
        dataBuffer := DataBuffer{Data: data, cursorLast: 0}
        if len(data) < 1 {
                dataBuffer.Error = errors.New("data is empty")
        } else if len(data) < 17 {
                dataBuffer.Error = errors.New("data is broken")
        }

        dataBuffer.GetFirstFourByte()
        dataLength := dataBuffer.GetDataLength()
        crc := dataBuffer.GetCrc()

        fmt.Printf("dataLength: %+v\n", dataLength)
        fmt.Printf("crc: %+v\n", crc)

        return dataBuffer
}

//перевод из base16 в base10 (9 = 9, f = 16)
func (dataBuffer *DataBuffer) getData(sizeByte int) int64 {
        if (dataBuffer.Error != nil) {
                return 0
        }
        mySlice := dataBuffer.Data[dataBuffer.cursorLast:dataBuffer.cursorLast+sizeByte]
        var value int64
        if (sizeByte == oneByte) {
                value = int64(uint16(mySlice[0]))
        } else if (sizeByte == twoByte) {
                valueRaw := binary.BigEndian.Uint16(mySlice)
                value = int64(valueRaw)
        } else if (sizeByte == fourByte) {
                valueRaw := binary.BigEndian.Uint32(mySlice)
                value = int64(valueRaw)
        } else {
                valueRaw := binary.BigEndian.Uint64(mySlice)
                value = int64(valueRaw)
        }
        dataBuffer.cursorLast += sizeByte
        return value
}

//перевод из base16 в base10 (9 = 9, f = 16)
func (dataBuffer *DataBuffer) getDataWithoutnCursor(start, sizeByte int) int64 {
        if (dataBuffer.Error != nil) {
                return 0
        }
        mySlice := dataBuffer.Data[start:start+sizeByte]
        var value int64
        if (sizeByte <= 2) {
                valueRaw := binary.BigEndian.Uint16(mySlice)
                value = int64(valueRaw)
        } else if (sizeByte <= 4) {
                valueRaw := binary.BigEndian.Uint32(mySlice)
                value = int64(valueRaw)
        } else {
                valueRaw := binary.BigEndian.Uint64(mySlice)
                value = int64(valueRaw)
        }
        return value
}

func (dataBuffer *DataBuffer) GetFirstFourByte() int64 {
        return dataBuffer.getData(4)
}

func (dataBuffer *DataBuffer) GetDataLength() int64 {
        return dataBuffer.getData(4)
}

func (dataBuffer *DataBuffer) GetCrc() int64 {
        start := len(dataBuffer.Data) - 4
        return dataBuffer.getDataWithoutnCursor(start, 4)
}

func (dataBuffer *DataBuffer) GetCodecID() int64 {
        return dataBuffer.getData(1)
}

func (dataBuffer *DataBuffer) GetNumberOfData() int64 {
        return dataBuffer.getData(1)
}

func (dataBuffer *DataBuffer) GetTimestampMilliseconds() int64 {
        return dataBuffer.getData(8)
}

func (dataBuffer *DataBuffer) GetPriority() int64 {
        return dataBuffer.getData(1)
}

//todo float maybe
func (dataBuffer *DataBuffer) GetGpsLatitude() int64 {
        return dataBuffer.getData(4)
}

//todo float maybe
func (dataBuffer *DataBuffer) GetGpsLongitude() int64 {
        return dataBuffer.getData(4)
}

func (dataBuffer *DataBuffer) GetGpsAltitude() int64 {
        return dataBuffer.getData(2)
}

func (dataBuffer *DataBuffer) GetGpsAngle() int64 {
        return dataBuffer.getData(2)
}

func (dataBuffer *DataBuffer) GetGpsVisibleSattelites() int64 {
        return dataBuffer.getData(1)
}

func (dataBuffer *DataBuffer) GetGpsSpeed() int64 {
        return dataBuffer.getData(2)
}

func (dataBuffer *DataBuffer) GetIOElementIDofEventGenerated() int64 {
        return dataBuffer.getData(1)
}

func (dataBuffer *DataBuffer) GetCountAllIOElementsForRecord() int64 {
        return dataBuffer.getData(1)
}

func (dataBuffer *DataBuffer) GetCountIOElementsNext() int64 {
        return dataBuffer.getData(1)
}

func (dataBuffer *DataBuffer) GetIOElementId() int64 {
        return dataBuffer.getData(1)
}

func (dataBuffer *DataBuffer) GetIOElementValueWithSize(countByte int) int64 {
        return dataBuffer.getData(countByte)
}
