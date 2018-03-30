package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var strs []string

	var dat map[string]interface{}

	// Read reader input into []string
	for scanner.Scan() {
		strs = append(strs, scanner.Text(), "\n")
	}
	str := strings.Join(strs, "")

	// QUESTION: Can I MarshalIndent directly?
	if err := json.Unmarshal([]byte(str), &dat); err != nil {
		panic(err)
	}

	prettyStr, err := json.MarshalIndent(dat, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	prettyStr = append(prettyStr, '\n')
	os.Stdout.Write(prettyStr)

}

//https://gist.github.com/jd-boyd/119b290f881a0148b515
