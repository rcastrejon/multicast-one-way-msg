package multicast

import "net"

type Client struct {
	conn *net.UDPConn
	sub  chan []byte
}

func (c *Client) listenForMessages() {
	buf := make([]byte, 1024)
	for {
		n, _, _ := c.conn.ReadFrom(buf)
		c.sub <- buf[:n]
	}
}

func NewMulticastClient(addr string) (*Client, error) {
	gaddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenMulticastUDP("udp4", nil, gaddr)
	c := &Client{conn, make(chan []byte)}

	go c.listenForMessages()

	return c, err
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) Receive() string {
	data := <-c.sub
	return string(data)
}
