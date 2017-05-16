package socketbridge

import (
	"net"
	"testing"
)

func TestSocketBridge(t *testing.T) {
	sock, err := net.Dial("unix", "/tmp/xapper.sock")
	if err != nil {
		t.Fatal(err)
	}

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		var n int
		for {
			println("new client nr", n)
			c, err := net.Dial("tcp", ":8080")
			if err != nil {
				t.Fatal(err)
			}
			c.Write([]byte("{ \"command\": [\"get_property\", \"playlist-pos\"] }"))
			n += 1
		}
	}()

	var n int
	for {
		c, err := ln.Accept()
		if err != nil {
			t.Fatal(err)
		}
		println("new bridge nr", n)

		go BidirectionalBridge(sock, c)
		n += 1
	}
}
