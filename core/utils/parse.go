package utils

import (
	"Findu/core/logger"
	"os"
	"regexp"
	"strings"
)

func ParseUrls(urls string) []string {
	var result []string
	result = strings.Split(urls, ",")
	//var re = regexp.MustCompile(`(?m)^(http(s)?:\/\/)?(www\.)?[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+(:\d+)*(\/\w+\.\w+)*[\/]{0,}$`)
	var re = regexp.MustCompile(`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`)
	for _, u := range result {
		if !re.MatchString(u) {
			logger.Warn("Please input your urls： [http|https]://www.example.com[,http://www.baidu.com]")
			logger.Errorf("Url格式错误:%v \n", u)
			os.Exit(0)
		}
	}

	return result
}
