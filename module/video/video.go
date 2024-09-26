package video

import (
	"fmt"
	"os"

	url "github.com/tharun13055/go_tube/module/url_to_id"
)

func Video(){
	// Getting the ID of the video
	Id := url.UrlSplit(os.Args[2])
	fmt.Println(Id)

	// Download the video
	


}