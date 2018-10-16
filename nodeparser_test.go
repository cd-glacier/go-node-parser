package nodeparser

import (
	"testing"
)

func TestParseStmt(t *testing.T) {
	code := `
tests := []struct{
	input string
	output string
}{
	{
		"input0",
		"output0",
	},
	{
		"input1",
		"output1",
	},
	{
		"input2",
		"output2",
	},
}	`

	_, err := ParseStmt(code)
	if err != nil {
		t.Fatalf("Failed to ParseStmt: %s", err.Error())
	}
}
