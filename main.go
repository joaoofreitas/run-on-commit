package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)
func main() {
    res, err := http.Get("http://api.github.com/users/joaoofreitas")

    if err != nil {
	fmt.Println(err)
    }

    reponseData, _ := ioutil.ReadAll(res.Body)
    fmt.Println(string(reponseData))

}

