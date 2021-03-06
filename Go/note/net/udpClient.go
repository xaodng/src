package main

import (
	"fmt"
	"os"
	"net"
)

func main()  {
	service := "127.0.0.1:1200"
	udpAddr, err := net.ResolveUDPAddr("udp4",service)
	checkError(err)
	conn,err := net.DialUDP("udp",nil, udpAddr)
	checkError(err)
	_,err = conn.Write([]byte("anything"))
	checkError(err)
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)
	fmt.Println(n)
	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}

func checkError(err error)  {
	if err != nil{
		fmt.Fprintf(os.Stderr, "Fatal error: ",err.Error())
		os.Exit(1)
	}
}