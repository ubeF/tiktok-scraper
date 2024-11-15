package scrape

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SoundDetails struct {
	StatusCode  int
	Status_Code int
	Status_Msg  string
	MusicInfo   struct {
		Author Author
		Music  struct {
			CoverThumb string
			Duration   int
			Id         string
			Title      string
		}
		Stats struct {
			VideoCount int
		}
	}
	ShareMeta struct {
		Desc string
	}
}

func GetSoundDetails(soundId string) (SoundDetails, error) {
	url := fmt.Sprintf("https://www.tiktok.com/api/music/detail/?aid=1988&app_language=en&app_name=tiktok_web&browser_language=en-US&browser_name=Mozilla&browser_online=true&browser_platform=Linux%%20x86_64&browser_version=5.0%%20%%28X11%%29&channel=tiktok_web&cookie_enabled=true&data_collection_enabled=true&device_id=1&device_platform=web_pc&focus_state=true&from_page=music&is_fullscreen=false&is_page_visible=true&language=en&musicId=%v&os=linux&region=DE&screen_height=1080&screen_width=1920&tz_name=Europe%%2FBerlin&user_is_login=false", soundId)

	var soundDetails SoundDetails

	resp, err := http.Get(url)
	if err != nil {
		return soundDetails, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return soundDetails, err
	}

	err = json.Unmarshal(body, &soundDetails)
	if err != nil {
		return soundDetails, err
	}

	if soundDetails.StatusCode != 0 || soundDetails.Status_Code != 0 {
		return soundDetails, fmt.Errorf("API responded with: statusCode=%v, status_code=%v, msg=%v", soundDetails.StatusCode, soundDetails.Status_Code, soundDetails.Status_Msg)
	}

	return soundDetails, nil
}
