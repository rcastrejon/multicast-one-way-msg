package multicast

import "net"

type MulticastClient struct {
	conn *net.UDPConn
	sub  chan []byte
}

func (c *MulticastClient) listenForMessages() {
	buf := make([]byte, 1024)
	for {
		n, _, _ := c.conn.ReadFrom(buf)
		c.sub <- buf[:n]
	}
}

func NewMulticastClient(addr string) (*MulticastClient, error) {
	gaddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenMulticastUDP("udp4", nil, gaddr)
	c := &MulticastClient{conn, make(chan []byte)}

	go c.listenForMessages()

	return c, err
}

func (c *MulticastClient) Close() error {
	return c.conn.Close()
}

func (c *MulticastClient) Receive() string {
	data := <-c.sub
	return string(data)
}
