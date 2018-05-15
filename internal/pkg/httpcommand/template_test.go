package httpcommand

import (
	"testing"
	"log"
)

func TestExecuteCommand(t *testing.T)  {
	cmdtmpl, err := Loadf("test_command.json")

	if err != nil {
		t.Error("could not read the template file.")
		log.Fatal("error", err)
	}

	if cmdtmpl.Attributes["url"] != "/api/execute" {
		t.Error("something wrong with the template file")
	}
}