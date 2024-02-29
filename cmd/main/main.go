package main

import (
	"fmt"
	"log"

	"github.com/rcastrejon/multicast-channels/pkg/multicast"
)

func main() {
	addresses := map[string]string{
		"foo": "224.0.0.250:9999",
		"bar": "224.0.0.249:9999",
	}
	srv, err := multicast.NewMulticastServer(addresses)
	if err != nil {
		log.Fatal(err)
	}
	defer srv.Close()

	cl, err := multicast.NewMulticastClient("224.0.0.250:9999")
	if err != nil {
		log.Fatal(err)
	}
	defer cl.Close()

	c2, err := multicast.NewMulticastClient("224.0.0.249:9999")
	if err != nil {
		log.Fatal(err)
	}
	defer c2.Close()

	srv.SendTo("foo", "hello, world! from foo")
	srv.SendTo("bar", "hello, world! from bar")

	fmt.Println(string(cl.Receive()))
	fmt.Println(string(c2.Receive()))
}
