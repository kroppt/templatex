package latex

import (
	"encoding/json"
	"io"
	"strings"
)

type entry struct {
	Type  string
	Value string
}

var m map[string]entry

// GetConfig takes in a file in the form of an io Reader and returns a JSON object that corresponds to the config parameters of the template
func GetConfig(reader io.Reader) ([]byte, error) {
	var stack string
	m = make(map[string]entry)
	b := make([]byte, 1)
	enclosed := false
	for {
		_, err := reader.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if !enclosed && b[0] == '<' {
			enclosed = true
			// refresh the stack
			stack = ""
		} else if enclosed && b[0] == '>' {
			enclosed = false
			// take the stack and parse it
			getEntry(stack)
		} else if enclosed {
			stack += string(b[0])
		}
	}
	return json.Marshal(m)
}

func getEntry(str string) {
	strs := strings.Split(str, ":")
	m[strs[0]] = entry{strs[1], ""}
}
