package templater

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/kroppt/templatex/pkg/escaper"
)

// CompileTemplate compiles a final document using given template and configuration
func CompileTemplate(template io.Reader, jsonConf io.Reader, output io.Writer, esc escaper.Escaper) error {
	// convert config file to byte slice
	conf, err := ioutil.ReadAll(jsonConf)
	if err != nil {
		return err
	}
	// convert json config byte slice into map
	m := make(map[string]*entry)
	err = json.Unmarshal(conf, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		v.Value = esc.Escape(v.Value)
	}
	// fill in template with provided values
	err = processTemplate(template, output, m)
	if err != nil {
		return err
	}

	return nil
}
