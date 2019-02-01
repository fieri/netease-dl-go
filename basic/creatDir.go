package basic

import (
	"os"
	"os/user"
)

var (
	Usr, _      = user.Current()
	UserHomeDir = Usr.HomeDir
)

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
