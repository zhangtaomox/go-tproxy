package main

import (
	"flag"
	"io"
	"log"
	"net"
)

var (
	l string
	r string
)

func init() {
	flag.StringVar(&l, "l", ":8080", "local listen address")
	flag.StringVar(&r, "r", ":27017", "remote address. eg-mongo proxy: 127.0.0.1:27017(default), redis: 127.0.0.1:6379")
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", l)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	cc, _ := net.Dial("tcp", r)
	copyConn := func(src, dst net.Conn) {
		defer src.Close()
		defer dst.Close()
		_, err := io.Copy(dst, src)
		if err != nil {
			return
		}
	}
	go copyConn(c, cc)
	go copyConn(cc, c)
}
