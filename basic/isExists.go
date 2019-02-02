package basic

import "os"

// IsExist return true if dir or file exists, return false or doesn't exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
