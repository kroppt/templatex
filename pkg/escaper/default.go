package escaper

func init() {
	RegisterEscaper("", latexEscaper{})
}

type defaultEscaper struct {
}

func (esc defaultEscaper) Escape(in string) (out string) {
	return in
}
