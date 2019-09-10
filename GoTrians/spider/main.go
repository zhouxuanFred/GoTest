package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	
)

func gbk2utf8(str []byte) ([]byte, error) {
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewDecoder()))
}

func main() {
	clint := &http.Client{}
	var resp *http.Response
	reqest, err := http.NewRequest("POST", "https://cn.sterlingcommerce.com/mainMenu.html", strings.NewReader("j_username=zhouxun5%40lenovo.com&j_password=gPHwO7wXmuBOld6n&loginButton=Sign+In"))
	if err != nil {
		panic(err)
	}

	reqest.Header.Set("Accept", "application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	reqest.Header.Set("Accept-Encoding", "gzip, deflate, br")
	reqest.Header.Set("Cache-Control", "max-age=0")
	reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-GB;q=0.8,en;q=0.7")
	reqest.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36")
	reqest.Header.Set("Host", "cn.sterlingcommerce.com")
	reqest.Header.Set("Connection", "keep-alive")
	reqest.Header.Set("Upgrade-Insecure-Requests", "1")
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	clint.Jar = nil
	resp, err = clint.Do(reqest)

	if resp.StatusCode != http.StatusOK {
		fmt.Print("Cannot get URL connent, pls check the reason,", resp.StatusCode)
		return
	}

	//all,err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	for k, v := range reqest.Header {
		fmt.Println("header info：", k, "values:", v)
	}

	fmt.Println()

	for k, v := range resp.Header {
		fmt.Println("header info：", k, "values:", v)
	}

	var buff []byte
	buff, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	buff, _ = gbk2utf8(buff)

}
