package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "os"
    "os/exec"
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

    //fmt.Println(responseObject.Content.SHA)
    return string(responseObject.Content.SHA)
}

func runCommand(commands ...string){
    arr := []int{}
    for i := len(commands) - 1; i >= 0; i-- {
	arr = append(arr, i)
    }
    fmt.Println(arr)
    cmd := exec.Command(commands[arr[]], commands[len(commands) -1])
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func main() {

    file, _ := ioutil.ReadFile("sha")
    commit := getCommitSHA("http://api.github.com/repos/joaoofreitas/vue-portfolio/git/refs/heads/master")

    if string(file) != commit{
	runCommand("ls", "-lah")
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
    Change ---> Run commands as run a script.
    Add argv for user and repository
    Write the conditions
*/

