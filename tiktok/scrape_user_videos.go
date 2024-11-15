package tiktok

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type UserVideoList struct {
	HasMore     bool
	StatusCode  int
	Status_Code int
	Status_Msg  string
	ItemList    []UserVideo
}

type UserVideo struct {
	Id         string
	Desc       string
	CreateTime int
	Author     Author
	Challenges []Challenge
	Stats      Stats
	Video      Video
}

func GetUserVideos(userId string, count int) ([]UserVideo, error) {
	var videos []UserVideo

	log.Printf("Trying to get %v videos from user %v", count, userId)

	for len(videos) < count {
		result, err := GetUserVideoList(userId, len(videos))
		if err != nil {
			return videos, err
		}
		if result.StatusCode != 0 || result.Status_Code != 0 {
			return videos, fmt.Errorf("API responded with: statusCode=%v, status_code=%v, msg=%v", result.StatusCode, result.Status_Code, result.Status_Msg)
		}

		log.Printf("Received %v videos", len(result.ItemList))

		videos = append(videos, result.ItemList...)
		if !result.HasMore {
			log.Println("Nothing left to request")
			break
		}
	}

	log.Printf("Returning %v videos", len(videos))

	return videos, nil
}

func GetUserVideoList(userId string, cursor int) (UserVideoList, error) {
	url := fmt.Sprintf("https://www.tiktok.com/api/post/item_list/?aid=1988&app_language=en&app_name=tiktok_web&browser_language=en-US&browser_name=Mozilla&browser_online=true&browser_platform=Linux%%20x86_64&browser_version=5.0%%20%%28X11%%29&channel=tiktok_web&cookie_enabled=true&count=35&coverFormat=0&cursor=%v&data_collection_enabled=true&device_id=1device_platform=web_pc&focus_state=true&from_page=user&history_len=0&is_fullscreen=false&is_page_visible=true&language=en&os=linux&priority_region=&referer=&region=DE&screen_height=1920&screen_width=1080&secUid=%v&tz_name=Europe%%2FBerlin&user_is_login=false&webcast_language=en", cursor, userId)
	var videoList UserVideoList

	resp, err := http.Get(url)
	if err != nil {
		return videoList, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return videoList, err
	}

	err = json.Unmarshal(body, &videoList)
	if err != nil {
		return videoList, err
	}

	return videoList, nil
}
