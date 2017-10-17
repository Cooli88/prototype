package clienttest

import (
    "fmt"
    //"destest"
    "destest"
)

//Верный
var triplekey []byte = []byte{
    0xE1, 0xB7, 0xA1, 0x8B, 0x09, 0x13, 0xF7, 0x8E,
    0xC0, 0x89, 0x8E, 0xD6, 0x28, 0x2D, 0xD8, 0xD3,
    0x15, 0x00, 0x3F, 0x6B, 0xFD, 0xA4, 0x69, 0x6E,
}
//var triplekey []byte = []byte{
//    0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0x12,
//    0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0x11,
//    0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0x01, 0x33,
//}
//var triplekey []byte = []byte{0xE1, 0xB7, 0xA1, 0x8B, 0x09, 0x13, 0xF7, 0x8E, 0xC0, 0x89, 0x8E, 0xD6, 0x28, 0x2D, 0xD8, 0xD3, 0x15, 0x00, 0x3F, 0x6B, 0xFD, 0xA4, 0x69, 0x6E}

func Decrypt(plaintext []byte) []byte {

    c, _ := destest.NewTripleDESCipher(triplekey)
    //
    //out  := make([]byte, 0)
    //out1 := make([]byte, 8)
    var tmp []byte = []byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38}

    outputNew := make([]byte, 8)
    X := GET_UINT32_BE(triplekey, 0)
    Y := GET_UINT32_BE(triplekey, 4)

    outputNew = PUT_UINT32_BE(Y, outputNew, 0)
    fmt.Printf("tmp %v\n", outputNew)
    outputNew = PUT_UINT32_BE(X, outputNew, 4)

    fmt.Printf("resultDes3Enc %v\n", outputNew)
    fmt.Printf("resultDes3Dec %v\n", X)
    fmt.Printf("resultDes3Dec %v\n", Y)
    panic("test")

    resultDes3Enc := Des3Enc(tmp)
    resultDes3Dec := Des3Dec(resultDes3Enc)
    fmt.Printf("tmp %v\n", tmp)
    fmt.Printf("resultDes3Enc %v\n", resultDes3Enc)
    fmt.Printf("resultDes3Dec %v\n", resultDes3Dec)
    panic("test")

    result := Des3Dec(plaintext)
    fmt.Printf("Des3Dec %x\n", result)

    //for i := 0; i < len(plaintext)/8; i++ {
    //    startBlock := i*8
    //    c.Decrypt(out1, plaintext[startBlock:startBlock+8])
    //    //fmt.Printf("\n")
    //    //fmt.Printf("block in %d %x\n", i, plaintext[startBlock:startBlock+8])
    //    //fmt.Printf("block out %d %x\n", i, out1)
    //    //fmt.Printf("\n")
    //    out = append(out, out1...)
    //}
    //
    //fmt.Printf("Key %x\n", triplekey)
    //fmt.Printf("out %x\n", out)
    //fmt.Printf("in %x\n", plaintext)

    return result
}
