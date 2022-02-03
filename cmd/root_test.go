package cmd

import (
	"bytes"
	"testing"
)

func TestRootCmda(t *testing.T) {
	buf := bytes.NewBufferString("")
	rootCmd.SetOut(buf)

	_ = rootCmd.Execute()

	out := buf.String()
	if out != "Hello Burak, First go project to commit!!!" {
		t.Errorf("it should be equal %s == Merhaba Go Türkiye!!", out)
	}
}
