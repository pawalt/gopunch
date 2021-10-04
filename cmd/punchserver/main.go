package main

import (
	"fmt"
	"net"
)

func main() {
	keys := make(map[string]*net.UDPAddr)

	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: 1338,
		IP:   net.ParseIP("0.0.0.0"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}
	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Read message from %v %s \n", remoteaddr, p)
		msg := string(p)
		if addr, found := keys[msg]; found {
			err = sendInfo(ser, remoteaddr, addr)
			if err != nil {
				panic(err)
			}
			delete(keys, msg)
		} else {
			keys[msg] = remoteaddr
		}
	}
}

func sendInfo(conn *net.UDPConn, first *net.UDPAddr, sec *net.UDPAddr) error {
	_, err := conn.WriteToUDP([]byte(fmt.Sprintf(`local:%v
remote:%v`, first, sec)), first)
	if err != nil {
		return err
	}

	_, err = conn.WriteToUDP([]byte(fmt.Sprintf(`local:%v
remote:%v`, sec, first)), sec)
	if err != nil {
		return err
	}

	return nil
}
