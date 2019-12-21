package BLC

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"log"
)

/*
公钥 ——》pubkeyhash: ripemd160(sha256(公钥)) :20个字节 ——》
address = base58（version:0x00 + pubkeyhash + sha256(sha256(pubkeyhash)) :截取最末尾4个字节 ）
共25字节 1KoKTUHRrGn3ZChS1zPyzn721oiZzJc84V
*/
/*
1、创建钱包{私钥、公钥}
2、先将公钥进行一次sha256,再进行一次160hash（RIPEMD160），返回20字节数组
3、base58编码(25字节)
*/
//地址由三部分组成：version，public key hash， checksum
const version = byte(0x00)
const addressChecksumLen  = 4  //截取生成的字符串（然后拼接

type Wallet struct {
	PrivateKey ecdsa.PrivateKey  //私钥通过椭圆加密算法（钥匙
	PublicKey []byte            //公钥（锁,从私钥得到
}

//生成新钱包
func NewWallet() *Wallet {
	private ,public := newKeyPair()
	wallet := Wallet{private,public}

	return &wallet
}

func (w *Wallet) GetAddress() []byte {
	pubKeyHash := HashPubKey(w.PublicKey)

	//version+pubkeyhash
	versionedPayload := append([]byte{version},pubKeyHash...)
	checksum := checksum(versionedPayload)

	//全部hash
	fullPayload := append(versionedPayload,checksum...)
	//base58（version:0x00 + pubkeyhash + sha256(sha256(pubkeyhash))
	address := Base58Encode(fullPayload)

	return address
}

//对公钥进行hash生成可读懂的内容
func HashPubKey(pubKey []byte) []byte {
	//先256hash
	publicSHA256 := sha256.Sum256(pubKey)

	//再160hash
	RIPEMD160Hasher := ripemd160.New()
	_,err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
}

//验证地址有效性
func IsValidateAddress(address []byte) bool {
	pubKeyHash := Base58Decode(address)  //地址解码
	//取后面4个字节
	actualChecksum := pubKeyHash[len(pubKeyHash)-addressChecksumLen:]
	version := pubKeyHash[0]
	//除version和后面4位外的hash值
	pubKeyHash = pubKeyHash[1: len(pubKeyHash)-addressChecksumLen]
	//两次hash后的目标值
	targetChecksum := checksum(append([]byte{version},pubKeyHash...))

	return bytes.Compare(actualChecksum,targetChecksum) == 0
}

//checksum为一个公钥生成checksum（两次256hash
func checksum(payload []byte) []byte {
	//两次进行sum256 hash
	firstSHA := sha256.Sum256(payload)  // 生成32字节
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:addressChecksumLen]  // 返回最末尾4位
}

//创建私钥、公钥（通过私钥拿到公钥
func newKeyPair() (ecdsa.PrivateKey,[]byte) {
	curve := elliptic.P256() //椭圆曲线
	private,err := ecdsa.GenerateKey(curve,rand.Reader)  //传入随机数
	if err != nil {
		log.Panic(err)
	}

	pubKey := append(private.PublicKey.X.Bytes(),private.PublicKey.Y.Bytes()...)

	return *private,pubKey
}
