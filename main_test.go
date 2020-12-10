package main

import (
	"bytes"
	"io/ioutil"
	"netshare/cmd"
	"testing"
)

func TestHelp(t *testing.T) {
	t.Log("Root cmd - Show help")

	cmd := cmd.NewRootCmd()
	result := bytes.NewBufferString("")

	cmd.SetOut(result)
	cmd.SetArgs([]string{"-h"})

	cmd.Execute()
	out, err := ioutil.ReadAll(result)
	if err != nil {
		t.Fatal(err)
	}

	if string(out) == "" {
		t.Fatal(err)
	}
}
