package httpgo

import (
	"bytes"
	"reflect"
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

func (kv KeyValue) Parse(arg string, separators []string) {
	tokens := tokenizer(arg)
	for i, t := range tokens {
		if t.escaped {
			continue
		}

		var separator string
		for _, sep := range separators {
			if pos := bytes.Index(t.data, sep); pos != -1 {
				separator = sep
				break
			}
		}
		if separator != "" {
			key, value := bytes.Split(token, sep)

			keyTokens := tokenizedData(token[:i])
			keyTokens = append(keyTokens, key)
			valueTokens := tokenizedData(token[i+1:])

			kv.Key = string(bytes.Join(keyTokens, ""))
			kv.Value = string(value) + string(bytes.Join(valueTokens, ""))
			kv.Sep = separator
			kv.Orig = arg
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
		fmt.Println("parsing: ", b, " ", string(b))
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
