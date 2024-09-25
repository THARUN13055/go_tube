package main

import (
	"fmt"
	"log"

	"github.com/kkdai/youtube/v2"
)

func main() {
	// Define the playlist ID
	playlistID := "PLdsu0umqbb8PQC4AhdbC97DtyZ9eI1qKG"

	// Create a new YouTube client
	client := youtube.Client{}

	// Get the playlist using the playlist ID
	playlist, err := client.GetPlaylist(playlistID)
	if err != nil {
		log.Fatalf("Error fetching playlist: %v", err)
	}

	fmt.Printf("Playlist Title: %s\n", playlist.Title)
	fmt.Printf("Number of Videos in Playlist: %d\n", len(playlist.Videos))

	for index,video := range playlist.Videos {
		fmt.Printf("TITLE:%v || INDEX:%d || URL:%v \n",video.Title,index,video.ID)

		videoDetail, err := client.GetVideo(video.ID)
		if err != nil {
			log.Printf("Error fetching video details: %v", err)
			continue
		}

		format := videoDetail.Formats.WithAudioChannels()
		if len(format) == 0 {
			fmt.Printf("No audio formate founded in this video: %v",video.Title)
		}else {
			// fmt.Printf("audio formate founded in this video: %v\n",video.Title)
			if len(videoDetail.Formats) > 0 {
				fmt.Printf("The Video quality are: %v\n",videoDetail.Formats[0].QualityLabel)
			}
		}	
	}
	
}