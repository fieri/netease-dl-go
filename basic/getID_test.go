package basic

import (
	"fmt"
	"testing"
)

type data struct {
	url  string
	want string
	cacl string
}

func TestGetID(t *testing.T) {
	testdata := []data{
		data{
			"https://music.163.com/#/song?id=721137",
			"721137",
			"",
		},
		data{
			"123456",
			"123456",
			"",
		},
		data{
			"https://music.163.com/#/song?id=721137",
			"721137",
			"",
		},
		data{
			"78777777778",
			"78777777778",
			"",
		},
	}
	for _, v := range testdata {
		v.cacl = GetID(v.url)
		fmt.Println(v.cacl)
		if v.cacl != v.want {
			t.Error("err")
		}
	}
}
