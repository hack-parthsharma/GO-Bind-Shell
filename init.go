package gobindshell

import (
	"fmt"
	"log"
	"net"
)

func (b *GoBindShell) Init(ip, port string) {
	if ip == "" || port == "" {
		log.Fatal("[-] Provide correct argument")
	}
	var err error
	socket := fmt.Sprintf("%s:%s", ip, port)
	b.addr, err = net.ResolveTCPAddr("tcp", socket) //resolves the address
	termError("[-] Error resolving: ", err)
}
