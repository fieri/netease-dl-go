package basic

import "regexp"

var idPattern = `id=(\d+)`

// GetID get song id or playlist id.
func GetID(s string) string {
	if len(s) < 20 {
		return s
	}
	re := regexp.MustCompile(idPattern)
	return re.FindAllStringSubmatch(s, -1)[0][1]
}
