package bitshares

import (
	"bytes"
	"encoding/binary"
	"os"
	"unsafe"
	"fmt"
)

type MyData struct {
	Block_pos uint64
	Block_size uint32
	Block_id [5]uint32
}

type blockData struct {
	Block_previus_hash [5]uint32
	time_point_sec uint32  // a lower resolution time_point accurate only to seconds from 1970

	/**
	witness_id
	 */
	space_id uint8
	type_id uint8
	instance uint32

	//transaction_merkle_root, an array of uint32. [5]uint32
	transaction_merkle_root [5]uint32

	//extensions, a flat_set<extension>.
	extensions uint64

	// signature_type. array<unsigned char, 65>
	witness_signature [65]uint8

	// vector<processed_transaction>

}

func struct_write(name string) {
	fp, _ := os.Create(name) //"struct.binary"
	defer fp.Close()

	// 将结构体转成bytes, 按照字段的声明顺序，但是"_"被放在最后
	var array = [5]uint32{1,2,3,4,5}
	var x uint64 = 1
	var y uint32 = 2
	data := &MyData{Block_pos:x, Block_size:y, Block_id:array}
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, data)

	// 将bytes写入文件
	fp.Write(buf.Bytes())
	fp.Sync()
}

func struct_read(name string) {

	fp, _ := os.Open(name) //"struct.binary"
	defer fp.Close()

	// 创建byte slice, 以读取bytes. 此处MyData的size为16，因为有字节对齐
	dataBytes := make([]byte, unsafe.Sizeof(MyData{}))
	data := MyData{}
	n, _ := fp.Read(dataBytes)
	dataBytes = dataBytes[:n]

	// 将bytes转成对应的struct
	binary.Read(bytes.NewBuffer(dataBytes), binary.LittleEndian, &data)
	fmt.Println(data)
}

func cat(f *os.File) {
	const NBUF = 32
	var buf [NBUF]byte
	var count int = 0
	for {
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0: // EOF
			fmt.Println("it`s over!!")
			return
		case nr > 0:
			dataBytes := make([]byte, unsafe.Sizeof(MyData{}))
			dataBytes = buf[:nr]
			data := MyData{}
			binary.Read(bytes.NewBuffer(dataBytes), binary.LittleEndian, &data)
			if count == 166748 || count == 166747 { //data.Block_size > 112 &&
				fmt.Println("count=",count)
				fmt.Println(data)

			}
			count++
			//var nw, ew = os.Stdout.Write(buf[0:nr]);
			////if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
			//if nw != nr {
			//	fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
			//}
		}
	}
}

func Char_read(name string, offset int64, len int64) {

	fp, _ := os.Open(name) //"struct.binary"
	defer fp.Close()

	const NBUF = 332
	//var buf [NBUF]byte
	var count int = 0
	//offset = 5652337
	buf := make([]byte, len)
	switch nr, err := fp.ReadAt(buf, offset); true {
	case nr < 0:
		fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
		os.Exit(1)
	case nr == 0: // EOF
		fmt.Println("count = ", count)
		return
	case nr > 0:

		//for id, ch := range buf {
		//	fmt.Println(id, ch)
		//	str := string(buf[id:id+1])
		//	fmt.Println(str)
		//}
		str2 := string(buf[:])
		fmt.Println(str2)
		if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
			fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
		}
	}
}

func ReadBuf(name string, offset int64, len int64) []byte {
	fp, _ := os.Open(name) //"struct.binary"
	defer fp.Close()

	var count int = 0

	buf := make([]byte, len)
	switch nr, err := fp.ReadAt(buf, offset); true {
	case nr < 0:
		fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
		os.Exit(1)
	case nr == 0: // EOF
		fmt.Println("count = ", count)
		return buf
	case nr > 0:
		return buf
	}
	return buf
}
