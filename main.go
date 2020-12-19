package main

import (
	"fmt"
	"github.com/shakesearch/server"
)

func main() {
	server := new(server.Server)
	if(server != nil) {
		server.Init()
	}else{
		fmt.Println("Error in Init Server")
	}
}
