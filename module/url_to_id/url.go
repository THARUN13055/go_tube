package url

import (
	"fmt"
	"strings"
)

func UrlSplit(url string) string {
	// Split the URL at the '=' character
	urlParts := strings.Split(url, "=")
	if len(urlParts) < 2 {
		fmt.Println("Error in URL, please check that :)")
		return "" // Return an empty string if there's an error null value
	}
	
	// Extract the ID part
	urlID := urlParts[1]
	return urlID
	
}