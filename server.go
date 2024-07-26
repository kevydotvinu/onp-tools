package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

// HTTP handler function
func httpHandler(w http.ResponseWriter, r *http.Request) {
	clientIP := r.RemoteAddr
	fmt.Printf("HTTP request from %s\n", clientIP)
	fmt.Fprintf(w, "Hello from HTTP! Your IP is %s", clientIP)
}

// UDP server function
func udpServer(addr string) {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Fatalf("Failed to resolve UDP address: %v", err)
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalf("Failed to create UDP connection: %v", err)
	}

	fmt.Printf("UDP server listening on %s\n", addr)

	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("Error reading from UDP connection: %v", err)
			continue
		}
		fmt.Printf("UDP message received from %s: %s\n", addr, string(buf[:n]))
		log.Printf("UDP request from %s\n", addr.String())
		_, err = conn.WriteToUDP([]byte("Message received"), addr)
		if err != nil {
			log.Printf("Error sending response: %v", err)
		}
	}
}

// TCP server function
func tcpServer(addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to start TCP server: %v", err)
	}
	defer listener.Close()

	fmt.Printf("TCP server listening on %s\n", addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		go handleTCPConnection(conn)
	}
}

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()
	clientIP := conn.RemoteAddr().String()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Printf("Error reading from TCP connection: %v", err)
		return
	}
	fmt.Printf("TCP message received from %s: %s\n", clientIP, string(buf[:n]))
	log.Printf("TCP request from %s\n", clientIP)
	_, err = conn.Write([]byte("Message received"))
	if err != nil {
		log.Printf("Error sending response: %v", err)
	}
}

// HTTP server function
func httpServer(addr string) {
	http.HandleFunc("/", httpHandler)
	fmt.Printf("HTTP server listening on %s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}

func main() {
	// Define command-line flags for protocol, stack, and port
	protocol := flag.String("protocol", "http", "Protocol to use (tcp, udp, http)")
	stack := flag.String("stack", "ipv4", "IP stack to use (ipv4, ipv6)")
	port := flag.String("port", "8080", "Port to listen on")
	flag.Parse()

	var addr string
	switch *stack {
	case "ipv4":
		addr = "0.0.0.0:" + *port
	case "ipv6":
		addr = "[::]:" + *port
	default:
		log.Fatalf("Unsupported IP stack: %s", *stack)
	}

	switch *protocol {
	case "tcp":
		tcpServer(addr)
	case "udp":
		udpServer(addr)
	case "http":
		httpServer(addr)
	default:
		log.Fatalf("Unsupported protocol: %s", *protocol)
	}
}
