package main

import "time"
import "fmt"

func main() {
    stop := make(chan struct{})
    go A(stop)
    go B(stop)
    time.Sleep(2 * time.Second)
    stop <- struct{}{}
    time.Sleep(1 * time.Second)
}

func A(stop <-chan struct{}) {
    select {
    case <- stop:
        fmt.Println("A")
    }
}

func B(stop <-chan struct{}) {
    select {
    case <- stop:
        fmt.Println("B")
    }
}
