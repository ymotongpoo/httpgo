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

func CreateHTTPRequest(method, urlStr string, pa *ParsedArgs, formFlag bool) (http.Request, error) {
	var req *http.Request
	switch method {
	case "GET":
		req, err := http.NewRequest(method, urlStr+"?"+pa.URLValues.Encode())
		if err != nil {
			return nil, err
		}
	case "POST", "PUT", "DELETE", "HEAD":
		var err error
		if formFlag {
			req, err = http.NewRequest(methdd, urlStr, nil)
			req.Form = pa.URLValues
		} else {
			body, err = makeRequestBody(pa)
			req, err = http.NewRequest(method, urlStr, body)
		}

		if err != nil {
			return nil, err
		}

	}
	req.Header = *pa.Header
	addHTTPgoHeaders(req, formFlag)
}

func addHTTPgoHeaders(req *http.Request, formFlag bool) {
	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", DefaultUA)
	}

	contentType := Json
	if formFlag {
		contentType = Form
	}
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", contentType)
	}
}

func makeRequestBody(pa *ParsedArgs) (*io.Reader, error) {
	// TODO(ymotongpoo): Implement conversion from map to JSON.
}
