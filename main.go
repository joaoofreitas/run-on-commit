package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)
func main() {
    type CommitSHA struct {
	SHA string `json:"sha"`
    }
    type Response struct {
	Content CommitSHA `json:"object"`
    }

    res, _ := http.Get("http://api.github.com/repos/joaoofreitas/vue-portfolio/git/refs/heads/master") 

    responseInBytes, _ := ioutil.ReadAll(res.Body)
    //responseInStr := string(responseInBytes)

    var responseObject Response
    json.Unmarshal(responseInBytes, &responseObject)

    fmt.Println(responseObject.Content.SHA)
}

