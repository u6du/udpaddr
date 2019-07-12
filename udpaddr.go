package udpaddr

import (
	"encoding/binary"
	"net"
)

func Byte(addr *net.UDPAddr) (ipPort []byte) {
	v4 := addr.IP.To4()
	port := uint16(addr.Port)
	if v4 == nil {
		ipPort = make([]byte, 18)
		copy(ipPort[0:16], addr.IP.To16()[:])
		binary.LittleEndian.PutUint16(ipPort[16:18], port)
	} else {
		ipPort = make([]byte, 6)
		copy(ipPort[0:4], v4)
		binary.LittleEndian.PutUint16(ipPort[4:6], port)
	}
	return
}

func Addr(byte []byte) (addr *net.UDPAddr) {
	pos := len(byte) - 2
	return &net.UDPAddr{IP: byte[:pos], Port: int(binary.LittleEndian.Uint16(byte[pos:]))}
}
