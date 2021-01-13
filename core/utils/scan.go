package utils

import (
	"Findu/core/http"
	"Findu/core/logger"
	"Findu/core/ruler"
	"sync"
	"time"
)

var lock sync.RWMutex

// todo 优化扫描结果展示，建立结构体


// Scan 内置指纹库 扫描函数 return: ["url": ["规则名称"]]
func Scan(urls []string, threads int) map[string][]string {
	result := make(map[string][]string)
	var wg sync.WaitGroup
	var lock sync.RWMutex
	logger.Debugf("[*] Current threads is %d", threads)
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			logger.Debugf("[*] Scan of %s", url)
			if r, err := checkRules(url, threads); err == nil {
				lock.Lock()
				result[url] = r
				lock.Unlock()
			}
		}(url)
	}
	wg.Wait()
	return result
}

// removeDuplicateElement 给切片去重复
func removeDuplicateElement(languages []string) []string {
	result := make([]string, 0, len(languages))
	temp := map[string]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// checkRules 内置指纹库 扫描工人函数 return： 返回匹配到的规则名称
func checkRules(url string, threads int) ([]string, error) {
	var foundResult []string
	done := make(chan bool, threads)
	result := make(chan string, 20)
	rule := make(chan ruler.Rule, 20)
	// pushCount 预检查数量
	var pushCount = 0
	// getCount 检查次数
	var getCount = 0
	resp,err := http.Get(url)
	if err != nil{
		logger.Errorf("Have a error: %s", err.Error())
		return nil, err
	}
	for worker := 0; worker < threads; worker++ {
		// 核心函数，用于调度指纹规则检查
		go func(rule <-chan ruler.Rule) {
			for {
				x, ok := <-rule
				if !ok {
					break
				}
				lock.Lock()
				getCount++
				lock.Unlock()
				found,name := x.Check(url, resp)
				if found {
					result <- name
				}
			}
			done <- true
		}(rule)
	}
	for _, x := range ruler.Default.Rules {
		rule <- x
		pushCount++
	}
	close(rule)
	for {
		select {
		case found := <-result:
			foundResult = append(foundResult, found)
		case <-done:
			threads--
			if threads == 0 {
				logger.Successf("The [%s] scanned ! Title: [%s] Queue count: %d, Check count: %d, Get count: %d", url, resp.Title, pushCount, getCount, len(foundResult))
				return foundResult,nil
			}
		case <-time.After(3 * time.Second):
			return foundResult, nil
		}
	}
}
