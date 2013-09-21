package main

import (
	"flag"
	"os"
	_ "path/filepath"

	"httpgo"
)

var (
	json = flag.StringVar("json", "",
		"Path to request template JSON file. Only avaliable with POST method.")
	xml = flag.StringVar("xml", "",
		"Path to request temmplate XML file. Only available with POST method.")
	verbose = flag.BoolVar("v", false,
		"Show requested data to stdout as well")
	form = flag.BoolVar("f", false,
		"Explicitly specify passed data as form request data, not JSON")
	download = flag.BoolVar("download", false,
		"Download file in wget style")
)

func main() {
	flag.Parse()

	var file os.File
	if json != "" {
		file, err := os.Open(json)
		if err != nil {
			LogFatal("cannot open JSON file: " + json)
		}
	} else if xml != "" {
		file, err := os.Open(xml)
		if err != nil {
			LogFatal("cannot open XML file: " + xml)
		}
	}

	// TODO(ymotongpoo): call httpgo.ParseArgs()
}
