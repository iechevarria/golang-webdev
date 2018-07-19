package main

import (
    "io"
    "net"
		"log"
		"fmt"
		"bufio"
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
	for s.Scan() {
		ln := s.Text()
		fmt.Println(ln)
		if ln == "" {
			fmt.Println("End of headers")
			break
		}
	}

	body := "Hey guys, just dropping in"
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/plain\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}