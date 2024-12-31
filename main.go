package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const TIMEOUT = 5 * time.Second

func Listen(conn *net.UDPConn) {
	buffer := make([]byte, 1024)
	for {
		bytes_read, _, _ := conn.ReadFromUDP(buffer)
		log.Printf(string(buffer[:bytes_read]))
	}
}

func Write(conn *net.UDPConn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		conn.Write([]byte(text))
		time.Sleep(TIMEOUT)
	}
}

func main() {
	remoteHost := flag.String("remoteHost", "127.0.0.1", "remoteHost")
	remotePort := flag.Int("remotePort", 2048, "remotePort")
	localHost := flag.String("localHost", "127.0.0.1", "localHost")
	localPort := flag.Int("localPort", 1024, "localPort")

	flag.Parse()
	remoteAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", *remoteHost, *remotePort))
	if err != nil {
		log.Fatal(err)
	}
	localAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", *localHost, *localPort))
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialUDP("udp", localAddr, remoteAddr)
	if err != nil {
		log.Fatal(err)
	}

	go Write(conn)
	go Listen(conn)

	select {}

}
