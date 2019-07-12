package udpaddrhash

import (
	"encoding/binary"
	"net"

	"github.com/u6du/highwayhash"

	"github.com/u6du/udpaddr"
)

func Uint64(addr *net.UDPAddr) uint64 {
	return highwayhash.Sum64(udpaddr.Byte(addr))
}

func Byte8(addr *net.UDPAddr) (msg [8]byte) {
	binary.LittleEndian.PutUint64(msg[:], Uint64(addr))
	return
}

func Byte(addr *net.UDPAddr) []byte {
	b := Byte8(addr)
	return b[:]
}
