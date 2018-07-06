package bitshares

import (
	"encoding/binary"
)


func Byte20To5Uint(buf []byte) []uint32{
	var result []uint32
	result = make([]uint32, 5)
	for i := 0; i < 5; i++ {
		hash := BytesToUint32(buf[i*4 : i*4+4])
		result[i] = hash
		//fmt.Println(hash)
	}
	return result
}

func Int32ToBytes(i uint32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func BytesToUint32(buf []byte) uint32 {
	return uint32(binary.BigEndian.Uint32(buf))
}

func Byte20To5UintLittleEndian(buf []byte) []uint32 {
	var result []uint32
	result = make([]uint32, 5)
	for i := 0; i < 5; i++ {
		hash := BytesToUint32LittleEndian(buf[i*4 : i*4+4])
		result[i] = hash
		//fmt.Println(hash)
	}
	return result
}

func Int32ToBytesLittleEndian(i uint32) []byte {
	var buf = make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(i))
	return buf
}

func BytesToUint32LittleEndian(buf []byte) uint32 {
	return uint32(binary.LittleEndian.Uint32(buf))
}

func BytesToUint64LittleEndian(buf []byte) uint64 {
	return uint64(binary.LittleEndian.Uint64(buf))
}


func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToUint64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}