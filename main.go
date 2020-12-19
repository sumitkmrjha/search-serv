package main

import (
	"fmt"
	"github.com/shakesearch/server"
)

func main() {
	fmt.Println("Init Server")
	server := new(server.Server)
	fmt.Println("Init Server")
	if(server != nil) {
		server.Init()
	}else{
		fmt.Println("Error in Init Server")
	}
}
