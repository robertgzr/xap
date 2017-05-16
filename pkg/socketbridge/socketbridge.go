package socketbridge

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type Bridgeable struct {
	net.Conn
}

const bufsize int = 1024

// RunListeningProxy starts a listener on laddr that accepts new connections
func RunListeningProxy(netw, laddr string, target net.Conn) {
	ln, err := net.Listen(netw, laddr)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		// listen for new connections
		c, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		// handle connection async
		BidirectionalBridge(c, target)
	}
}

// BidirectionalBridge mirrors traffic from A to B and from B to A
func BidirectionalBridge(connA, connB net.Conn) {
	go UnidirectionalBridge(connA, connB)
	UnidirectionalBridge(connB, connA)
}

// UnidirectionalBridge mirrors traffic from src to dst
func UnidirectionalBridge(src, dst net.Conn) {
	defer log.Printf("ended UnidirectionalBridge from %v to %v\n", src.RemoteAddr(), dst.RemoteAddr())
	for {
		dst.SetDeadline(time.Now().Add(1 * time.Second))
		buf := make([]byte, bufsize)
		nr, err := src.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Println("Err on read:", err)
			return
		}
		nw, err := dst.Write(buf[0:nr])
		if err != nil {
			log.Println("Err on write:", err)
			return
		}
		if nr != nw {
			panic(fmt.Sprintf("Lost data! Read: %v bytes, wrote: %v bytes.", nr, nw))
		}
		log.Printf("read %v, wrote %v\n", nr, nw)
	}
}
