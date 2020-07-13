package main

import "net/http"

var baseURL string = "https://kr.indeed.com/jobs?q=python"

func main() {
	pages := getPages()

}

func getPages() int {
	res, err := http.Get(baseURL)
	return 0
}
asdfsdfsdghhjjjbranch