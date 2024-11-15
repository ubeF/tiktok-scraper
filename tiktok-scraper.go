package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type SoundVideoList struct {
	Cursor      string
	HasMore     bool
	StatusCode  int
	Status_Code int
	Status_Msg  string
	ItemList    []SoundVideo
}

type SoundVideo struct {
	Id            string
	Desc          string
	DuetEnabled   bool
	StitchEnabled bool
	CreateTime    int
	Author        VideoAuthor
	AuthorStats   VideoAuthorStats
	Challenges    []Challenge
	Stats         VideoStats
	Video         VideoContent
}

type VideoAuthor struct {
	AvatarThumb string
	Nickname    string
	UniqueId    string
	SecUId      string
	Signature   string
}

type VideoContent struct {
	Cover string
}

type Challenge struct {
	Title string
}

type VideoStats struct {
	CollectCount int
	CommentCount int
	DiggCount    int
	PlayCount    int
	ShareCount   int
}

type VideoAuthorStats struct {
	DiggCount     int
	FollowerCount int
	FollowedCount int
	FriendCount   int
	Heart         int
	HeartCount    int
	VideoCount    int
}

func GetSoundVideos(soundId string, count int) ([]SoundVideo, error) {
	var videos []SoundVideo

	log.Printf("Trying to get %v videos from sound %v", count, soundId)

	for len(videos) < count {
		result, err := GetSoundVideoList(soundId, len(videos))
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

func GetSoundVideoList(soundId string, cursor int) (SoundVideoList, error) {
	url := fmt.Sprintf("https://www.tiktok.com/api/music/item_list/?aid=1988&app_language=en&app_name=tiktok_web&browser_language=en-US&browser_name=Mozilla&browser_online=true&browser_platform=Linux%%20x86_64&browser_version=5.0%%20(X11)&channel=tiktok_web&cookie_enabled=true&count=30&coverFormat=2&cursor=%v&data_collection_enabled=true&device_id=1&device_platform=web_pc&focus_state=false&from_page=music&history_len=1&is_fullscreen=false&is_page_visible=true&language=en&musicID=%v&os=linux&priority_region=&referer=&region=DE&screen_height=1080&screen_width=1920&tz_name=Europe%%2FBerlin&user_is_login=false&webcast_language=en", cursor, soundId)
	var videoList SoundVideoList

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

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	_, error := GetSoundVideos("7413835514468059936", 60)

	if error != nil {
		log.Fatal(error)
	}
}
