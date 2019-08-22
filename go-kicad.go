package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/antchfx/htmlquery"
)

// PathExists test if path exst
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func mkdir(Dir string) {
	exist, err := PathExists(Dir)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return
	}
	if exist {
		fmt.Printf("has dir![%v]\n", Dir)
	} else {
		fmt.Printf("no dir![%v]\n", Dir)
		// 创建文件夹
		err := os.MkdirAll(Dir, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}
}

func main() {
	doc, err := htmlquery.LoadURL("https://kicad-downloads.s3.cern.ch")
	if err != nil {
		panic(err)
	}

	for _, n := range htmlquery.Find(doc, "//key") {
		// fmt.Printf("https://kicad-downloads.s3.cern.ch/%s\n", htmlquery.OutputHTML(n, false))
		suburl := htmlquery.OutputHTML(n, false)
		fmt.Printf("suburl is : %s\n", suburl)
		// t := strings.Split(suburl, "/")
		// fmt.Printf("Split is : %s\n", t)
		// fmt.Printf("Last / is at : %d\n", strings.LastIndex(suburl, "/"))

		dirRegexp := regexp.MustCompile(`.+([a-z].[a-z]\/)*\/`)
		dir := dirRegexp.FindStringSubmatch(suburl)
		fmt.Printf("dir is :%s\n", dir)

		fmt.Println(len(dir))
		if len(dir) != 0 {
			Dir := dir[0]

			defer func() {
				// if err := recover(); err != nil {
				// 	fmt.Println(err)
				// }
				recover()
			}()

			mkdir(Dir)
		}

	}

}
