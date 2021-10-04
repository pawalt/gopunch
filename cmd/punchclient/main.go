package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
)

var token = flag.String("token", "test", "token to use for matching up pairs")
var serverAddr = flag.String("serverAddr", "127.0.0.1:1338", "stun server to connect to")

func main() {
	p := make([]byte, 2048)
	addr := net.UDPAddr{
		IP: net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}

	serverTypedAddr, err := net.ResolveUDPAddr("udp", *serverAddr)
	if err != nil {
		panic(err)
	}
	_, err = ser.WriteToUDP([]byte(*token), serverTypedAddr)
	if err != nil {
		panic(err)
	}

	_, _, err = ser.ReadFromUDP(p)
	if err != nil {
		panic(err)
	}

	resp := string(p)
	resp = strings.ReplaceAll(resp, "local:", "")
	resp = strings.ReplaceAll(resp, "remote:", "")
	addresses := strings.Split(resp, "\n")

	localStrAddr := addresses[0]
	remoteStrAddr := addresses[1]

	localAddr, err := net.ResolveUDPAddr("udp", localStrAddr)
	if err != nil {
		panic(err)
	}
	remoteAddr, err := net.ResolveUDPAddr("udp", remoteStrAddr)
	if err != nil {
		panic(err)
	}

	_, err = ser.WriteToUDP([]byte(fmt.Sprintf("hello from %v", localAddr)), remoteAddr)
	if err != nil {
		panic(err)
	}

	for {
		_, _, err := ser.ReadFromUDP(p)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", p)
	}
}
