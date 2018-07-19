package main

import (
    "io"
    "net"
		"log"
		"fmt"
		"bufio"
		"strings"
)

func main () {
	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln(err)
	}
	
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go serve(c)		
	}
}

func serve(c net.Conn) {
	defer c.Close()

	s := bufio.NewScanner(c)
	first := true

	var method string
	var uri string

	for s.Scan() {
		ln := s.Text()
		fmt.Println(ln)

		if first {
			method = strings.Split(ln, " ")[0]
			uri = strings.Split(ln, " ")[1]
			first = false
		}

		if ln == "" {
			fmt.Println("End of headers")
			break
		}
	}
	fmt.Println(method)
	fmt.Println(uri)

	body := "Hey guys, just dropping in   "
	body += method
	body += "   "
	body += uri

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/plain\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}