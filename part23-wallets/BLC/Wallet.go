package BLC

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/ripemd160"
	"log"
)

const version = byte(0x00)
const addressChecksumLen = 4

type Wallet struct {
	//1. 私钥
	PrivateKey ecdsa.PrivateKey

	//2. 公钥
	PublicKey []byte
}

func IsValidForAddress(address []byte) bool {

	versionPublicChecksumbytes := Base58Decode(address)

	fmt.Println(versionPublicChecksumbytes)

	checkSumBytes := versionPublicChecksumbytes[len(versionPublicChecksumbytes)-addressChecksumLen:]

	versionRipemd160 := versionPublicChecksumbytes[:len(versionPublicChecksumbytes)-addressChecksumLen]

	checkBytes := CheckSum(versionRipemd160)

	if bytes.Compare(checkSumBytes, checkBytes) == 0 {
		return true
	}

	return false
}

func (w *Wallet) GetAddress() []byte {

	//1. hash160

	ripemd160Hash := w.Ripemd160Hash(w.PublicKey)

	versionRipemd160hash := append([]byte{version}, ripemd160Hash...)

	checkSumBytes := CheckSum(versionRipemd160hash)

	bytes := append(versionRipemd160hash, checkSumBytes...)

	return Base58Encode(bytes)
}

func CheckSum(payload []byte) []byte {

	hash1 := sha256.Sum256(payload)

	hash2 := sha256.Sum256(hash1[:])

	return hash2[:addressChecksumLen]
}

func (w *Wallet) Ripemd160Hash(publicKey []byte) []byte {

	//1. 256

	hash256 := sha256.New()
	hash256.Write(publicKey)
	hash := hash256.Sum(nil)

	//2. 160

	ripemd160 := ripemd160.New()
	ripemd160.Write(hash)

	return ripemd160.Sum(nil)
}

// NewWallet 创建钱包
func NewWallet() *Wallet {

	privateKey, publicKey := newKeyPair()

	return &Wallet{privateKey, publicKey}
}

// 通过私钥产生公钥
func newKeyPair() (ecdsa.PrivateKey, []byte) {

	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubKey
}
