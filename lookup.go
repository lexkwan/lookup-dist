package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	var queryWords string
	switch len(os.Args) {
	case 1:
		fmt.Printf("Please input the word(s) you want to query.\n E.g. lookup awesome\n")
		os.Exit(1)
	case 2:
		queryWords = os.Args[1]
	default:
		queryWords = strings.Join(os.Args[1:], "%20")
	}

	queryLinks := "http://dict.cn/" + queryWords
	resp, err := http.Get(queryLinks)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	reg := regexp.MustCompile("(?s)<div class=\"basic.*((?s)<li.*)<li style")
	results := string(reg.Find(body))
	// fmt.Println(results)

	r := strings.NewReplacer(
		"\n", "",
		"<div class=\"basic clearfix\">", "",
		"<ul class=\"dict-basic-ul\">", "",
		"<ul >", "",
		"<", "",
		"li", "",
		">", "",
		"strong", "",
		"/", "",
		"span", "",
		"style", "",
		"\t", "",
	)

	fmt.Println(r.Replace(results))
}
