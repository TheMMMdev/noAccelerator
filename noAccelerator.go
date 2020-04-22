package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	file := os.Args[1]
	readFile, err := os.Open(file)

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	for _, eachline := range fileTextLines {

		client := &http.Client{}

		req, err := http.NewRequest("GET", eachline, nil)
		// ...
		req.Header.Add("X-Bug-Bounty", "MrMustacheMan")
		resp, err := client.Do(req)

		if err != nil {
			panic(err)
		}
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		matched, err := regexp.MatchString("Not Found on Accelerator", string(bodyBytes))
		if err != nil {
			fmt.Println(err)
		}
		if !matched {
			fmt.Printf("%s\n", eachline)
		}
		// ...
	}
}
