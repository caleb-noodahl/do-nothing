package models

import (
	"bytes"
	"unicode"
)

type Step struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"desc"`
	Cmds        []string `yaml:"cmd"`
}

func (s Step) NameCamelCase() string {
	out := bytes.Buffer{}

	capNext := true
	for _, c := range s.Name {
		switch c {
		//ignore spaces but mark the next
		case rune(' '):
			capNext = true
		default:
			if capNext {
				out.WriteRune(unicode.ToUpper(c))
			} else {
				out.WriteRune(c)
			}
			capNext = false
		}
	}
	return out.String()
}
