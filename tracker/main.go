package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var scanned_input = strings.Split(scanner.Text(), " ")
		service := scanned_input[0]

		udpAddr, err := net.ResolveUDPAddr("udp4", service)
		CheckFatalError(err)

		conn, err := net.DialUDP("udp", nil, udpAddr)
		CheckFatalError(err)

		_, err = conn.Write([]byte(scanned_input[1]))
		CheckFatalError(err)

		var buf [512]byte
		n, err := conn.Read(buf[0:])
		CheckFatalError(err)

		fmt.Println(string(buf[0:n]))
	}

	if err := scanner.Err(); err != nil {
		CheckFatalError(err)
	}
}

func CheckFatalError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
