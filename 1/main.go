package main

/*

References:
https://medium.com/@marcus.olsson/writing-a-go-client-for-your-restful-api-c193a2f4998c
*/

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Client struct {
	BaseUrl    *url.URL
	httpClient *http.Client
	UserAgent  string
}

type User struct {
}

func (c *Client) ListUsers() ([]User, error) {
	rel := &url.URL{Path: "/users"}
	u := c.BaseUrl.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println(rel)
	fmt.Println(u)
	return []User{}, nil
}

func main() {

	if len(os.Args) != 2 {
		log.Fatal("Takes exactly one input argument. You entered ", os.Args)
	}
	log.Println("Input args", os.Args)
	args := os.Args[1]

	fmt.Println(args)

	c := Client{httpClient: http.DefaultClient, UserAgent: "gogo"}
	c.ListUsers()

}
