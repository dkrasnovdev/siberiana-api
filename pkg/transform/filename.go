package transform

import "strings"

func Filename(url string) string {
	urlSplit := strings.Split(url, "/")
	filename := urlSplit[len(urlSplit)-1]
	return filename
}
