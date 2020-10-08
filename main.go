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

    res, err := http.Get(url)
    if err != nil {
	fmt.Println("Error fetching the GitHub API, please check your internet connection. Exiting...")
	os.Exit(3)
    }

    responseInBytes, err := ioutil.ReadAll(res.Body)
    if err != nil {
	fmt.Println("Error reading the GitHub API. Exiting...")
	os.Exit(3)
    }

    var responseObject Response
    json.Unmarshal(responseInBytes, &responseObject)
    if len(responseObject.Content.SHA) == 0{
	fmt.Println("The username or the repository flagged does not exist. Please check the username and the repository name.")
	os.Exit(3)
    }
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

    file, err := ioutil.ReadFile("sha")

    if err != nil {
	fmt.Println("Error reading the SHA memory file. Please check if the sha file exists")
	os.Exit(3)
    }

    commit := getCommitSHA("http://api.github.com/repos/" + *username + "/" + *repositoryName +"/git/refs/heads/master")


    if string(file) != commit {
	fmt.Println("There are some changes, updating containers...")
	ioutil.WriteFile("sha", []byte(commit), 0644)

	runScript(*scriptFile)
    } else {
	fmt.Println("Nothing changed...Exiting")
    }

}

