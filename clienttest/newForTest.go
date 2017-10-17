package clienttest

//import "fmt"
//fmt.Printf("start - end %v - %v \n", startBlock, endBlock)

//#include "StdAfx.h"
//#include "desn.h"
//#include <string.h>
//#include <stdio.h>
//#define mbedtls_printf printf

/* Implementation that should never be optimized out by the compiler */
//func mbedtls_zeroize( void *v, size_t n ) {
//    volatile unsigned char *p = (volatile unsigned char *)v; while( n-- ) *p++ = 0;
//}

/*
 * 32-bit integer manipulation macros (big endian)
 */
func GET_UINT32_BE(b []byte, i uint32) uint32 {
    return uint32(b[i]<<24) | uint32(b[i+1]<<16) | uint32(b[i+2]<<8) | uint32(b[i+3])
}

func PUT_UINT32_BE(n uint32, b []byte, i uint32) []byte {
    b[(i)  ] = (byte)((n) >> 24)
    b[(i)+1] = (byte)((n) >> 16)
    b[(i)+2] = (byte)((n) >> 8)
    b[(i)+3] = (byte)(n)
    return b
}

type mbedtls_des3_context struct {
    SK      uint32
    SKArray []uint32
}

var (
    SB1 []uint32 = []uint32{
        0x01010400, 0x00000000, 0x00010000, 0x01010404,
        0x01010004, 0x00010404, 0x00000004, 0x00010000,
        0x00000400, 0x01010400, 0x01010404, 0x00000400,
        0x01000404, 0x01010004, 0x01000000, 0x00000004,
        0x00000404, 0x01000400, 0x01000400, 0x00010400,
        0x00010400, 0x01010000, 0x01010000, 0x01000404,
        0x00010004, 0x01000004, 0x01000004, 0x00010004,
        0x00000000, 0x00000404, 0x00010404, 0x01000000,
        0x00010000, 0x01010404, 0x00000004, 0x01010000,
        0x01010400, 0x01000000, 0x01000000, 0x00000400,
        0x01010004, 0x00010000, 0x00010400, 0x01000004,
        0x00000400, 0x00000004, 0x01000404, 0x00010404,
        0x01010404, 0x00010004, 0x01010000, 0x01000404,
        0x01000004, 0x00000404, 0x00010404, 0x01010400,
        0x00000404, 0x01000400, 0x01000400, 0x00000000,
        0x00010004, 0x00010400, 0x00000000, 0x01010004,
    }
    SB2 []uint32 = []uint32{
        0x80108020, 0x80008000, 0x00008000, 0x00108020,
        0x00100000, 0x00000020, 0x80100020, 0x80008020,
        0x80000020, 0x80108020, 0x80108000, 0x80000000,
        0x80008000, 0x00100000, 0x00000020, 0x80100020,
        0x00108000, 0x00100020, 0x80008020, 0x00000000,
        0x80000000, 0x00008000, 0x00108020, 0x80100000,
        0x00100020, 0x80000020, 0x00000000, 0x00108000,
        0x00008020, 0x80108000, 0x80100000, 0x00008020,
        0x00000000, 0x00108020, 0x80100020, 0x00100000,
        0x80008020, 0x80100000, 0x80108000, 0x00008000,
        0x80100000, 0x80008000, 0x00000020, 0x80108020,
        0x00108020, 0x00000020, 0x00008000, 0x80000000,
        0x00008020, 0x80108000, 0x00100000, 0x80000020,
        0x00100020, 0x80008020, 0x80000020, 0x00100020,
        0x00108000, 0x00000000, 0x80008000, 0x00008020,
        0x80000000, 0x80100020, 0x80108020, 0x00108000,
    }
    SB3 []uint32 = []uint32{
        0x00000208, 0x08020200, 0x00000000, 0x08020008,
        0x08000200, 0x00000000, 0x00020208, 0x08000200,
        0x00020008, 0x08000008, 0x08000008, 0x00020000,
        0x08020208, 0x00020008, 0x08020000, 0x00000208,
        0x08000000, 0x00000008, 0x08020200, 0x00000200,
        0x00020200, 0x08020000, 0x08020008, 0x00020208,
        0x08000208, 0x00020200, 0x00020000, 0x08000208,
        0x00000008, 0x08020208, 0x00000200, 0x08000000,
        0x08020200, 0x08000000, 0x00020008, 0x00000208,
        0x00020000, 0x08020200, 0x08000200, 0x00000000,
        0x00000200, 0x00020008, 0x08020208, 0x08000200,
        0x08000008, 0x00000200, 0x00000000, 0x08020008,
        0x08000208, 0x00020000, 0x08000000, 0x08020208,
        0x00000008, 0x00020208, 0x00020200, 0x08000008,
        0x08020000, 0x08000208, 0x00000208, 0x08020000,
        0x00020208, 0x00000008, 0x08020008, 0x00020200,
    }
    SB4 []uint32 = []uint32{
        0x00802001, 0x00002081, 0x00002081, 0x00000080,
        0x00802080, 0x00800081, 0x00800001, 0x00002001,
        0x00000000, 0x00802000, 0x00802000, 0x00802081,
        0x00000081, 0x00000000, 0x00800080, 0x00800001,
        0x00000001, 0x00002000, 0x00800000, 0x00802001,
        0x00000080, 0x00800000, 0x00002001, 0x00002080,
        0x00800081, 0x00000001, 0x00002080, 0x00800080,
        0x00002000, 0x00802080, 0x00802081, 0x00000081,
        0x00800080, 0x00800001, 0x00802000, 0x00802081,
        0x00000081, 0x00000000, 0x00000000, 0x00802000,
        0x00002080, 0x00800080, 0x00800081, 0x00000001,
        0x00802001, 0x00002081, 0x00002081, 0x00000080,
        0x00802081, 0x00000081, 0x00000001, 0x00002000,
        0x00800001, 0x00002001, 0x00802080, 0x00800081,
        0x00002001, 0x00002080, 0x00800000, 0x00802001,
        0x00000080, 0x00800000, 0x00002000, 0x00802080,
    }
    SB5 []uint32 = []uint32{
        0x00000100, 0x02080100, 0x02080000, 0x42000100,
        0x00080000, 0x00000100, 0x40000000, 0x02080000,
        0x40080100, 0x00080000, 0x02000100, 0x40080100,
        0x42000100, 0x42080000, 0x00080100, 0x40000000,
        0x02000000, 0x40080000, 0x40080000, 0x00000000,
        0x40000100, 0x42080100, 0x42080100, 0x02000100,
        0x42080000, 0x40000100, 0x00000000, 0x42000000,
        0x02080100, 0x02000000, 0x42000000, 0x00080100,
        0x00080000, 0x42000100, 0x00000100, 0x02000000,
        0x40000000, 0x02080000, 0x42000100, 0x40080100,
        0x02000100, 0x40000000, 0x42080000, 0x02080100,
        0x40080100, 0x00000100, 0x02000000, 0x42080000,
        0x42080100, 0x00080100, 0x42000000, 0x42080100,
        0x02080000, 0x00000000, 0x40080000, 0x42000000,
        0x00080100, 0x02000100, 0x40000100, 0x00080000,
        0x00000000, 0x40080000, 0x02080100, 0x40000100,
    }
    SB6 []uint32 = []uint32{
        0x20000010, 0x20400000, 0x00004000, 0x20404010,
        0x20400000, 0x00000010, 0x20404010, 0x00400000,
        0x20004000, 0x00404010, 0x00400000, 0x20000010,
        0x00400010, 0x20004000, 0x20000000, 0x00004010,
        0x00000000, 0x00400010, 0x20004010, 0x00004000,
        0x00404000, 0x20004010, 0x00000010, 0x20400010,
        0x20400010, 0x00000000, 0x00404010, 0x20404000,
        0x00004010, 0x00404000, 0x20404000, 0x20000000,
        0x20004000, 0x00000010, 0x20400010, 0x00404000,
        0x20404010, 0x00400000, 0x00004010, 0x20000010,
        0x00400000, 0x20004000, 0x20000000, 0x00004010,
        0x20000010, 0x20404010, 0x00404000, 0x20400000,
        0x00404010, 0x20404000, 0x00000000, 0x20400010,
        0x00000010, 0x00004000, 0x20400000, 0x00404010,
        0x00004000, 0x00400010, 0x20004010, 0x00000000,
        0x20404000, 0x20000000, 0x00400010, 0x20004010,
    }
    SB7 []uint32 = []uint32{
        0x00200000, 0x04200002, 0x04000802, 0x00000000,
        0x00000800, 0x04000802, 0x00200802, 0x04200800,
        0x04200802, 0x00200000, 0x00000000, 0x04000002,
        0x00000002, 0x04000000, 0x04200002, 0x00000802,
        0x04000800, 0x00200802, 0x00200002, 0x04000800,
        0x04000002, 0x04200000, 0x04200800, 0x00200002,
        0x04200000, 0x00000800, 0x00000802, 0x04200802,
        0x00200800, 0x00000002, 0x04000000, 0x00200800,
        0x04000000, 0x00200800, 0x00200000, 0x04000802,
        0x04000802, 0x04200002, 0x04200002, 0x00000002,
        0x00200002, 0x04000000, 0x04000800, 0x00200000,
        0x04200800, 0x00000802, 0x00200802, 0x04200800,
        0x00000802, 0x04000002, 0x04200802, 0x04200000,
        0x00200800, 0x00000000, 0x00000002, 0x04200802,
        0x00000000, 0x00200802, 0x04200000, 0x00000800,
        0x04000002, 0x04000800, 0x00000800, 0x00200002,
    }
    SB8 []uint32 = []uint32{
        0x10001040, 0x00001000, 0x00040000, 0x10041040,
        0x10000000, 0x10001040, 0x00000040, 0x10000000,
        0x00040040, 0x10040000, 0x10041040, 0x00041000,
        0x10041000, 0x00041040, 0x00001000, 0x00000040,
        0x10040000, 0x10000040, 0x10001000, 0x00001040,
        0x00041000, 0x00040040, 0x10040040, 0x10041000,
        0x00001040, 0x00000000, 0x00000000, 0x10040040,
        0x10000040, 0x10001000, 0x00041040, 0x00040000,
        0x00041040, 0x00040000, 0x10041000, 0x00001000,
        0x00000040, 0x10040040, 0x00001000, 0x00041040,
        0x10001000, 0x00000040, 0x10000040, 0x10040000,
        0x10040040, 0x10000000, 0x00040000, 0x10001040,
        0x00000000, 0x10041040, 0x00040040, 0x10000040,
        0x10040000, 0x10001000, 0x10001040, 0x00000000,
        0x10041040, 0x00041000, 0x00041000, 0x00001040,
        0x00001040, 0x00040040, 0x10000000, 0x10041000,
    }

    /*
    * PC1: left and right halves bit-swap
    */
    LHs []uint32 = []uint32{
        0x00000000, 0x00000001, 0x00000100, 0x00000101,
        0x00010000, 0x00010001, 0x00010100, 0x00010101,
        0x01000000, 0x01000001, 0x01000100, 0x01000101,
        0x01010000, 0x01010001, 0x01010100, 0x01010101,
    }
    RHs []uint32 = []uint32{
        0x00000000, 0x01000000, 0x00010000, 0x01010000,
        0x00000100, 0x01000100, 0x00010100, 0x01010100,
        0x00000001, 0x01000001, 0x00010001, 0x01010001,
        0x00000101, 0x01000101, 0x00010101, 0x01010101,
    }
    des3_test_keys []byte = []byte{
        0xE1, 0xB7, 0xA1, 0x8B, 0x09, 0x13, 0xF7, 0x8E,
        0xC0, 0x89, 0x8E, 0xD6, 0x28, 0x2D, 0xD8, 0xD3,
        0x15, 0x00, 0x3F, 0x6B, 0xFD, 0xA4, 0x69, 0x6E,
    }
    //des3_test_keys []byte = []byte{
    //    0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0x12,
    //    0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0x11,
    //    0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0x01, 0x33,
    //}
)

/*
 * Initial Permutation macro
 */
func DES_IP(X uint32, Y uint32) (uint32, uint32) {
    T := ((X >> 4) ^ Y) & 0x0F0F0F0F
    Y ^= T
    X ^= T << 4
    T = ((X >> 16) ^ Y) & 0x0000FFFF
    Y ^= T
    X ^= T << 16
    T = ((Y >> 2) ^ X) & 0x33333333
    X ^= T
    Y ^= T << 2
    T = ((Y >> 8) ^ X) & 0x00FF00FF
    X ^= T
    Y ^= T << 8
    Y = ((Y << 1) | (Y >> 31)) & 0xFFFFFFFF
    T = (X ^ Y) & 0xAAAAAAAA
    Y ^= T
    X ^= T
    X = ((X << 1) | (X >> 31)) & 0xFFFFFFFF

    return X, Y
}

/*
 * Final Permutation macro
 */
func DES_FP(X uint32, Y uint32) (uint32, uint32) {
    X = ((X << 31) | (X >> 1)) & 0xFFFFFFFF
    T := (X ^ Y) & 0xAAAAAAAA
    X ^= T
    Y ^= T
    Y = ((Y << 31) | (Y >> 1)) & 0xFFFFFFFF
    T = ((Y >> 8) ^ X) & 0x00FF00FF
    X ^= T
    Y ^= T << 8
    T = ((Y >> 2) ^ X) & 0x33333333
    X ^= T
    Y ^= T << 2
    T = ((X >> 16) ^ Y) & 0x0000FFFF
    Y ^= T
    X ^= T << 16
    T = ((X >> 4) ^ Y) & 0x0F0F0F0F
    Y ^= T
    X ^= T << 4

    return X, Y
}

/*
 * DES round macro
 */
func DES_ROUND(X uint32, Y uint32, ctx *mbedtls_des3_context) (uint32, uint32) {
    T := ctx.SKArray[ctx.SK] ^ X
    Y ^= SB8[ (T )&0x3F ] ^
        SB6[ (T>>8)&0x3F ] ^
        SB4[ (T>>16)&0x3F ] ^
        SB2[ (T>>24)&0x3F ]
    ctx.SK++

    T = ctx.SKArray[ctx.SK] ^ ((X << 28) | (X >> 4))
    Y ^= SB7[ (T)&0x3F ] ^
        SB5[ (T>>8)&0x3F ] ^
        SB3[ (T>>16)&0x3F ] ^
        SB1[ (T>>24)&0x3F ]
    ctx.SK++

    return X, Y
}

//#if !defined(MBEDTLS_DES_SETKEY_ALT)
func mbedtls_des_setkey(SKArray []uint32, key []byte, SKIteration uint32) ([]uint32, uint32) {
    var (
        i int    = 0
        X uint32 = 0
        Y uint32 = 0
        T uint32
    )

    X = GET_UINT32_BE(key, 0)
    Y = GET_UINT32_BE(key, 4)

    /*
     * Permuted Choice 1
     */
    T = ((Y >> 4) ^ X) & 0x0F0F0F0F
    X ^= T
    Y ^= T << 4
    T = ((Y      ) ^ X) & 0x10101010
    X ^= T
    Y ^= T
    X = (LHs[ (X)&0xF] << 3) | (LHs[ (X>>8)&0xF ] << 2) |
        (LHs[ (X>>16)&0xF] << 1) | (LHs[ (X>>24)&0xF ]) |
        (LHs[ (X>>5)&0xF] << 7) | (LHs[ (X>>13)&0xF ] << 6) |
        (LHs[ (X>>21)&0xF] << 5) | (LHs[ (X>>29)&0xF ] << 4)

    Y = (RHs[ (Y>>1)&0xF] << 3) | (RHs[ (Y>>9)&0xF ] << 2) |
        (RHs[ (Y>>17)&0xF] << 1) | (RHs[ (Y>>25)&0xF ]) |
        (RHs[ (Y>>4)&0xF] << 7) | (RHs[ (Y>>12)&0xF ] << 6) |
        (RHs[ (Y>>20)&0xF] << 5) | (RHs[ (Y>>28)&0xF ] << 4)
    X &= 0x0FFFFFFF
    Y &= 0x0FFFFFFF

    /*
     * calculate subkeys
     */
    for i = 0; i < 16; i++ {
        if i < 2 || i == 8 || i == 15 {
            X = ((X << 1) | (X >> 27)) & 0x0FFFFFFF
            Y = ((Y << 1) | (Y >> 27)) & 0x0FFFFFFF
        } else {
            X = ((X << 2) | (X >> 26)) & 0x0FFFFFFF
            Y = ((Y << 2) | (Y >> 26)) & 0x0FFFFFFF
        }

        SKArray[SKIteration] = ((X << 4) & 0x24000000) | ((X << 28) & 0x10000000) |
            ((X << 14) & 0x08000000) | ((X << 18) & 0x02080000) |
            ((X << 6) & 0x01000000) | ((X << 9) & 0x00200000) |
            ((X >> 1) & 0x00100000) | ((X << 10) & 0x00040000) |
            ((X << 2) & 0x00020000) | ((X >> 10) & 0x00010000) |
            ((Y >> 13) & 0x00002000) | ((Y >> 4) & 0x00001000) |
            ((Y << 6) & 0x00000800) | ((Y >> 1) & 0x00000400) |
            ((Y >> 14) & 0x00000200) | ((Y      ) & 0x00000100) |
            ((Y >> 5) & 0x00000020) | ((Y >> 10) & 0x00000010) |
            ((Y >> 3) & 0x00000008) | ((Y >> 18) & 0x00000004) |
            ((Y >> 26) & 0x00000002) | ((Y >> 24) & 0x00000001)
        SKIteration++

        SKArray[SKIteration] = ((X << 15) & 0x20000000) | ((X << 17) & 0x10000000) |
            ((X << 10) & 0x08000000) | ((X << 22) & 0x04000000) |
            ((X >> 2) & 0x02000000) | ((X << 1) & 0x01000000) |
            ((X << 16) & 0x00200000) | ((X << 11) & 0x00100000) |
            ((X << 3) & 0x00080000) | ((X >> 6) & 0x00040000) |
            ((X << 15) & 0x00020000) | ((X >> 4) & 0x00010000) |
            ((Y >> 2) & 0x00002000) | ((Y << 8) & 0x00001000) |
            ((Y >> 14) & 0x00000808) | ((Y >> 9) & 0x00000400) |
            ((Y      ) & 0x00000200) | ((Y << 7) & 0x00000100) |
            ((Y >> 7) & 0x00000020) | ((Y >> 3) & 0x00000011) |
            ((Y << 2) & 0x00000004) | ((Y >> 21) & 0x00000002)
        SKIteration++
    }

    return SKArray, SKIteration
}

func des3_set3key(esk []uint32, dsk []uint32, key []byte) ([]uint32, []uint32) {
    var i int
    var SKIteration uint32 = 0
    esk, SKIteration = mbedtls_des_setkey(esk, key[0:8], SKIteration)
    dsk, SKIteration = mbedtls_des_setkey(dsk, key[8:16], SKIteration)
    esk, SKIteration = mbedtls_des_setkey(esk, key[16:24], SKIteration)

    for i = 0; i < 32; i += 2 {

        dsk[i] = esk[94-i]
        dsk[i+1] = esk[95-i]

        esk[i+32] = dsk[62-i]
        esk[i+33] = dsk[63-i]

        dsk[i+64] = esk[30-i]
        dsk[i+65] = esk[31-i]
    }

    return dsk, esk
}

/*
 * Triple-DES key schedule (168-bit, decryption)
 */
func mbedtls_des3_set3key_enc(ctx *mbedtls_des3_context, key []byte) int {
    var sk = make([]uint32, 96)

    des3_set3key(ctx.SKArray, sk, key)
    return 0
}

/*
 * Triple-DES key schedule (168-bit, decryption)
 */
func mbedtls_des3_set3key_dec(ctx *mbedtls_des3_context, key []byte) int {
    var sk = make([]uint32, 96)

    des3_set3key(sk, ctx.SKArray, key)
    return 0
}

func Des3Enc(Text []byte) []byte {
    var i int = 0
    var ctx3 mbedtls_des3_context
    ctx3.SKArray = make([]uint32, 96)
    textDecrypted := []byte{}
    var len int = len(Text)
    mbedtls_des3_set3key_enc(&ctx3, des3_test_keys)

    for i = 0; i < len/8; i++ {
        startBlock := i * 8
        rawDecryptedArr := mbedtls_des3_crypt_ecb(&ctx3, Text[startBlock:], Text[startBlock:])
        textDecrypted = append(textDecrypted, rawDecryptedArr...)
    }
    return textDecrypted
}

func Des3Dec(Text []byte) []byte {
    var i int = 0
    var ctx3 mbedtls_des3_context
    ctx3.SKArray = make([]uint32, 96)

    mbedtls_des3_set3key_dec(&ctx3, des3_test_keys)

    textDecrypted := []byte{}
    var len int = len(Text)
    for i = 0; i < len/8; i++ {
        startBlock := i * 8
        endBlock := startBlock + 8
        rawDecryptedArr := mbedtls_des3_crypt_ecb(&ctx3, Text[startBlock:endBlock], Text[startBlock:endBlock])
        textDecrypted = append(textDecrypted, rawDecryptedArr...)
    }
    return textDecrypted
}

func mbedtls_des3_crypt_ecb(ctx *mbedtls_des3_context, input []byte, output []byte) []byte {
    var (
        i int
        X uint32
        Y uint32
    )

    ctx.SK = 0
    outputNew := make([]byte, 8)
    X = GET_UINT32_BE(input, 0)
    Y = GET_UINT32_BE(input, 4)
    X, Y = DES_IP(X, Y)
    for i = 0; i < 8; i++ {
        Y, X = DES_ROUND(Y, X, ctx)
        X, Y = DES_ROUND(X, Y, ctx)
    }
    for i = 0; i < 8; i++ {
        X, Y = DES_ROUND(X, Y, ctx)
        Y, X = DES_ROUND(Y, X, ctx)
    }
    for i = 0; i < 8; i++ {
        Y, X = DES_ROUND(Y, X, ctx)
        X, Y = DES_ROUND(X, Y, ctx)
    }
    Y, X = DES_FP(Y, X)
    outputNew = PUT_UINT32_BE(Y, outputNew, 0)
    outputNew = PUT_UINT32_BE(X, outputNew, 4)
    ctx.SK = 0
    return outputNew
}
