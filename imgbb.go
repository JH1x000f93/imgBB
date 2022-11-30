package imgBB

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ImgBBError - is a struct with upload error response
type ImgBBError struct {
	StatusCode int     `json:"status_code"`
	StatusText string  `json:"status_txt"`
	Err        errInfo `json:"error"`
}

type errInfo struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Context string `json:"context"`
}

// ImgBBResult - is a struct with upload success response
type ImgBBResult struct {
	Data       data `json:"data"`
	StatusCode int  `json:"status"`
	Success    bool `json:"success"`
}

type data struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	UrlViewer  string `json:"url_viewer"`
	Url        string `json:"url"`
	DisplayUrl string `json:"display_url"`
	Size       int    `json:"size"`
	Time       string `json:"time"`
	Expiration string `json:"expiration"`
	Image      info   `json:"image"`
	Thumb      info   `json:"thumb"`
	Medium     info   `json:"medium"`
	DeleteUrl  string `json:"delete_url"`
}

type info struct {
	Filename  string `json:"filename"`
	Name      string `json:"name"`
	Mime      string `json:"mime"`
	Extension string `json:"extension"`
	Url       string `json:"url"`
}

// imgBB.Upload("0x8d58lf04hbn84053gb26j","https://sindominio.io/profile.png")
func Upload(key, img string) (string, error) {
	var imgR ImgBBResult
	//15 days 1 300 000
	url := fmt.Sprintf("https://api.imgbb.com/1/upload?expiration=1300000&key=%s&image=%s", key, img)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&imgR)
	if imgR.Success == false {
		return "Ocurio un error", nil
	}
	return imgR.Data.Thumb.Url, nil
}
