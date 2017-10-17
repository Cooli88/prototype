package parse

import (
    "errors"
    "fmt"
    "encoding/binary"
    "device-listener-go-new-device/clienttest"
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

//var triplekey []byte = []byte{225, 183, 161, 139, 9, 19, 247, 142, 192, 137, 142, 214, 40, 45, 216, 211, 21, 0, 63, 107, 253, 164, 105, 110}
var triplekey []byte = []byte{0xE1, 0xB7, 0xA1, 0x8B, 0x09, 0x13, 0xF7, 0x8E, 0xC0, 0x89, 0x8E, 0xD6, 0x28, 0x2D, 0xD8, 0xD3, 0x15, 0x00, 0x3F, 0x6B, 0xFD, 0xA4, 0x69, 0x6E}

type DataBuffer struct {
    DataRaw    []byte
    Data       []byte
    imei       string
    cursorLast int
    Error      error
}

func newDataBuffer(data []byte) DataBuffer {
    dataBuffer := DataBuffer{DataRaw: data, cursorLast: 0}
    if len(data) < 1 {
        dataBuffer.Error = errors.New("data is empty")
    } else if len(data) < 17 {
        dataBuffer.Error = errors.New("data is broken")
    }

    headerFrame := dataBuffer.GetFrameHeaderBytes()
    messageLength := dataBuffer.GetMessageLength()
    dataBuffer.imei = dataBuffer.GetImei()

    dataBuffer.SetData()
    frameType := dataBuffer.FrameType()
    frameNumber := dataBuffer.FrameNumber()
    dataLength := dataBuffer.DataLength()
    messageGetProcolV := dataBuffer.MessageGetProcolV()

    fmt.Printf("headerFrame: %+v\n", headerFrame)
    fmt.Printf("MessageLength: %+v\n", messageLength)
    fmt.Printf("imei: %+v\n", dataBuffer.imei)
    fmt.Printf("frameType: %+v\n", frameType)
    fmt.Printf("frameNumber: %+v\n", frameNumber)
    fmt.Printf("dataLength: %+v\n", dataLength)
    fmt.Printf("dataBuffer.Data: %+v\n", dataBuffer.Data)
    fmt.Printf("messageGetProcolV: %+v\n", messageGetProcolV)

    panic(1)

    return dataBuffer
}

func byteSliceToInt64(sizeByte int, mySlice []byte) int64 {
    var value int64
    switch sizeByte {
    case oneByte:
        value = int64(uint16(mySlice[0]))
    case twoByte:
        valueRaw := binary.BigEndian.Uint16(mySlice)
        value = int64(valueRaw)
    case fourByte:
        valueRaw := binary.BigEndian.Uint32(mySlice)
        value = int64(valueRaw)
    default:
        //eightByte
        valueRaw := binary.BigEndian.Uint64(mySlice)
        value = int64(valueRaw)
    }
    return value
}

//перевод из base16 в base10 (9 = 9, f = 16)
func (dataBuffer *DataBuffer) getDataRaw(sizeByte int) int64 {
    if dataBuffer.Error != nil {
        return 0
    }
    mySlice := dataBuffer.DataRaw[dataBuffer.cursorLast:dataBuffer.cursorLast+sizeByte]
    dataBuffer.cursorLast += sizeByte
    return byteSliceToInt64(sizeByte, mySlice)
}

//перевод из base16 в base10 (9 = 9, f = 16)
func (dataBuffer *DataBuffer) getData(sizeByte int) int64 {
    if dataBuffer.Error != nil {
        return 0
    }
    mySlice := dataBuffer.Data[dataBuffer.cursorLast:dataBuffer.cursorLast+sizeByte]
    dataBuffer.cursorLast += sizeByte
    return byteSliceToInt64(sizeByte, mySlice)
}

func (dataBuffer *DataBuffer) getDataRawString(sizeByte int) string {
    if dataBuffer.Error != nil {
        return ""
    }
    mySlice := dataBuffer.DataRaw[dataBuffer.cursorLast:dataBuffer.cursorLast+sizeByte]
    dataBuffer.cursorLast += sizeByte
    return string(mySlice)
}

//перевод из base16 в base10 (9 = 9, f = 16)
func (dataBuffer *DataBuffer) getDataWithoutnCursor(start, sizeByte int) int64 {
    if dataBuffer.Error != nil {
        return 0
    }
    mySlice := dataBuffer.Data[start:start+sizeByte]
    return byteSliceToInt64(sizeByte, mySlice)
}

func (dataBuffer *DataBuffer) GetFrameHeaderBytes() int64 {
    return dataBuffer.getDataRaw(2)
}

func (dataBuffer *DataBuffer) GetMessageLength() int64 {
    return dataBuffer.getDataRaw(2)
}

func (dataBuffer *DataBuffer) GetImei() string {
    return dataBuffer.getDataRawString(15)
}

func (dataBuffer *DataBuffer) SetData() {
    dataToDecrypt := dataBuffer.DataRaw[dataBuffer.cursorLast:len(dataBuffer.DataRaw)-2]

    //fmt.Printf("dataToDecrypt: %x\n", dataToDecrypt)
    dataBuffer.Data = clienttest.Decrypt(dataToDecrypt)
    dataBuffer.cursorLast = 0
    //var s = "0100010040000210C21C4D43555F544D4F564D363230305356302E302E3042303854313030351A544D4F5F55535F5A3632303056312E302E3042303754303932390000005D15B6128F"
    //dataFromDevice, _ := hex.DecodeString(s)
    //dataBuffer.Data = dataFromDevice
}

func (dataBuffer *DataBuffer) FrameType() int64 {
    return dataBuffer.getData(1)
}

func (dataBuffer *DataBuffer) FrameNumber() int64 {
    return dataBuffer.getData(2)
}

func (dataBuffer *DataBuffer) DataLength() int64 {
    return dataBuffer.getData(2)
}

func (dataBuffer *DataBuffer) MessageGetProcolV() int64 {
    return dataBuffer.getData(1)
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
