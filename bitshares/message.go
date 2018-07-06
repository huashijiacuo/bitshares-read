package bitshares

import (
	"fmt"
)

/**
采用ASE加密信息
 */
func entropyMessage(message string, key []byte) []byte {
	var messageByte []byte
	messageByte = []byte(message)
	fmt.Println("message byte:", messageByte)
	out, err := AesEncrypt(messageByte, key)
	if err != nil {
		fmt.Println("加密出错")
	}
	fmt.Println("encrypt message:", out)
	return out
}

/**
采用AES解密信息
 */
func decryptMessage(cipherMessage []byte, key []byte) string {
	plain, err :=  AesDecrypt(cipherMessage, key)
	if err != nil {
		fmt.Println("解密出错")
	}
	message := string(plain)
	//fmt.Println("plain byte:", plain)
	return message
}

