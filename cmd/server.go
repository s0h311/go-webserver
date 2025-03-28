package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}
}

type WebserverConfig struct {
	RootDir string
}

var config WebserverConfig
  RootDir: "tmp"
}

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	reader := bufio.NewReader(conn)
	size := reader.Size()

	buffer := make([]byte, size)
	_, err := conn.Read(buffer)

	if err != nil {
		if err == io.EOF {
			return
		}

		panic(err)
	}

	requestData := GetRequestData(buffer)
	fmt.Println(requestData)
}

func GetRequestData(buffer []byte) RequestData {
	lines := strings.Split(string(buffer), "\r\n")

	firstLine := strings.Split(lines[0], " ")
	headersBodySeparatorIndex := len(lines)

	for i := range lines {
		if lines[i] == "" {
			headersBodySeparatorIndex = i
			break
		}
	}

	headerLines := lines[1:headersBodySeparatorIndex]

	var body *string

	if headersBodySeparatorIndex != len(lines) {
		body = &lines[headersBodySeparatorIndex+1]
	}

	return NewRequestData(
		firstLine[0],
		firstLine[1],
		firstLine[2],
		headerLines,
		body,
	)
}

type RequestData struct {
	Method          string
	Resource        string
	ProtocolVersion string
	Headers         []string
	Body            string
}

func NewRequestData(
	method string,
	resource string,
	protocolVersion string,
	headers []string,
	body *string,
) RequestData {
	return RequestData{
		Method:          method,
		Resource:        resource,
		ProtocolVersion: protocolVersion,
		Headers:         headers,
		Body:            *body,
	}
}
