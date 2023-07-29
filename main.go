package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://chegva.com/feed/")
	if err != nil {
		panic(err)
	}
	fmt.Println(feed)
	b, err := ioutil.ReadFile("./readme.tpl")
	if err != nil {
		log.Panicf("read tpl err: %v", err)
	}
	if err != nil {

	}
	tpl, err := template.New("").Parse(string(b))
	if err != nil {
		log.Panicf("tpl init err: %v", err)
	}
	fmt.Println(tpl)
	f, err := os.OpenFile("./README.md", os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Panicf("open readme err: %v", err)
	}

	err = tpl.Execute(f, map[string]interface{}{
		"Articles": feed.Items[:1],
	})

	if err != nil {
		log.Panicf("tpl exec err: %v", err)
	}
}
