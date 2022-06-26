package gobindshell

import (
	"net"
)

type GoBindShell struct {
	ip   string
	port string
	addr *net.TCPAddr
}
