package main_test

import (
	"bytes"
	"strings"
	"testing"

	main "github.com/kyeett/52projects/two_json_from_cmdline"
)

func printErrorNotEqual(t *testing.T, a, b string) {
	t.Errorf(`
"%s" is not equal to
"%s"
`, a, b)
}

func TestSimpleJson(t *testing.T) {
	testWriter := new(bytes.Buffer)
	reader := strings.NewReader(`{"hej":"foo"}`)
	main.PrettyWrite(reader, testWriter)
	expectedPrettyString := `{
    "hej": "foo"
}
`
	if (testWriter.String() == expectedPrettyString) == false {
		printErrorNotEqual(t, testWriter.String(), expectedPrettyString)
	}
}
