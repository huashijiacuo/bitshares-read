package bitshares

import (
	"testing"
	"os"
	"fmt"
	"time"
)

func TestReadIndex(t *testing.T) {
	var name string
	//索引文件地址
	name = "e:\\bitshares-databse\\index"

	//struct_write(name)
	//struct_read(name)
	var fp, _ = os.Open(name)
	cat(fp)
	name = "e:\\bitshares-databse\\blocks"
	Char_read(name, 18694285, 387)
}

/**
前20byt位[5]hash int32
紧接着的20-23byte为时间戳
后续扫描，可以发现转账金额1000.00000 BTS，和转账手续费21.05468
 */
func TestReadByte(t *testing.T) {
	name := "e:\\bitshares-databse\\blocks"
	var i int64
	var hashArray []uint32

	buf := ReadBuf(name, 18694285, 387)

	buf160 := buf[i:i+20]

	hashArray = Byte20To5UintLittleEndian(buf160)
	for j := 0; j < 5; j++ {
		if hashArray[j] == 1535836672 && hashArray[j+1] == 850237229 && hashArray[j+2] ==  1560314014 && hashArray[j+3] ==  2756601697 && hashArray[j+4] ==  3030724319 {
			fmt.Println("偏移：i = ", i)
			fmt.Println("*********************j = ", j, "***************")
		}
	}
	time_sec := BytesToUint32LittleEndian(buf[20 : 24])
	fmt.Println(time_sec)
	time_unix_sec := int64(time_sec)
	formatTimeStr:=time.Unix(time_unix_sec,0).Format("2006-01-02 15:04:05")

	fmt.Println(formatTimeStr)   //打印结果：2017-04-11 13:30:39
	the_time := time.Date(2018, 6, 29, 14, 53, 0, 0, time.Local)
	unix_time := the_time.Unix()
	fmt.Println(unix_time)

	// witness_id
	for j:=0; j < 350; j++ {
		space_id := buf[24+j:25+j]
		type_id := buf[25+j:26+j]
		instance := BytesToUint32(buf[26+j:30+j])
		fmt.Println("witness_id = ", uint8(space_id[0]), ".", uint8(type_id[0]), ".", instance)
	}

	//transaction_merkle_root
	buf160 = buf[30:34]
	hashArray = Byte20To5UintLittleEndian(buf160)
	for _, hash := range hashArray {
		fmt.Println(hash)
	}

	//extensions  uint64
	extensions := BytesToUint64LittleEndian(buf[34:42])
	fmt.Println("extensions = ", extensions)


	for j:=0; j < 350; j++ {
		data := buf[24+j:32+j]
		amount := BytesToUint64LittleEndian(data)
		fmt.Println("amount = ", amount)
	}

	//count= 166747
	//{18694173 112 [1535836672 850237229 1560314014 2756601697 3030724319]}
	//count= 166748
	//{18694285 387 [1552613888 1843633297 3625019220 321969644 2852454278]}
	// {18694285 387 [1552613888 1843633297 3625019220 321969644 2852454278]}
}

func TestGetIndex(t *testing.T)  {
		var name string
		//name = "struct.bin"
		name = "e:\\bitshares-databse\\index"

		//struct_write(name)
		//struct_read(name)
		var fp, _ = os.Open(name)
		cat(fp)
}
