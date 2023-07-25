package main

import (
	"fmt"
	"os"
	"golang.org/x/crypto/ssh"
)

func main() {
	host := "localhost:22"
	user := "zhanliang"
	pwd := ""	
	pKey, _ := os.ReadFile(`testdata/privatekey`)


	var err error
	var signer ssh.Signer

	signer, err = ssh.ParsePrivateKey(pKey)
	if err != nil {
		fmt.Println(err.Error())
	}

	var hostkeyCallback ssh.HostKeyCallback
	hostkeyCallback = ssh.InsecureIgnoreHostKey()
	if err != nil {
		fmt.Println(err.Error())
	}

	conf := &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: hostkeyCallback,
		Auth: []ssh.AuthMethod{
			ssh.Password(pwd),
			ssh.PublicKeys(signer),
		},
	}
	var conn *ssh.Client

	conn, err = ssh.Dial("tcp", host, conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()
}