package udpaddrhash

import (
	"net"

	"github.com/u6du/highwayhash"

	"github.com/u6du/udpaddr"
)

func Hash(addr *net.UDPAddr) [32]byte {
	return highwayhash.Rand.Sum(udpaddr.Byte(addr))
}
