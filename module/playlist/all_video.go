package all_video

import (
	"fmt"
	"os"
	"strings"

	"github.com/kkdai/youtube/v2"
)

func all_video() {
	if len(os.Args) < 4 {
		fmt.Println("Please Enter the URL")
		return
	}
	URL := os.Args[3]

	// Here I need to split the url to id
	playlistIDPart := strings.Split(URL, "=")
	if len(playlistIDPart) < 2 {
		fmt.Println("Error in url please check that :)")
	}
	playlistID := playlistIDPart[1]

	// Create client

	client := youtube.Client{}

	forth_command := os.Args[4]

	// if arg 4 is --list
	switch forth_command {
	case "--list":
		Number_of_playlist,err := client.GetPlaylist(playlistID)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Number of Videos in playlist is :  %v\n",len(Number_of_playlist.Videos))
		fmt.Println("List of Videos in Playlist")
		for index , video := range Number_of_playlist.Videos {
			// Index of video and Video Title
			fmt.Printf("%d :  %v",index,video.Title)
		}
		return
	case "-d":
		
	}

	
}
