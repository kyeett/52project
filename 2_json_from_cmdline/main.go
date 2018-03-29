package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var dat map[string]interface{}

	for scanner.Scan() {

		fmt.Println(scanner.Text())

		if err := json.Unmarshal([]byte(scanner.Text()), &dat); err != nil {
			panic(err)
		}

		str, err := json.MarshalIndent(dat, "", "    ")
		if err != nil {
			log.Fatal(err)
		}

		str = append(str, '\n')
		os.Stdout.Write(str)
	}
}

//https://gist.github.com/jd-boyd/119b290f881a0148b515
