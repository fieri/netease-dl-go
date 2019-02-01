package downloader

import (
	"fmt"
	"testing"
)

var (
	url1        = "https://music.163.com/#/song?id=554242032"
	url2        = "http://music.163.com/api/song/detail/?id=554242032&ids=[554242032]"
	urlplaylist = "https://music.163.com/#/playlist?id=2621874307"
)

func TestGetHTTPResponse(t *testing.T) {
	resp, err := GetHTTPResponse(urlplaylist)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s", string(resp))
}
