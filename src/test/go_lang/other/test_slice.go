package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	pack_Header := make([]byte, 10)
	binary.BigEndian.PutUint32(pack_Header[0:4], uint32(10) + 6)
	binary.BigEndian.PutUint16(pack_Header[4:6], uint16(1))
	binary.BigEndian.PutUint16(pack_Header[6:8], uint16(2))
	binary.BigEndian.PutUint16(pack_Header[8:10], uint16(3))

	fmt.Println(pack_Header)

}
