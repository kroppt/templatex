package escaper

import (
	"strings"
)

func init() {
	RegisterEscaper("LaTeX", latexEscaper{})
}

type latexEscaper struct {
}

func (esc latexEscaper) Escape(in string) (out string) {
	rin := strings.NewReader(in)
	m := map[rune]string{
		'\\': "\\textbackslash{}",
		'&':  "\\&",
		'%':  "\\%",
		'$':  "\\$",
		'#':  "\\#",
		'_':  "\\_",
		'{':  "\\{",
		'}':  "\\}",
		'~':  "\\textasciitilde{}",
		'^':  "\\textasciicircum{}",
	}
	r, _, err := rin.ReadRune()
	for err == nil {
		if s, ok := m[r]; ok {
			out = out + s
		} else {
			out = out + string(r)
		}
		r, _, err = rin.ReadRune()
	}
	return out
}
