package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
)

const (
	TypeSize   = 2
	LengthSize = 4
	maxLen     = 1 << 20 // 1 MB
)

var (
	ErrTooLarge      = errors.New("length too large")
	ErrUnexpectedEOF = errors.New("unexpected EOF")
)

type TLV struct {
	Type  uint16
	Value []byte
}

// TLV ni oqimdan o‘qish
func ReadTLV(r io.Reader) (TLV, error) {
	header := make([]byte, TypeSize+LengthSize)
	if _, err := io.ReadFull(r, header); err != nil {
		return TLV{}, ErrUnexpectedEOF
	}
	t := binary.BigEndian.Uint16(header[0:2])
	l := binary.BigEndian.Uint32(header[2:6])

	if l > uint32(maxLen) {
		return TLV{}, ErrTooLarge
	}

	v := make([]byte, l)
	if _, err := io.ReadFull(r, v); err != nil {
		return TLV{}, ErrUnexpectedEOF
	}

	return TLV{Type: t, Value: v}, nil
}

// Server ishga tushadi
func startServer() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server listening on :9000")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New client connected:", conn.RemoteAddr())

	for {
		tlv, err := ReadTLV(conn)
		if err != nil {
			fmt.Println("Client disconnected:", conn.RemoteAddr())
			return
		}
		fmt.Printf("[Server] Got message: Type=%d, Value=%s\n", tlv.Type, string(tlv.Value))
	}
}

func main() {
	startServer()
}
