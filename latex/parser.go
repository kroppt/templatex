package latex

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
)

type entry struct {
	Type  string
	Value string
}

// GetConfig takes in a file in the form of an io Reader and returns a JSON object that corresponds to the config parameters of the template
func GetConfig(reader io.Reader) ([]byte, error) {
	m := make(map[string]entry)
	err := useFile(reader, ioutil.Discard, m)
	if err != nil {
		return nil, err
	}
	return json.Marshal(m)
}

func useFile(reader io.Reader, writer io.Writer, m map[string]entry) error {
	var stack string
	b := make([]byte, 1)
	enclosed := false
	for {
		_, err := reader.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if !enclosed && b[0] == '<' {
			enclosed = true
			// refresh the stack
			stack = ""
		} else if enclosed && b[0] == '>' {
			enclosed = false
			// take the stack and parse it
			name := getEntry(m, stack)
			_, err := writer.Write([]byte(m[name].Value))
			if err != nil {
				return err
			}
		} else if enclosed {
			stack += string(b[0])
		} else {
			_, err := writer.Write(b)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getEntry(m map[string]entry, str string) string {
	if m == nil {
		return ""
	}
	strs := strings.Split(str, ":")
	m[strs[0]] = entry{strs[1], ""}
	return strs[0]
}
