package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func PrettyWrite(r io.Reader, w io.Writer) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, b, "", "    ")

	if error != nil {
		log.Fatal("JSON parse error: ", error)
	}

	f := bufio.NewWriter(w)
	defer f.Flush()

	// Add newline at end of text if missing
	bs := prettyJSON.Bytes()
	bs = append(bytes.TrimSpace(bs), '\n')

	f.Write(bs)
}

func main() {
	r := os.Stdin
	w := os.Stdout
	PrettyWrite(r, w)

}
