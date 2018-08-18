package escaper

import "strings"

func init() {
	RegisterEscaper("LaTeX", latexEscaper{})
}

type latexEscaper struct {
}

func (esc latexEscaper) Escape(in string) (out string) {
	out = in
	steps := []struct {
		i, o string
	}{
		{"\\", "\\textbackslash"},
		{"&", "\\&"},
		{"%", "\\%"},
		{"$", "\\$"},
		{"#", "\\#"},
		{"_", "\\_"},
		{"{", "\\{"},
		{"}", "\\}"},
		{"~", "\\textasciitilde"},
		{"^", "\\textasciicircum"},
		// replace gobbled spaces
		{"\\textbackslash ", "\\textbackslash~"},
		{"\\textasciitilde ", "\\textasciitilde~"},
		{"\\textasciicircum ", "\\textasciicircum~"},
	}
	for _, st := range steps {
		strings.Replace(out, st.i, st.o, -1)
	}
	return out
}
