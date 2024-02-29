package main

import (
	"fmt"
	"log"

	"github.com/rcastrejon/multicast-channels/pkg/multicast"
)

func main() {
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

	fmt.Println(string(cl.Receive()))
	fmt.Println(string(c2.Receive()))
}
