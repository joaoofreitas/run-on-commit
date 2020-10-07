package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "os"
    "os/exec"
    "flag"
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

    return string(responseObject.Content.SHA)
}

func runScript(script string){
    cmd := exec.Command("./" + script)
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func main() {
    scriptFile := flag.String("file", "", "-file <nameofthefile>") 
    username := flag.String("user", "", "-user <githubusername>")
    repositoryName := flag.String("repo", "", "-repo <nameofrepository>")

    flag.Parse()

    file, _ := ioutil.ReadFile("sha")
    commit := getCommitSHA("http://api.github.com/repos/" + *username + "/" + *repositoryName +"/git/refs/heads/master")

    if string(file) != commit {
	fmt.Println("There are some changes, updating containers...")
	ioutil.WriteFile("sha", []byte(commit), 0644)

	runScript(*scriptFile)
    } else {
	fmt.Println("Nothing changed...Exiting")
    }

}

