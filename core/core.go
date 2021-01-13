package core

import (
	"Findu/core/logger"
	"Findu/core/ruler"
	"Findu/core/utils"
	"flag"
	"os"
	"sync"
)

var (
	Urls string
	Filename string
	Thread int
	List bool
	Help bool
	once sync.Once
)


func InitFlags() {
	flag.StringVar(&Urls, "url", "", "Please input your urls： [http|https]://www.example.com[,http://www.baidu.com]")
	flag.StringVar(&Filename, "file", "", "Check urls in this filename")
	flag.IntVar(&Thread, "thread", 20, "Default threads 20")
	flag.BoolVar(&Help, "help", false, "Print usage")
	flag.BoolVar(&List, "show", false, "Show all rules")
	flag.Parse()
	if Help {
		flag.Usage()
		os.Exit(0)
	}

	if Urls == "" && Filename == "" && List == false{
		flag.Usage()
		os.Exit(0)
	}
	// 打印Debug 信息
	logger.Default.Level = 5
}

func init() {
	InitFlags()
	once.Do(func(){
		ruler.LoadDefault()
	})
}

func Execute() {

	if List {
		ruler.ShowInfo()
	}

	if Urls != "" {
		pUrls := utils.ParseUrls(Urls)
		logger.Success("Parsed urls")
		result := utils.Scan(pUrls, Thread)
		// todo 优化输出结果展示：Banner、Title等
		for url, rules := range result {
			logger.Infof("%s: %v", url, rules)
		}
	} else if Filename != "" {
		// todo 从文件中读取Urls 并且处理后开始扫描
	}
}