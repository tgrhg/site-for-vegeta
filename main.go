package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano()) // 現在時刻で乱数のシードを初期化

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if rand.Intn(5) < 1 { // 0-4の乱数を生成し、20%の確率で0になる
            http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
            return
        }

        switch r.Method {
        case http.MethodGet:
            fmt.Fprintln(w, "Hello, GET request received!")
        case http.MethodPost:
            fmt.Fprintln(w, "Hello, POST request received!")
        case http.MethodPut:
            fmt.Fprintln(w, "Hello, PUT request received!")
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    http.ListenAndServe(":8080", nil)
}

