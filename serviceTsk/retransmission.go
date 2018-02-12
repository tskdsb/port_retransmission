package serviceTsk

import (
	"io"
	"log"
	"net"
)

type Line struct {
	Accept net.Conn
	Dial   net.Conn
}

type Rule struct {
	Local    string
	Remote   string
	Listener net.Listener
	Line     []Line
}

var (
	RULES []Rule
)

// Retransmission //
func Retransmission(network, local, remote string) {

	lis, err := net.Listen(network, local)
	if err != nil {
		log.Printf("Listen local(%s) error: %s\n", local, err)
		return
	}

	rule := Rule{Local: local, Remote: remote, Listener: lis, Line: *new([]Line)}

	for {
		conn1, err := lis.Accept()
		if err != nil {
			log.Printf("Accept local(%s) error: %s\n", local, err)
			continue
		}

		conn2, err := net.Dial(network, remote)
		if err != nil {
			conn1.Close()
			log.Printf("Dial remote(%s) error: %s\n", remote, err)
			continue
		}

		go io.Copy(conn2, conn1)
		go io.Copy(conn1, conn2)

		rule.Line = append(rule.Line, Line{conn1, conn2})
		RULES = append(RULES, rule)
	}
}
