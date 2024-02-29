package multicast

import (
	"errors"
	"net"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
)

type Server struct {
	rooms map[string]*net.UDPConn
}

func NewMulticastServer(addrs map[string]string) (*Server, error) {
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
	return &Server{rooms}, nil
}

func (s *Server) Close() {
	for _, conn := range s.rooms {
		conn.Close()
	}
}

func (s *Server) SendTo(name string, msg string) error {
	conn, ok := s.rooms[name]
	if !ok {
		return errors.New("channel not found")
	}
	_, err := conn.Write([]byte(msg))
	return err
}

func (s *Server) BuildRoomsTable() ([]table.Column, []table.Row) {
	columns := []table.Column{
		{Title: "ID", Width: 4},
		{Title: "Name", Width: 10},
		{Title: "Address", Width: 20},
	}
	rows := make([]table.Row, len(s.rooms))
	i := 1
	for name, address := range s.rooms {
		rows[i-1] = table.Row{strconv.Itoa(i), name, address.RemoteAddr().String()}
		i++
	}
	return columns, rows
}
