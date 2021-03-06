package netease

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/tidwall/gjson"
	"github.com/trytwice/netease-dl-go/basic"
	"github.com/trytwice/netease-dl-go/downloader"
	"github.com/trytwice/netease-dl-go/model"
)

func getSongInfoByID(id string) (model.SongInfo, error) {
	song := model.SongInfo{}
	url := fmt.Sprintf(model.SongAPI, id, id)
	resp, err := downloader.GetHTTPResponse(url)
	if err != nil {
		return song, err
	}
	info := gjson.ParseBytes(resp).Get("songs")
	for _, s := range info.Array() {
		song.ID = s.Get("id").String()
		song.SongName = strings.TrimSpace(s.Get("name").String())
		for _, v := range s.Get("artists").Array() {
			song.ArtistsName = append(song.ArtistsName, v.Get("name").String())
		}
	}
	return song, nil
}

func getSongNameByID(id string) (string, error) {
	songInfo, err := getSongInfoByID(id)
	if err != nil {
		return "", err
	}
	name := songInfo.SongName
	for _, v := range songInfo.ArtistsName {
		name += "-" + v
	}
	return name, nil
}

func getSongResp(id string) ([]byte, error) {
	url := fmt.Sprintf(model.BaseURL, id)
	resp, err := downloader.GetHTTPResponse(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DownloadSongByID download a song by id.
func DownloadSongByID(id string) error {
	resp, err := getSongResp(id)
	if err != nil {
		return err
	}
	name, err := getSongNameByID(id)
	if err != nil {
		return err
	}
	err = saveSong(name, resp)
	return nil
}

func saveSong(name string, song []byte) error {
	fileName := basic.FilePath + "/" + name + ".mp3"
	if ok := basic.IsExist(fileName); ok {
		return nil
	}
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(song)
	if err != nil {
		return err
	}
	return nil
}

func getPlaylistInfoByID(id string) (model.PlaylistInfo, error) {
	playlist := model.PlaylistInfo{}
	url := fmt.Sprintf(model.PlaylistAPI, id, id)
	resp, err := downloader.GetHTTPResponse(url)
	if err != nil {
		return playlist, err
	}
	info := gjson.ParseBytes(resp).Get("result")
	for _, v := range info.Get("tracks").Array() {
		playlist.SongID = append(playlist.SongID, v.Get("id").String())
	}
	playlist.ID = info.Get("id").String()
	playlist.SongCount = info.Get("trackCount").String()
	return playlist, nil
}

// DownloadSongByPlaylist download songs by playlist's id.
func DownloadSongByPlaylist(id string) error {
	err := basic.CreatDir(basic.FilePath + "/" + id)
	if err != nil {
		return err
	}
	pl, err := getPlaylistInfoByID(id)
	if err != nil {
		return err
	}
	for _, v := range pl.SongID {
		resp, _ := getSongResp(v)
		name, _ := getSongNameByID(v)
		_ = saveSong(id+"/"+name, resp)
	}
	return nil
}

var waitgroup sync.WaitGroup

// DownloadSongByPlaylistCur download songs by playlist's id use go currency.
func DownloadSongByPlaylistCur(id string) error {
	err := basic.CreatDir(basic.FilePath + "/" + id)
	if err != nil {
		return err
	}
	pl, err := getPlaylistInfoByID(id)
	if err != nil {
		return err
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	for _, v := range pl.SongID {
		waitgroup.Add(1)
		go gao(v, id)
	}
	waitgroup.Wait()
	return nil
}

func gao(songID, listID string) {
	resp, _ := getSongResp(songID)
	name, _ := getSongNameByID(songID)
	_ = saveSong(listID+"/"+name, resp)
	waitgroup.Done()
}

// SearchSong search song info by given name, return 10 results.
func SearchSong(name string) ([]model.SongInfo, error) {
	result := []model.SongInfo{}
	url := fmt.Sprintf(model.SearchAPI, name)
	resp, err := downloader.GetHTTPResponse(url)
	if err != nil {
		return result, err
	}
	info := gjson.ParseBytes(resp).Get("result")
	for _, v := range info.Get("songs").Array() {
		song := model.SongInfo{}
		song.ID = v.Get("id").String()
		song.SongName = strings.TrimSpace(v.Get("name").String())
		for _, k := range v.Get("artists").Array() {
			song.ArtistsName = append(song.ArtistsName, k.Get("name").String())
		}
		result = append(result, song)
	}
	return result, nil
}
