package basic

import (
	"testing"
)

func TestCreatDir(t *testing.T) {
	err := CreatDir("/" + "123456")
	if err != nil {
		t.Error(err)
	}
}
