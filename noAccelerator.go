package main

import (
        "net/http"
        "fmt"
		"os"
        "regexp"
        "bufio"
        "log"
        "crypto/tls"
        "io/ioutil"
//        "sync"
//        "time"
        "flag"
)

func checkUrl(eachline string, header *string){
	  client := &http.Client{}
	req, err := http.NewRequest("GET", eachline, nil)
// ...
req.Header.Add("X-Bug-Bounty", "MrMustacheMan")
resp, err := client.Do(req)

if err != nil {
        return
}
bodyBytes, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                fmt.Println(err)
        }

matched, err := regexp.MatchString("Accelerator", string(bodyBytes)) 
if err != nil {
        fmt.Println(err)
        }
if !matched {
                fmt.Printf("%s\n", eachline)
        } 
}

func main() {        

        http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
        file := flag.String("f", "", "file to use for parsing")
        var header = flag.String("h", "", "H1 Username to pass to the request")
        flag.Parse()
        
        readFile, err := os.Open(*file)
 
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
        
        	checkUrl(eachline, header)
}
}
