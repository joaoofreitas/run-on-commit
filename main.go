package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

func getCommitSHA(url string) string{ 
    type CommitSHA struct {
	SHA string `json:"sha"`
    }
    type Response struct {
	Content CommitSHA `json:"object"`
    }

    res, _ := http.Get(url) 

    responseInBytes, _ := ioutil.ReadAll(res.Body)

    var responseObject Response
    json.Unmarshal(responseInBytes, &responseObject)

    fmt.Println(responseObject.Content.SHA)
    return string(responseObject.Content.SHA)
}
func main() {

    file, _ := ioutil.ReadFile("sha")
    commit := getCommitSHA("http://api.github.com/repos/joaoofreitas/vue-portfolio/git/refs/heads/master")

    if string(file) != commit{
	// Run Update Function
	fmt.Println("There are some changes, updating containers...")
	ioutil.WriteFile("sha", []byte(commit), 0644)
    } else {
	fmt.Println("Nothing changed...Exiting")
    }

}

/*
Todo: 
    Store ---> responseObject.Content.SHA in a file. DONE
    @START ---> read content in the file and perform the request DONE
    Add argv for user and repository
    Write the conditions
*/

