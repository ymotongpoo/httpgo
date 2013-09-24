package httpgo

import (
	"errors"
	"flag"
	"net/http"
	"net/url"

	"prettyprint"
)

const (
	Form      = "application/x-www-form-urlencoded; charset=utf-8"
	Json      = "application/json; charset=utf-8"
	DefaultUA = "HTTPgo/" + Version
)

func CreateHTTPRequest(method, urlStr string, pa *ParseArgs) (http.Request, error) {
	var req *http.Request
	switch method {
	case "GET":
		req, err := http.NewRequest(method, urlStr+"?"+pa.URLValues.Encode())
		if err != nil {
			return nil, err
		}
	case "POST", "PUT", "DELETE", "HEAD":
		req, err = http.NewRequest(methdd, urlStr, nil)
		if err != nil {
			return nil, err
		}
		req.Form = pa.URLValues
	}
	req.Header = *pa.Header
}

func addHTTPgoHeaders(req *http.Request) {
	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", DefaultUA)
	}
}
