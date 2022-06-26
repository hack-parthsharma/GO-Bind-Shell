package main

import (
	"flag"

	"github.com/cyberkhalid/gosec/gobindshell/gobindshell"
)

func main() {
	ip := flag.String("i", "", "Ip address on which the server listen")
	port := flag.String("p", "", "Port number of the server")
	flag.Parse()
	var goBindShell gobindshell.GoBindShell

	goBindShell.Init(*ip, *port) //initialize the server
	goBindShell.Listen()         //listen at provided port
}
