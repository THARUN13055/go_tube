package playlist

// Input --all-video || --all-audio || indexing videos

import (
	"fmt"
	"os"
)

func playlist() {
	if len(os.Args) > 2 {
		return 
	}
	if len(os.Args) > 3 {
		second_command := os.Args[2]
		switch second_command {
		case "--all-video":
			fmt.Println("all videos")
		case "--all-audio":
			fmt.Println("all audio")
		default:
			return
		}
	}
}