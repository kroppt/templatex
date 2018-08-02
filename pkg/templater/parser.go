package templater

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
)

type entry struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// GetConfig takes in a file in the form of an io Reader and returns a JSON object that corresponds to the config parameters of the template
func GetConfig(reader io.Reader, human bool) ([]byte, error) {
	m := make(map[string]entry)
	err := useFile(reader, ioutil.Discard, m)
	if err != nil {
		return nil, err
	}
	if human {
		return json.MarshalIndent(m, "", "\t")
	}
	return json.Marshal(m)
}

// useFile extracts config from template AND writes to final document
func useFile(template io.Reader, writer io.Writer, config map[string]entry) error {
	var stack string
	b := make([]byte, 1)
	enclosed := false
	for {
		_, err := template.Read(b)
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
			name := getEntry(config, stack)
			_, err := writer.Write([]byte(config[name].Value))
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

// getEntry returns variable name and adds pair to map
func getEntry(m map[string]entry, str string) string {
	if m == nil {
		return ""
	}
	strs := strings.Split(str, ":")
	name := strs[0]
	// check if name is not in map
	if _, ok := m[name]; !ok {
		m[strs[0]] = entry{strs[1], ""}
	}
	return strs[0]
}
