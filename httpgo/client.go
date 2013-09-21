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
