package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

func init() {

}
func main() {
	mode := flag.String("mode", "send", "Specify 'send' for sender or 'recv' for receiver.")
	address := flag.String("address", "127.0.0.1", "Specify the IP address for sender or receiver.")
	port := flag.Int("port", 10086, "Specify the port number for sender or receiver.")

	flag.Parse()

	switch *mode {
	case "send":
		sendData(*address, *port)
	case "recv":
		receiveData(*port)
	default:
		fmt.Println("Invalid mode. Use '-mode send' for sender or '-mode recv' for receiver.")
	}
}

func sendData(address string, port int) {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		fmt.Println("ResolveUDPAddr error:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("DialUDP error:", err)
		return
	}
	defer conn.Close()

	for {
		// 获取当前时间戳
		now := time.Now().UnixNano()

		// 将时间戳转换成字节流
		timeBytes := make([]byte, 8)
		for i := 0; i < 8; i++ {
			timeBytes[i] = byte(now >> (i * 8))
		}

		// 发送数据包
		_, err = conn.Write(timeBytes)
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}

		time.Sleep(1 * time.Second)
	}
}

func receiveData(port int) {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("ResolveUDPAddr error:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("ListenUDP error:", err)
		return
	}
	defer conn.Close()

	for {
		// 接收数据包
		timeBytes := make([]byte, 8)
		_, _, err := conn.ReadFromUDP(timeBytes)
		if err != nil {
			fmt.Println("ReadFromUDP error:", err)
			return
		}

		// 将字节流转换成时间戳
		var now int64
		for i := 0; i < 8; i++ {
			now |= int64(timeBytes[i]) << (i * 8)
		}

		// 计算延迟并输出
		delay := time.Now().UnixNano() - now
		fmt.Println("单边延迟:", delay/1e6, "毫秒")
	}
}
