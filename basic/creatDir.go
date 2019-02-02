package basic

import (
	"os"
	"os/user"
)

var (
	usr, _      = user.Current()
	UserHomeDir = usr.HomeDir
	FilePath    = UserHomeDir + "/Downloads/music"
)

func init() {
	CreatDir(UserHomeDir + "/Downloads/music")
}

//CreatDir ...
func CreatDir(p string) error {
	// filePath := UserHomeDir + p
	if ok := IsExist(p); !ok {
		err := os.Mkdir(p, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
