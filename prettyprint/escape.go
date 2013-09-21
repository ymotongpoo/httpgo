// This package deal with ANSI escape code, especially Control Sequence Introducer (CSI).
// For detailed explanation, please refer ANSI.
//   c.f. http://en.wikipedia.org/wiki/ANSI_escape_code
package prettyprint

import (
	"bytes"
	"io"
	"os"
	"strconv"
)

// SGR Parameter
const (
	ESC       = "\033"
	Reset     = ESC + "[0m"
	Bold      = ESC + "[1m"
	Underline = ESC + "[4m"
	Invert    = ESC + "[7m"
	FgBlack   = ESC + "[30m"
	FgRed     = ESC + "[31m"
	FgGreen   = ESC + "[32m"
	FgYellow  = ESC + "[33m"
	FgBlue    = ESC + "[34m"
	FgMagenta = ESC + "[35m"
	FgCyan    = ESC + "[36m"
	FgWhite   = ESC + "[37m"
	FgDefault = ESC + "[39m"
	BgBlack   = ESC + "[40m"
	BgRed     = ESC + "[41m"
	BgGreen   = ESC + "[42m"
	BgYellow  = ESC + "[43m"
	BgBlue    = ESC + "[44m"
	BgMagenta = ESC + "[45m"
	BgCyan    = ESC + "[46m"
	BgWhite   = ESC + "[47m"
	BgDefault = ESC + "[49m"
)

var operationMap = map[string]string{
	"up":         "A",
	"down":       "B",
	"forward":    "C",
	"back":       "D",
	"nextline":   "E",
	"prevline":   "F",
	"holisontal": "G",
	"position":   "H",
	"erasedisp":  "J",
	"eraseline":  "K",
	"scrollup":   "S",
	"scrolldown": "T",
}

func CUU(n int) []byte {
	return ESC + "[" + strconv.Itoa(n) + operationMap["up"]
}

func CUD(n int) string {
	return ESC + "[" + strconv.Itoa(n) + operationMap["down"]
}

func CUF(n int) string {
	return ESC + "[" + strconv.Itoa(n) + operationMap["forward"]
}

func CUB(n int) string {
	return ESC + "[" + strconv.Itoa(n) + operationMap["back"]
}
