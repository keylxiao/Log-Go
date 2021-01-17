package main

import (
    "Log-Go/config"
    "fmt"
)

func main() {

   message := map[string]string{
     "error": "Timeout",
     "grade": "C",
     "ip": "127.0.101.1" ,
     "requetaddress": "/Product" ,
     "time": "2018-3-7",
   }
    _ = config.Write("./config/error.toml", message)
    content, _ := config.Read("./config/error.toml")
    fmt.Println(content)
}
