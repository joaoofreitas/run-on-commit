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

/*
Todo: 
    Store ---> responseObject.Content.SHA in a file.
    @START ---> read content in the file and perform the request
    Add argv for user and repository
    Write the conditions
*/

