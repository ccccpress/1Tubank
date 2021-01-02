package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/studio-b12/gowebdav"
)

func main() {
	ad, err := ioutil.ReadFile("admin")
	adm := strings.Fields(string(ad))
	if len(adm) != 3 {
		log.Println("第一行网址，第二行账户，第三行应用密码")
	}
	root := adm[0]
	user := adm[1]
	password := adm[2]
	c := gowebdav.NewClient(root, user, password)
	err = c.Mkdir("CCCC", 0644)
	if err != nil {
		log.Println(err, "配置有误")
		log.Println("第一行网址，第二行账户，第三行应用密码")
		os.Exit(2)
	}

	log.Println(" Ctrl+C 结束")
	http.HandleFunc("/", index)
	/////////////////////////////////////////////////////////////////////
	// 改端口看这里uヾ(￣▽￣)Bye~Bye~
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}
func index(w http.ResponseWriter, r *http.Request) {

	// log.Println(r.URL.Path)

	file, err := os.Open("./CCCC" + r.URL.Path)
	defer file.Close()
	if err != nil {
		ad, err := ioutil.ReadFile("admin")
		adm := strings.Fields(string(ad))
		if len(adm) != 3 {
			log.Println("第一行网址，第二行账户，第三行应用密码")
		}

		root := adm[0]
		user := adm[1]
		password := adm[2]

		c := gowebdav.NewClient(root, user, password)

		bytes, err := c.Read("CCCC" + r.URL.Path)
		if err != nil {
			log.Println(err, "文件路径错误或不存在")
		}
		fileWithDir("./CCCC", r.URL.Path, bytes)
		file, err = os.Open("./CCCC" + r.URL.Path)
		defer file.Close()
		if err != nil {
			log.Println(err, "这里不可能出错好吧")
		}
	}
	w.Header().Set("Content-Type", "image")
	io.Copy(w, file)

}
func fileWithDir(dir string, link string, content []byte) {

	// link:=/2021/01/02.png
	links := strings.Split(link, "/")
	path := strings.Join(links[:len(links)-1], "/")
	name := string(links[len(links)-1])
	///////////////////////////////////////////////////////
	os.MkdirAll(dir+path, os.ModePerm)
	file, err := os.OpenFile(dir+path+"/"+name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	file.Write(content)
}
