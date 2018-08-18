package escaper

import "strings"

func init() {
	RegisterEscaper("LaTeX", latexEscaper{})
}

type latexEscaper struct {
}

func (esc latexEscaper) Escape(in string) (out string) {
	out = in
	m := make(map[string]string)
	m["\\"] = "\\textbackslash"
	m["&"] = "\\&"
	m["%"] = "\\%"
	m["$"] = "\\$"
	m["#"] = "\\#"
	m["_"] = "\\_"
	m["{"] = "\\{"
	m["}"] = "\\}"
	m["~"] = "\\textasciitilde"
	m["^"] = "\\textasciicircum"
	// replace gobbled spaces
	m["\\textbackslash "] = "\\textbackslash~"
	m["\\textasciitilde "] = "\\textasciitilde~"
	m["\\textasciicircum "] = "\\textasciicircum~"
	for k, v := range m {
		strings.Replace(out, k, v, -1)
	}
	return out
}
