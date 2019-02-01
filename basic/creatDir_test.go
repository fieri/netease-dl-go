package basic

import (
	"fmt"
	"testing"
)

func TestCreatDir(t *testing.T) {
	s, err := CreatDir("/" + "123456")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(s)
}
