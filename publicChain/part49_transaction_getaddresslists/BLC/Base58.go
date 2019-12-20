package BLC

import (
	"bytes"
	"math/big"
)

//字节数组转base58（加密和解密
//ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890+/
//base58比base64少了几个字符（“0”，“O”，“I”，“l","+","/"

//需要种子数据
var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

//加密
func Base58Encode(input []byte) []byte {
	var result []byte

	x := big.NewInt(0).SetBytes(input)
	base := big.NewInt(int64(len(b58Alphabet)))
	zero := big.NewInt(0)
	mod := &big.Int{}  //取模

	for x.Cmp(zero) != 0 {
		x.DivMod(x,base,mod)
		result = append(result,b58Alphabet[mod.Int64()])
	}

	ReverseBytes(result)
	for b := range input{
		if b == 0x00 {
			result = append([]byte{b58Alphabet[0]},result...)
		}else {
			break
		}
	}
	return result
}

//解密
/*
func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	zeroBytes := 0

	for b := range input{
		if b == 0x00 {
			zeroBytes++
		}
	}

	payload := input[zeroBytes:]
	for _,b := range payload{
		charIndex := bytes.IndexByte(b58Alphabet,b)
		result.Mul(result,big.NewInt(58))
		result.Add(result,big.NewInt(int64(charIndex)))
	}

	decoded := result.Bytes()
	decoded = append(bytes.Repeat([]byte{byte(0x00)},zeroBytes))

	return decoded
}

 */

func Base58Decode(input []byte) []byte {

	result := big.NewInt(0)

	for _, b := range input {

		charIndex := bytes.IndexByte(b58Alphabet, b)

		result.Mul(result, big.NewInt(58))

		result.Add(result, big.NewInt(int64(charIndex)))

		//fmt.Println(result)

	}

	decoded := result.Bytes()

	if input[0] == b58Alphabet[0] {

		decoded = append([]byte{0x00}, decoded...)

	}

	return decoded

}
