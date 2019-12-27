package main

import (
	"context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
	"log"
)

func GetYoutubeService(ctx context.Context,  b []byte) *youtube.Service {
	config, err := google.ConfigFromJSON(b, youtube.YoutubeScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := GetClient(ctx, config)
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}
	log.Printf("Authorized on account Youtube account ")

	return service
}

func playlistItemInsert(service *youtube.Service, videoId string, playlistId string) string {
	resourceId := youtube.ResourceId{
		Kind:    "youtube#video",
		VideoId: videoId,
	}
	log.Printf("VideoId:"+videoId)

	playlistItemSnippet := youtube.PlaylistItemSnippet{
		PlaylistId: playlistId,
		ResourceId: &resourceId,
	}

	playlistItem :=  youtube.PlaylistItem{
		Snippet: &playlistItemSnippet,
	}

	call := service.PlaylistItems.Insert("snippet", &playlistItem)
	_, err := call.Do()

	if err != nil {
		return err.Error()
	}
	return "As you wish"
}