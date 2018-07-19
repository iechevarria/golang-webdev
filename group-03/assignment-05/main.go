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
		c.Close()
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

	io.WriteString(c, "I see you connected.")
}