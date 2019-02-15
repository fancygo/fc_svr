package svr

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func ClientTest() {
	fmt.Println("client test")
	conn, err := net.Dial("tcp", "127.0.0.1:7002")
	if err != nil {
		fmt.Println("client strt err")
		return
	}

	for {
		_, err := conn.Write([]byte("hello fancy"))
		if err != nil {
			fmt.Println("write err ", err)
			return
		}
		buf := make([]byte, 512)
		count, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}
		fmt.Printf("sever call: %s, %d\n", buf, count)
		time.Sleep(time.Second)
	}
}

func TestServer(t *testing.T) {
	s := NewServer("fancy svr")

	go ClientTest()
	s.Serve()
}
