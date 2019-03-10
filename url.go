package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/otium/ytdl"
)

func validYoutubeUrl(u string) bool {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		log.Println(err)
		return false
	}
	return strings.HasPrefix(u, "https://www.youtube.com/")
}

func getDownloadURL(url string) (string, error) {
	video, err := ytdl.GetVideoInfo(url)
	if err != nil {
		fmt.Println("Failed to get video info")
		return "", err
	}
	return video.GetDownloadURL(video.Formats[0]), nil
}
