package main

import (
    "Log-Go/config"
    "encoding/json"
)

func main() {

   message := map[string]string{
    "error": "Timeoutaaaaaa",
    "grade": "Caaaa",
    "ip": "127.0.101.1" ,
    "requetaddress": "/Product" ,
    "time": "2018-3-7",
   }


    b, _ := json.Marshal(message)
    _ = config.Write("./config/error.toml", b)

}
