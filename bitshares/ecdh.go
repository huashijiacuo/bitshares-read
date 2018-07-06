package bitshares

import (
	"crypto/elliptic"
	"math/big"
	"bytes"
	"crypto/sha512"
)
// Package ecdh implments Elliptic curve Diffie–Hellman key sharing

// ComputeShared computes the shared key for the private key material priv and
// the x and y public coordinates
func ComputeShared(curve elliptic.Curve, x, y *big.Int, priv []byte) []byte {
	x, _ = curve.ScalarMult(x, y, priv)
	return x.Bytes()
}


/**
返回值为共享秘钥，经过hash后长度固定为512bit，nonce为随机扰动，避免再进行对称加密时的秘钥总是相同的
 */
func ComputeSharedPlusNonce(nonce []byte, shared []byte) []byte {

	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	buffer.Write(nonce)
	buffer.Write(shared)
	mySha512 := sha512.New()
	mySha512.Write(buffer.Bytes())
	nonce_plus_shared := mySha512.Sum(nil)
	//fmt.Println("nonce_plus_shared = ", nonce_plus_shared, "; len(nonce_plus_shared = ", len(nonce_plus_shared))
	return nonce_plus_shared
}