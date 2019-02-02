package basic

import (
	"log"
	"os"
	"os/user"
)

var (
	usr, _      = user.Current()
	UserHomeDir = usr.HomeDir
	FilePath    = UserHomeDir + "/Downloads/music"
)

func init() {
	err := CreatDir(UserHomeDir + "/Downloads/music")
	if err != nil {
		log.Fatal(err)
	}
}

// CreatDir creat path or dir if it doesn't exists.
func CreatDir(p string) error {
	if ok := IsExist(p); !ok {
		err := os.Mkdir(p, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
