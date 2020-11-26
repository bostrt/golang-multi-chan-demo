package main

import "time"
import "fmt"

func main() {
    stop := make(chan struct{})
    fmt.Println("Configuring A() and B() go-routines as stop channel readers...")
    go A(stop)
    go B(stop)
    time.Sleep(2 * time.Second)
    fmt.Println("Sending stop <- struct{}{}...") 
    stop <- struct{}{}
    time.Sleep(1 * time.Second)

    fmt.Println()
    fmt.Println("Configuring A() and B() again...")
    go A(stop)
    go B(stop)
    time.Sleep(2 * time.Second)
    fmt.Println("Calling close(stop)...")
    close(stop)
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
