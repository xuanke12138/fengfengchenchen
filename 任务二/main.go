package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

type WebSite struct {
	url string
	// header map[string]string
}

func (keyword WebSite) get_html_header() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", keyword.url, nil)
	if err != nil {
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	return string(body)

}

func getdata() {
	f, err := os.Create("asd.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for i := 1; i <= 7; i++ {
		fmt.Println("正在抓取第" + strconv.Itoa(i) + "页......")
		url := "https://blog.lenconda.top/page/" + strconv.Itoa(i) + "/"
		WebSite := &WebSite{url}
		html := WebSite.get_html_header()
		//评分
		pattern3 := `datetime="(.*?)"`
		rp3 := regexp.MustCompile(pattern3)
		find_txt3 := rp3.FindAllStringSubmatch(html, -1)

		//电影名称
		pattern4 := `class="post-excerpt">(.*?)</p>`
		rp4 := regexp.MustCompile(pattern4)
		find_txt4 := rp4.FindAllStringSubmatch(html, -1)

		// 写入UTF-8 BOM
		f.WriteString("\xEF\xBB\xBF")
		//  打印全部数据和写入txt文件
		for i := 0; i < len(find_txt3); i++ {
			//fmt.Printf("%s %s\n", find_txt4[i][1], find_txt3[i][1])
			f.WriteString(find_txt3[i][1] + "\n" + "\t" + find_txt4[i][1] + "\n" + "\t" + "\r\n")
		}
	}
}

func main() {
	getdata()
}
