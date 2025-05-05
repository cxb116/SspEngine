package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

//
//func test1(ch chan string) {
//	time.Sleep(5 * time.Second)
//	ch <- "hello"
//}
//func test2(ch chan string) {
//	time.Sleep(2 * time.Second)
//	ch <- "world"
//}
//
//func main() {
//
//	output1 := make(chan string)
//	output2 := make(chan string)
//	go test1(output1)
//	go test2(output2)
//
//	select {
//	case s1 := <-output1:
//		fmt.Println(s1)
//	case s2 := <-output2:
//		fmt.Println(s2)
//	}
//	time.Sleep(15 * time.Second)
//}

// 判断管道有没有存满
//func main() {
//	// 创建管道
//	output1 := make(chan string, 10)
//	// 子协程写数据
//	go write(output1)
//	// 取数据
//	for s := range output1 {
//		fmt.Println("res:", s)
//		time.Sleep(time.Second)
//	}
//}
//
//func write(ch chan string) {
//	for {
//		select {
//		// 写数据
//		case ch <- "hello":
//			fmt.Println("write hello")
//		default:
//			fmt.Println("channel full")
//		}
//		time.Sleep(time.Millisecond * 500)
//	}
//}

//var x int64
//
//func add(wg *sync.WaitGroup) {
//	for i := 0; i < 10; i++ {
//
//		x = x + 1
//
//	}
//	wg.Done()
//}
//
//func main() {
//	var wg sync.WaitGroup
//
//	wg.Add(6)
//	go add(&wg)
//	go add(&wg)
//	go add(&wg)
//	go add(&wg)
//	go add(&wg)
//	go add(&wg)
//
//	wg.Wait()
//	fmt.Println(x)
//}

// 读写锁
//var (
//	x      int64
//	wg     sync.WaitGroup
//	lock   sync.Mutex
//	rwlock sync.RWMutex
//)
//
//func write() {
//	rwlock.Lock()
//	fmt.Println("write 锁开始", x)
//	x = x + 1
//	//time.Sleep(time.Millisecond * 1)
//	fmt.Println("write 锁结束", x)
//	rwlock.Unlock()
//	wg.Done()
//}
//
//func read() {
//	rwlock.RLock()
//	fmt.Println("读锁开始 read")
//	time.Sleep(time.Millisecond * 1000)
//	fmt.Println("读锁结束 read")
//	rwlock.RUnlock()
//	wg.Done()
//}
//
//func read1() {
//	rwlock.RLock()
//	fmt.Println("读锁开始 read1")
//	time.Sleep(time.Millisecond * 1000)
//	fmt.Println("读锁结束 read1")
//	rwlock.RUnlock()
//	wg.Done()
//}
//
//func main() {
//
//	go func() {
//		for i := 0; i < 10; i++ {
//			wg.Add(1)
//			go write()
//		}
//	}()
//
//	for i := 0; i < 10; i++ {
//		wg.Add(2)
//		go read()
//		go read1()
//	}
//	wg.Wait()
//}

func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

// 下载图片，传入的是图片叫什么
func DownloadFile(url string, filename string) (ok bool) {
	resp, err := http.Get(url)
	HandleError(err, "http.get.url")
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "resp.body")
	filename = "D:/topgoer.com/src/github.com/student/3.0/img/" + filename
	// 写出数据
	err = ioutil.WriteFile(filename, bytes, 0666)
	if err != nil {
		return false
	} else {
		return true
	}
}

// 并发爬思路：
// 1.初始化数据管道
// 2.爬虫写出：26个协程向管道中添加图片链接
// 3.任务统计协程：检查26个任务是否都完成，完成则关闭数据管道
// 4.下载协程：从管道里读取链接并下载

var (
	// 存放图片链接的数据管道
	chanImageUrls chan string
	waitGroup     sync.WaitGroup
	// 用于监控协程
	chanTask chan string
	reImg    = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func main() {
	// myTest()
	// DownloadFile("http://i1.shaodiyejin.com/uploads/tu/201909/10242/e5794daf58_4.jpg", "1.jpg")

	// 1.初始化管道
	chanImageUrls = make(chan string, 1000000)
	chanTask = make(chan string, 26)
	// 2.爬虫协程
	for i := 1; i < 27; i++ {
		waitGroup.Add(1)
		go getImgUrls("https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html")
	}
	// 3.任务统计协程，统计26个任务是否都完成，完成则关闭管道
	waitGroup.Add(1)
	go CheckOK()
	// 4.下载协程：从管道中读取链接并下载
	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go DownloadImg()
	}
	waitGroup.Wait()
}

// 下载图片
func DownloadImg() {
	for url := range chanImageUrls {
		filename := GetFilenameFromUrl(url)
		ok := DownloadFile(url, filename)
		if ok {
			fmt.Printf("%s 下载成功\n", filename)
		} else {
			fmt.Printf("%s 下载失败\n", filename)
		}
	}
	waitGroup.Done()
}

// 截取url名字
func GetFilenameFromUrl(url string) (filename string) {
	// 返回最后一个/的位置
	lastIndex := strings.LastIndex(url, "/")
	// 切出来
	filename = url[lastIndex+1:]
	// 时间戳解决重名
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	filename = timePrefix + "_" + filename
	return
}

// 任务统计协程
func CheckOK() {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s 完成了爬取任务\n", url)
		count++
		if count == 26 {
			close(chanImageUrls)
			break
		}
	}
	waitGroup.Done()
}

// 爬图片链接到管道
// url是传的整页链接
func getImgUrls(url string) {
	urls := getImgs(url)
	// 遍历切片里所有链接，存入数据管道
	for _, url := range urls {
		chanImageUrls <- url
	}
	// 标识当前协程完成
	// 每完成一个任务，写一条数据
	// 用于监控协程知道已经完成了几个任务
	chanTask <- url
	waitGroup.Done()
}

// 获取当前页图片链接
func getImgs(url string) (urls []string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果\n", len(results))
	for _, result := range results {
		url := result[0]
		urls = append(urls, url)
	}
	return
}

// 抽取根据url获取内容
func GetPageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	// 字节转字符串
	pageStr = string(pageBytes)
	return pageStr
}
