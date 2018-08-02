package templater

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// CompileTemplate compiles a final document using given template and configuration
func CompileTemplate(template io.Reader, jsonConf io.Reader, output io.Writer) error {
	// convert config file to byte slice
	conf, err := ioutil.ReadAll(jsonConf)
	if err != nil {
		return err
	}
	// convert json config byte slice into map
	m := make(map[string]entry)
	err = json.Unmarshal(conf, &m)
	if err != nil {
		return err
	}
	// fill in template with provided values
	err = useFile(template, output, m)
	if err != nil {
		return err
	}

	return nil
}
