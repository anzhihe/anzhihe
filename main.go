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
	feed, _ := fp.ParseURL("https://chegva.com/feed/")
	fmt.Println(feed)
	b, err := ioutil.ReadFile("./readme.tpl")
	if err != nil {
		log.Panicf("read tpl err: %v", err)
	}

	tpl, err := template.New("").Parse(string(b))
	if err != nil {
		log.Panicf("tpl init err: %v", err)
	}

	f, err := os.OpenFile("./README.md", os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Panicf("open readme err: %v", err)
	}

	err = tpl.Execute(f, map[string]interface{}{
		"Articles": feed.Items[:5],
	})

	if err != nil {
		log.Panicf("tpl exec err: %v", err)
	}
}
