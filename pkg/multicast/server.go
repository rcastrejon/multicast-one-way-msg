package multicast

import (
	"errors"
	"net"
)

type MulticastServer struct {
	rooms map[string]*net.UDPConn
}

func NewMulticastServer(addrs map[string]string) (*MulticastServer, error) {
	rooms := make(map[string]*net.UDPConn, len(addrs))
	for name, address := range addrs {
		gaddr, err := net.ResolveUDPAddr("udp4", address)
		if err != nil {
			return nil, err
		}
		conn, err := net.DialUDP("udp4", nil, gaddr)
		if err != nil {
			return nil, err
		}
		rooms[name] = conn
	}
	return &MulticastServer{rooms}, nil
}

func (s *MulticastServer) Close() {
	for _, conn := range s.rooms {
		conn.Close()
	}
}

func (s *MulticastServer) SendTo(name string, msg string) error {
	conn, ok := s.rooms[name]
	if !ok {
		return errors.New("channel not found")
	}
	_, err := conn.Write([]byte(msg))
	return err
}
