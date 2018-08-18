package escaper

import (
	"errors"
	"strings"
)

var esc map[string]Escaper

func init() {
	esc = make(map[string]Escaper)
	RegisterEscaper("LaTeX", latexEscaper{})
}

// Escaper defines the behavior of returning an escaped string of the input.
type Escaper interface {
	Escape(string) string
}

// RegisterEscaper maps the given string to the given escaper for GetEscaper
func RegisterEscaper(key string, escaper Escaper) {
	key = strings.ToLower(key)
	esc[key] = escaper
}

// GetEscaper uses a string mapping to find associated escapers.
func GetEscaper(key string) (Escaper, error) {
	key = strings.ToLower(key)
	val, ok := esc[key]
	if !ok {
		return nil, errors.New("escaper type does not exist")
	}
	return val, nil
}
