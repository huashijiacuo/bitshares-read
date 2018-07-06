package bitshares

import (
	"io"
	"fmt"
	"encoding/binary"
	"bytes"
	"os"
	"math/rand"
	"time"
	"io/ioutil"
)


func readIndex(inputFile string) {
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
}

func bin_write(path string) {

	t := time.Now().Nanosecond()
	fp, _ := os.Create(path)
	defer fp.Close()

	rand.Seed(int64(t))

	buf := new(bytes.Buffer)
	for i := 0; i < 10; i++ {
		binary.Write(buf, binary.LittleEndian, int32(i))
		fp.Write(buf.Bytes())
	}

	// bin file contains: 0~9
}

func bin_read(path string) {

	fp, _ := os.Open(path)
	defer fp.Close()

	data := make([]byte, 4)
	var k int32
	for {
		data = data[:cap(data)]

		// read bytes to slice
		n, err := fp.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			break
		}

		// convert bytes to int32
		data = data[:n]
		binary.Read(bytes.NewBuffer(data), binary.LittleEndian, &k)
		fmt.Println(k)
	}
}