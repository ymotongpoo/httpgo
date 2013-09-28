package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"httpgo"
)

var (
	json = flag.String("json", "",
		"Path to request template JSON file. Only avaliable with POST method.")
	xml = flag.String("xml", "",
		"Path to request temmplate XML file. Only available with POST method.")
	verbose = flag.Bool("v", false,
		"Show requested data to stdout as well")
	form = flag.Bool("f", false,
		"Explicitly specify passed data as form request data, not JSON")
	download = flag.Bool("download", false,
		"Download file in wget style")
)

func main() {
	flag.Parse()

	var file *os.File
	var err error
	if *json != "" {
		file, err = os.Open(*json)
		if err != nil {
			httpgo.LogFatal("cannot open JSON file: " + *json)
		}
	} else if *xml != "" {
		file, err = os.Open(*xml)
		if err != nil {
			httpgo.LogFatal("cannot open XML file: " + *xml)
		}
	}
	_ = file

	// TODO(ymotongpoo): call httpgo.ParseArgs()
	method, urlStr, pa, err := httpgo.ParseArgs(flag.Args())
	if err != nil {
		panic(err)
	}
	req, err := httpgo.CreateHTTPRequest(method, urlStr, pa, *form)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(req)
		panic(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
