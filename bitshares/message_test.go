package bitshares

import (
	"testing"
	"crypto/elliptic"
	"crypto/rand"
	"bytes"
	"time"
	"crypto/sha256"
	"strconv"
	"fmt"
)

/**
测试基于椭圆曲线的秘钥交换 ECDH
以及AES加解密
 */
func TestMessage(t *testing.T) {
	for i := 100; i > 0; i-- {
		curve := elliptic.P256()
		prv1, x1, y1, err := elliptic.GenerateKey(curve, rand.Reader)
		if err != nil {
			t.Fatal(err)
		}
		if prv1 == nil {
			t.Fatal("expected prv1 to be non-nil")
		}
		if x1 == nil {
			t.Fatal("expected x1 to be non-nil")
		}
		if y1 == nil {
			t.Fatal("expected y1 to be non-nil")
		}

		prv2, x2, y2, err := elliptic.GenerateKey(curve, rand.Reader)
		if err != nil {
			t.Fatal(err)
		}
		if prv2 == nil {
			t.Fatal("expected prv2 to be non-nil")
		}
		if x2 == nil {
			t.Fatal("expected x2 to be non-nil")
		}
		if y2 == nil {
			t.Fatal("expected y2 to be non-nil")
		}

		shared1 := ComputeShared(curve, x2, y2, prv1)
		shared2 := ComputeShared(curve, x1, y1, prv2)

		if !bytes.Equal(shared1, shared2) {
			t.Fatal("expected shared1 and shared2 to be equal")
		} else {
			//fmt.Println("shared1 = ", shared1)
		}

		timeUnixNano:=time.Now().UnixNano()
		//fmt.Println(timeUnixNano)
		mySha256 := sha256.New();
		mySha256.Write(Int64ToBytes(timeUnixNano))
		hash := mySha256.Sum(nil)

		nonce_plus_shared := ComputeSharedPlusNonce(hash[0:8], shared1)
		//fmt.Println(nonce_plus_shared)

		message := "this is my first message"
		stri := strconv.Itoa(i)
		fmt.Println("stri:", stri)
		message = message + stri
		fmt.Println("message:",message)
		ciphertext := entropyMessage(message , nonce_plus_shared[0:32])
		plaintext := decryptMessage(ciphertext, nonce_plus_shared[0:32])
		fmt.Println("plaintext:", plaintext)
	}
}
