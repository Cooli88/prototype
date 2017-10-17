package main

import (
        "go.uber.org/zap"
        "device-listener-go-new-device/parse"
        "fmt"
        //"device-listener-go-new-device/clienttest"
        "encoding/hex"
)

func main() {
        logger, _ := zap.NewProduction()
        sugar := logger.Sugar()
        sugar.Infow("App start.", )
        //dataFromDevice := []byte{85, 85, 0, 37, 56, 54, 50, 54, 48, 57, 48, 48, 48, 49, 50, 55, 54, 54, 55, 43, 94, 130, 255, 202, 206, 101, 29, 141, 130, 149, 80, 119, 133, 250, 22, 170, 170}
        //dataFromDevice := []byte{85, 85,
        //0, 37,
        //56, 54, 50, 54, 48, 57, 48, 48, 48, 49, 50, 55, 54, 54, 55,
        //176,
        //185, 204,
        //26, 193,
        //33, 126, 25, 207, 140, 3, 33, 66, 217, 127, 19, 170, 170}
        //const s = "555500653836313235313033303030393333334481463F374D1FD9D485F4DD50CC39D9896C89DB98813E74753BB830BDCAAC24ED20816645AF3332DB53D45CE7DB870785BF90EFCD421A85B93BC8FE7430CFF9B7AA8B4C45028C6113EBE1AD4FBB9E1FAAAA"
        //Зашифрованные входные
        const s = "555500653836313235313033303030393333334481463F374D1FD9D485F4DD50CC39D9896C89DB98813E74753BB830BDCAAC24ED20816645AF3332DB53D45CE7DB870785BF90EFCD421A85B93BC8FE7430CFF9B7AA8B4C45028C6113EBE1AD4FBB9E1FAAAA"
        //const s = "5555002d383636363939303234333937383330030003000900010156f5f8120000560305b4ffffffffffffaaaa"
        dataFromDevice, _ := hex.DecodeString(s)

        fmt.Printf("dataFromDevice: %x\n", dataFromDevice)

        //clienttest.Decrypt()
        parse.CreateData(dataFromDevice)
        //tcpserver.ServerStart()
}

//555500653836313235313033303030393333330100010040000210C21C4D43555F544D4F564D363230305356302E302E3042303854313030351A544D4F5F55535F5A3632303056312E302E3042303754303932390000005D15B6128FAAAA
