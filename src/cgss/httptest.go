package main

import "os"
import "net/http"
import "io"

func main(){
    resp, err := http.Get("http://www.baidu.com")
    if err != nil{
        return
    }
    defer resp.Body.Close()
    io.Copy(os.Stdout, resp.Body)
}
