package escaper

import "testing"

func Test_latexEscaper_Escape(t *testing.T) {
	esc := latexEscaper{}
	tests := []struct {
		name    string
		esc     latexEscaper
		in      string
		wantOut string
	}{
		{"dollar sign", esc, "$", "\\$"},
		{"complex with slashes", esc, "$\\_/$", "\\$\\textbackslash{}\\_/\\$"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := tt.esc.Escape(tt.in); gotOut != tt.wantOut {
				t.Errorf("latexEscaper.Escape() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
