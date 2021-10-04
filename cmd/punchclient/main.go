package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

var token = flag.String("token", "test", "token to use for matching up pairs")
var serverAddr = flag.String("serverAddr", "127.0.0.1:1338", "stun server to connect to")

func main() {
	flag.Parse()
	p := make([]byte, 2048)
	addr := net.UDPAddr{
		IP: net.ParseIP("0.0.0.0"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}

	serverTypedAddr, err := net.ResolveUDPAddr("udp", *serverAddr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Sending STUN request to %v\n", serverTypedAddr)
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

	_, err = ser.WriteToUDP([]byte(fmt.Sprintf("Connected to host at %v", localAddr)), remoteAddr)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			_, _, err := ser.ReadFromUDP(p)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s\n", p)
		}
	}()

	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			text, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			text = strings.Replace(text, "\n", "", -1)
			_, err = ser.WriteToUDP([]byte(text), remoteAddr)
			if err != nil {
				panic(err)
			}
		}
	}()

	for {
	}
}
