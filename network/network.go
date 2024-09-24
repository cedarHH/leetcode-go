package network

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"sync"
	"time"
)

// udp
func startUDPServer(ctx context.Context, wg *sync.WaitGroup) {
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
	wg.Done()

	buf := make([]byte, 1024)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Shutting down UDP server...")
				return
			default:
				n, remoteAddr, err := conn.ReadFromUDP(buf)
				if err != nil {
					fmt.Println("Error reading from UDP:", err)
					continue
				}
				fmt.Printf("Received UDP message: %s from %v\n", string(buf[:n]), remoteAddr)
				_, err = conn.WriteToUDP(buf[:n], remoteAddr)
				if err != nil {
					fmt.Println("Error writing to UDP:", err)
				}
			}
		}
	}()

	<-ctx.Done()
	fmt.Println("UDP server gracefully stopped")
}

// tcp
func startTCPServer(ctx context.Context, wg *sync.WaitGroup) {
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP server started on port 8082")
	wg.Done()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stopping accepting new connections...")
				return
			default:
				conn, err := listener.Accept()
				if err != nil {
					fmt.Println("Error accepting connection:", err)
					continue
				}
				go handleTCPConnection(conn)
			}
		}
	}()

	<-ctx.Done()
	fmt.Println("TCP server gracefully stopped")
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

func startHTTPServer(ctx context.Context, wg *sync.WaitGroup) {
	server := &http.Server{
		Addr:    ":8083",
		Handler: http.DefaultServeMux,
	}

	http.HandleFunc("/echo", echoHandler)

	fmt.Println("HTTP server started on port 8083")
	wg.Done()

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Error starting HTTP server:", err)
		}
	}()

	<-ctx.Done()

	fmt.Println("Shutting down the server...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Println("Error shutting down server:", err)
	}

	fmt.Println("HTTP server gracefully stopped")
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

func startSSEServer(ctx context.Context, wg *sync.WaitGroup) {
	server := &http.Server{
		Addr:    ":8084",
		Handler: http.DefaultServeMux,
	}

	http.HandleFunc("/sse", sseHandler)

	fmt.Println("SSE server started on port 8084")
	wg.Done()

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Error starting SSE server:", err)
		}
	}()

	<-ctx.Done()

	fmt.Println("Shutting down the SSE server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Println("Error shutting down SSE server:", err)
	}

	fmt.Println("SSE server gracefully stopped")
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

func startWebSocketServer(ctx context.Context, wg *sync.WaitGroup) {
	server := &http.Server{
		Addr:    ":8085",
		Handler: http.DefaultServeMux,
	}

	http.HandleFunc("/ws", websocketHandler)

	fmt.Println("WebSocket server started on port 8085")
	wg.Done()

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Error starting WebSocket server:", err)
		}
	}()

	<-ctx.Done()

	fmt.Println("Shutting down the WebSocket server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Println("Error shutting down WebSocket server:", err)
	}

	fmt.Println("WebSocket server gracefully stopped")
}

func EchoServer(ctx context.Context, wg *sync.WaitGroup) {
	go startUDPServer(ctx, wg)
	go startTCPServer(ctx, wg)
	go startHTTPServer(ctx, wg)
	go startSSEServer(ctx, wg)
	go startWebSocketServer(ctx, wg)

	select {}
}
