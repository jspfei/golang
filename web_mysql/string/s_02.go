package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		url  string
		path string
	)

	fmt.Scanf("%s%s", &url, &path)
	fmt.Println(urlProcess(url))
	fmt.Println(pathProcess(path))
}

func urlProcess(url string) string {
	result := strings.HasPrefix(url, "http://")
	if !result {
		url = fmt.Sprintf("http://%s", url)
	}
	return url
}

func pathProcess(path string) string {
	result := strings.HasSuffix(path, "/")
	if !result {
		path = fmt.Sprintf("%s/", path)
	}
	return path
}
