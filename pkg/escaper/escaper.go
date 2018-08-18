package escaper

var esc map[string]Escaper

func init() {
	esc = make(map[string]Escaper)
	esc["latex"] = latexEscaper{}
}

// Escaper defines the behavior of returning an escaped string of the input.
type Escaper interface {
	Escape(string) string
}

// FindEscaper uses a string mapping to find associated escapers.
func FindEscaper(key string) (Escaper, error) {
	return nil, nil
}
