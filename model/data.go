package model

const (
	BaseURL     = "http://music.163.com/song/media/outer/url?id=%s"
	PlaylistAPI = "http://music.163.com/api/playlist/detail/?id=%s&ids=[%s]"
	SongAPI     = "http://music.163.com/api/song/detail/?id=%s&ids=[%s]"
	SearchAPI   = "http://music.163.com/api/search/get?s=%s&type=1&offset=0&total=true&limit=10"
)
