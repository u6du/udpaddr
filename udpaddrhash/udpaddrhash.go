package udpaddrhash

import (
	"net"

	"github.com/u6du/highwayhash"

	"github.com/u6du/udpaddr"
)

func Hash32(addr *net.UDPAddr) [32]byte {
	return highwayhash.Rand.Sum(udpaddr.Byte(addr))
}

func Hash(addr *net.UDPAddr) []byte {
	h := Hash32(addr)
	return h[:]
}
