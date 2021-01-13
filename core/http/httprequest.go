package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type Response struct {
	Body   string
	Header string
	Status int
	Title  string
}

func Get(url string) (Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}
	var result Response
	header := ""
	for key, value := range resp.Header {
		header += fmt.Sprintf("%s:%s\n", key, value)
	}
	var re = regexp.MustCompile(`<title>(.*)</title>`)
	title := "N/A"
	if x := re.FindStringSubmatch(string(body)); len(x) > 0 {
		title = x[1]
	}
	//
	//for i, match := range re.FindAllString(string(body), -1) {
	//	fmt.Println(match, "found at index", i)
	//}

	result.Header = strings.ToLower(header)
	result.Body = strings.ToLower(string(body))
	result.Status = resp.StatusCode
	result.Title = strings.ToLower(title)
	return result, nil
}
