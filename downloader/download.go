package downloader

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// GetHTTPResponse get http response by url
func GetHTTPResponse(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	request.Header.Add("Origin", "https://music.163.com")
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode >= 300 && response.StatusCode <= 500 {
		return nil, errors.New("response.Status error: check http response")
	}
	return ioutil.ReadAll(response.Body)
}
