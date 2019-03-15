package main

import (
	"fmt"
	"goscrap/scrapper"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	STR "strings"
)

func main() {

	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Please Make Sure You Have Provided A Valid URL And Have A Stable Network")
		}
	}()

	url := os.Args[1]

	images := scrapper.GetImageUrls(url)

	for _, image := range images {
		fmt.Println(image)
		downloadImage(image)
	}

}

func downloadImage(url string) {
	filename := STR.Split(url, "/")[len(STR.Split(url, "/"))-1]
	resp, err := http.Get(url)
	if err != nil {

		return
	}
	fmt.Println("fetched")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	name := ""
	if len(os.Args) > 2 {
		os.Mkdir(os.Args[2], os.ModePerm)
		name = os.Args[2]
	} else {
		os.Mkdir("images", os.ModePerm)
		name = "images"
	}

	f, _ := os.Create(filepath.Join(name, filename))
	defer f.Close()
	f.Write(body)

}
