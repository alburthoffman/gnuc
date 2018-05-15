package http

import (
	"testing"
	"log"
)

func TestExecuteCommand(t *testing.T)  {
	cmdtmpl, err := Loadf("C:\\Users\\hoffman\\Documents\\GitHub\\gnuc\\execute\\command.json")

	if err != nil {
		t.Error("could not read the template file.")
		log.Fatal("error", err)
	}

	if cmdtmpl.Attributes["url"] != "/api/execute" {
		t.Error("something wrong with the template file")
	}
}