package assessment

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"time"
)

// udp
func startUDPServer() {
	addr, err := net.ResolveUDPAddr("udp", ":8081")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error starting UDP server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("UDP server started on port 8081")

	buf := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}
		fmt.Printf("Received UDP message: %s from %v\n", string(buf[:n]), remoteAddr)
		conn.WriteToUDP(buf[:n], remoteAddr)
	}
}

// tcp
func startTCPServer() {
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP server started on port 8082")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleTCPConnection(conn)
	}
}

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}
		fmt.Printf("Received TCP message: %s\n", string(buf[:n]))
		conn.Write(buf[:n])
	}
}

// http
func echoHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	msg := r.FormValue("message")
	fmt.Fprintf(w, "Echo: %s", msg)
}

func startHTTPServer() {
	http.HandleFunc("/echo", echoHandler)
	fmt.Println("HTTP server started on port 8083")
	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		fmt.Println("Error starting HTTP server:", err)
	}
}

// sse
func sseHandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	msg := r.URL.Query().Get("message")

	for i := 0; i < 3; i++ {
		fmt.Fprintf(w, "data: Echo: %s\n\n", msg)
		flusher.Flush()
		time.Sleep(2 * time.Second)
	}
}

func startSSEServer() {
	http.HandleFunc("/sse", sseHandler)
	fmt.Println("SSE server started on port 8084")
	err := http.ListenAndServe(":8084", nil)
	if err != nil {
		fmt.Println("Error starting SSE server:", err)
	}
}

// websocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to websocket:", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading websocket message:", err)
			return
		}
		fmt.Printf("Received WebSocket message: %s\n", message)
		conn.WriteMessage(websocket.TextMessage, message)
	}
}

func startWebSocketServer() {
	http.HandleFunc("/ws", websocketHandler)
	fmt.Println("WebSocket server started on port 8085")
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		fmt.Println("Error starting WebSocket server:", err)
	}
}

func Echo() {
	go startUDPServer()
	go startTCPServer()
	go startHTTPServer()
	go startSSEServer()
	go startWebSocketServer()

	select {}
}
