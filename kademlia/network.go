package kademlia

import (
	"fmt"
	"net"
	"os"

	"error"
)

type Network struct {
}

func Listen(port string) {
	udpAddr, err := net.ResolveUDPAddr("udp4", "0.0.0.0:"+port)
	error.CheckFatalError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	error.CheckFatalError(err)

	for {
		handleConnection(conn)
	}
}

func handleConnection(conn *net.UDPConn) {

	var buf [512]byte

	n, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Handle Connection error %s\n", err.Error())
		return
	}

	var request = string(buf[0:n])
	fmt.Fprintf(os.Stdout, "Received request : %s from %s\n", request, addr.String())

	if request == "PING" {
		conn.WriteToUDP([]byte("PONG"), addr)
	} else if requestAddr, requestError := net.ResolveUDPAddr("udp4", request); requestError == nil {
		conn.WriteToUDP([]byte("Senging PING to "+requestAddr.String()), addr)
		var contact = NewContact(NewRandomKademliaID(), requestAddr.String())
		go SendPingMessage(&contact)
	} else if request == "EXIT" {
		_, exitErr := conn.WriteToUDP([]byte("EXITING"), addr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Handle Connection EXIT error %s", exitErr.Error())
		}

		os.Exit(0)
	} else {
		conn.WriteToUDP([]byte("Invalid Request"), addr)
	}
}

func SendPingMessage(contact *Contact) {
	udpAddr, err := net.ResolveUDPAddr("udp4", contact.Address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Send Ping Message error %s", err.Error())
		return
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Send Ping Message error %s", err.Error())
		return
	}

	_, err = conn.Write([]byte("PING"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Send Ping Message error %s", err.Error())
		return
	}

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Send Ping Message error %s", err.Error())
		return
	}

	fmt.Println(string(buf[0:n]))
}

func (network *Network) SendFindContactMessage(contact *Contact) {
	// TODO
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}
