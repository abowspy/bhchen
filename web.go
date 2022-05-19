package main

import (
    "fmt"
    "net/http"
    "log"
    "os"
    "net"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {

    // 取得IP
    ip, _, _ := net.SplitHostPort(r.RemoteAddr)

    // 取得hostname
    host, _ := os.Hostname()

    fmt.Fprintf(w, "v3 web -> " + host + " from " + ip + "\n")     //這個寫入到 w 的是輸出到客戶端
}

func main() {
    http.HandleFunc("/", sayhelloName)          //設定存取的路由
    err := http.ListenAndServe(":80", nil)      //設定監聽的port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}