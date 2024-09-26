package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return

	}
	command_one := os.Args[1]
	if len(os.Args) > 1 {
		switch command_one {
		case "video":
			fmt.Println("video only")
		case "audIo":
			fmt.Println("audio only")
		default:
			return
		}
	}
}
