package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	resp, err := http.Get("https://www.baidu.com/robots.txt")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		all, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", all)
	}
}
