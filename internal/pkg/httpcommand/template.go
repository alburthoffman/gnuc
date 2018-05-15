package httpcommand

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"fmt"
)

type Tmpl struct {
	Metadata map[string]string
	Defaults map[string]interface{}
	Headers map[string]interface{}
	Attributes map[string]interface{}
}

func Loadf(tmplfile string) (t *Tmpl, err error)  {
	content, err := ioutil.ReadFile(tmplfile)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var tmpl Tmpl
	err = json.Unmarshal(content, &tmpl)

	return &tmpl, err
}

func (t *Tmpl) PrintUsage()  {
	fmt.Printf("Usage:\n\t%s\nExample:\n\t%s\n", t.Metadata["usage"], t.Metadata["example"])
}