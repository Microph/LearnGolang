package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

type Conns struct {
	mu  sync.Mutex
	Map map[string]net.Conn
}

func (c *Conns) Add(conn net.Conn) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Map[conn.RemoteAddr().String()] = conn
}

func (c *Conns) Delete(conn net.Conn) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.Map, conn.RemoteAddr().String())
}

func (c *Conns) Broadcast(text string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, conn := range c.Map {
		fmt.Fprintln(conn, text)
	}
}

var conns = &Conns{
	Map: make(map[string]net.Conn),
}

func handle(conn net.Conn) {
	defer conn.Close()
	defer conns.Delete(conn)

	conns.Add(conn)
	fmt.Println(conn.RemoteAddr())
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "time" {
			fmt.Fprintln(conn, time.Now())
		} else {
			conns.Broadcast(input)
		}

	}
}

func main() {
	l, err := net.Listen("tcp", ":4567")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}
