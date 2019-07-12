package udpaddr

import (
	"encoding/binary"
	"net"
)

func Byte(addr *net.UDPAddr) (ipPort [18]byte) {
	copy(ipPort[0:16], addr.IP.To16()[:])
	binary.LittleEndian.PutUint16(ipPort[16:18], uint16(addr.Port))
	return
}

func Addr(byte [18]byte) (addr *net.UDPAddr) {
	return &net.UDPAddr{IP: byte[:16], Port: int(binary.LittleEndian.Uint16(byte[16:]))}
}
