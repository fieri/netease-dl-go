package netease

import (
	"fmt"
	"testing"
)

var (
	id  = "571338279"
	pid = "2649779495"
)

func TestGetSongInfoByID(t *testing.T) {
	song, err := getSongInfoByID(id)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", song)
}

func TestGetSongNameByID(t *testing.T) {
	name, err := getSongNameByID(id)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(name)
}

func TestGetPlaylistInfoByID(t *testing.T) {
	playlist, err := getPlaylistInfoByID(pid)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", playlist)
}

func TestDownloadSongByID(t *testing.T) {
	err := downloadSongByID(id)
	if err != nil {
		t.Error(err)
	}
}

func TestDownloadSongByPlaylist(t *testing.T) {
	err := downloadSongByPlaylist(pid)
	if err != nil {
		t.Error(err)
	}
}
