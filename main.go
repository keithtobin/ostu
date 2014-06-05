package main

import ( 
	log "github.com/cihub/seelog"
	"bytes"
    	"fmt"
    	ssh "code.google.com/p/go.crypto/ssh"
)

var (
	sshPassword string
	sshUser string
	sshServerIP string
	sshServerPort string
	
)

func main() {
    defer log.Flush()

    sshPassword = "R0yJ3nk1n5"
    sshUser = "root"
    sshServerIP = "10.49.116.2:22"
    sshServerPort = "22"

    config := &ssh.ClientConfig{
    User: "root",
    Auth: [] ssh.AuthMethod{ 
	ssh.Password("R0yJ3nk1n5"),
    },
}

client, err := ssh.Dial("tcp", "10.49.116.2:22", config)
if err != nil {
    panic("Failed to dial: " + err.Error())
}

log.Info("Connection (SSH) to " + sshServerIP + ":" + sshServerPort)

session, err := client.NewSession()
if err != nil {
    panic("Failed to create session: " + err.Error())
}

defer session.Close()

log.Info("Connected")


var b bytes.Buffer
session.Stdout = &b
if err := session.Run("/usr/bin/whoami"); err != nil {
    panic("Failed to run: " + err.Error())
}
fmt.Println(b.String())

}


