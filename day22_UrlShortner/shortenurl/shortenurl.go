package shortenurl

import (
	"fmt"
	"math/rand"
)

func ShortenURL(url string) string {
	s := ""
	for i := 0; i < 6; i++ {
		s += string(rand.Intn(26) + 97)

	}

	shortenedURL := fmt.Sprintf("https://localhost:5000/%s",s)
	return shortenedURL
}
