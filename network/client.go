package network

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"sync"
)

func udpClient(wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := net.Dial("udp", "localhost:8081")
	if err != nil {
		fmt.Println("Error dialing UDP server:", err)
		return
	}
	defer conn.Close()

	message := "Hello UDP!"
	fmt.Printf("Sending: %s\n", message)
	conn.Write([]byte(message))

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from UDP server:", err)
		return
	}

	fmt.Printf("Received: %s\n", string(buf[:n]))
}

func tcpClient(wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := net.Dial("tcp", "localhost:8082")
	if err != nil {
		fmt.Println("Error dialing TCP server:", err)
		return
	}
	defer conn.Close()

	message := "Hello TCP!"
	fmt.Printf("Sending: %s\n", message)
	conn.Write([]byte(message))

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from TCP server:", err)
		return
	}

	fmt.Printf("Received: %s\n", string(buf[:n]))
}

func httpClient(wg *sync.WaitGroup) {
	defer wg.Done()
	endpoint := "http://localhost:8083/echo"
	data := url.Values{}
	data.Set("message", "Hello HTTP!")

	resp, err := http.PostForm(endpoint, data)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading HTTP response:", err)
		return
	}

	fmt.Printf("Received: %s\n", string(body))
}

func sseClient(wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get("http://localhost:8084/sse?message=Hello%20SSE!")
	if err != nil {
		fmt.Println("Error connecting to SSE server:", err)
		return
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading SSE response:", err)
			return
		}
		fmt.Printf("Received: %s", line)
	}
}

func websocketClient(wg *sync.WaitGroup) {
	defer wg.Done()
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8085/ws", nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	defer c.Close()

	message := "Hello WebSocket!"
	fmt.Printf("Sending: %s\n", message)
	err = c.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Fatal("Error sending WebSocket message:", err)
	}

	_, msg, err := c.ReadMessage()
	if err != nil {
		log.Fatal("Error reading WebSocket message:", err)
	}
	fmt.Printf("Received: %s\n", msg)
}

func StartClient(wg *sync.WaitGroup) {
	udpClient(wg)
	tcpClient(wg)
	httpClient(wg)
	sseClient(wg)
	websocketClient(wg)
}
