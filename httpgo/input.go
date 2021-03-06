package httpgo

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	_ "reflect"
)

// Various separators used in args.
const (
	SepHeaders     = ":"
	SepCredentials = ":"
	SepProxy       = ":"
	SepData        = "="
	SepDataRawJSON = ":="
	SepFiles       = "@"
	SepQuery       = "=="
)

var (
	// Separators that become request data.
	sepGroupDataItems = []string{SepData, SepDataRawJSON, SepFiles}
	// Separators allows in ITEM arguments. Make sure to be in length descendant order.
	sepGroupItems = []string{SepDataRawJSON, SepQuery, SepHeaders, SepData, SepFiles}
)

type ParsedArgs struct {
	Header   *http.Header
	URLParam url.Values
	Data     map[string]string
	JSON     string
	File     map[string][]byte
}

func NewParsedArgs() *ParsedArgs {
	return &ParsedArgs{
		Header: &http.Header{},
	}
}

// ParseArgs parses arguments without flags and return values in following order:
//
//     * HTTP Method
//     * URL
//     * Required data
//     * error
func ParseArgs(args []string) (string, string, *ParsedArgs, error) {
	switch len(args) {
	case 0:
		return "", "", nil, errors.New("One argument is required")
	case 1:
		return "GET", args[0], NewParsedArgs(), nil
	case 2:
		return args[0], args[1], NewParsedArgs(), nil
	default:
		parsedArgs, err := ParseItems(args[2:])
		if err != nil {
			return "", "", nil, err
		}
		return args[0], args[1], parsedArgs, nil
	}
}

// ParseItems parses all ITEMs in args and returns ParsedArgs for HTTP requesst.
func ParseItems(args []string) (*ParsedArgs, error) {
	pa := &ParsedArgs{}
	for _, arg := range args {
		kv := NewKeyValue(arg)
		kv.Parse(sepGroupItems)
		// TODO(ymotongpoo): Implement function to fill kv data into ParsedArgs.

		switch kv.Sep {
		case SepHeaders:
			pa.Header.Add(kv.Key, kv.Value)
		case SepQuery:
			pa.URLParam.Add(kv.Key, kv.Value)
		case SepFiles:
			value, err := ioutil.ReadFile(kv.Value)
			if err != nil {
				return nil, err
			}
			pa.File[kv.Key] = value
		case SepData, SepDataRawJSON:
			pa.Data[kv.Key] = kv.Value
		}
	}
	return pa, nil
}

// tokenized is struct used during parsing an argument.
type tokenized struct {
	data    []byte
	escaped bool
}

type KeyValue struct {
	Key   string
	Value string
	Sep   string
	Orig  string
}

func NewKeyValue(arg string) *KeyValue {
	return &KeyValue{
		Orig: arg,
	}
}

func (kv *KeyValue) Parse(separators []string) {
	tokens := tokenize([]byte(kv.Orig))
	for i, t := range tokens {
		if t.escaped {
			continue
		}

		var separator string
		for _, sep := range separators {
			if pos := bytes.Index(t.data, []byte(sep)); pos != -1 {
				separator = sep
				break
			}
		}
		if separator != "" {
			splitted := bytes.SplitN(t.data, []byte(separator), 2)
			key := splitted[0]
			value := splitted[1]

			keyTokens := tokenizedData(tokens[:i])
			keyTokens = append(keyTokens, key)
			valueTokens := tokenizedData(tokens[i+1:])

			kv.Key = string(bytes.Join(keyTokens, []byte("")))
			kv.Value = string(value) + string(bytes.Join(valueTokens, []byte("")))
			kv.Sep = separator
			return
		}
	}
}

// tokenizedData takes all data field in each elements of tokens and
// return a slice of all those.
func tokenizedData(tokens []tokenized) [][]byte {
	allData := [][]byte{}
	for _, t := range tokens {
		allData = append(allData, t.data)
	}
	return allData
}

// Tokenize `s`. There are only two token types - strings and escaped characters:
// tokenize(r'foo\=bar\\baz')
// => [tokenized{[]bytes("foo"), false},
//     tokenized{[]bytes("="), true},
//     tokenized{[]bytes("bar"), false},
//     tokenized{[]bytes("\"), true},
//     tokenized{[]bytes("baz"), true}]
func tokenize(s []byte) []tokenized {
	tokens := []tokenized{tokenized{}}
	esc := false
	for _, b := range s {
		if esc {
			tokens = append(tokens, tokenized{[]byte{b}, true})
			tokens = append(tokens, tokenized{})
			esc = false
		} else {
			if b == '\\' {
				esc = true
			} else {
				tokens[len(tokens)-1].data = append(tokens[len(tokens)-1].data, b)
			}
		}
	}
	return tokens
}
