package gosock

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// SockHandler hanlder func for unix sock listener
type SockHandler func(net.Conn)

// Listen listen to some sock, register a handler
func Listen(sock string, handler SockHandler) error {
	defer log.Fatalln(syscall.Unlink(sock))
	ln, err := net.Listen("unix", sock)
	if err != nil {
		return err
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func(ln net.Listener, c chan os.Signal) {
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		log.Fatalln(ln.Close())
		os.Exit(0)
	}(ln, sigc)

	for {
		fd, err := ln.Accept()
		if err != nil {
			return err
		}

		go handler(fd)
	}
	return nil
}
