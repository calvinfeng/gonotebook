package main

import (
  "fmt"
  "time"
)

func BlockingFunc() {
  for {
    fmt.Println("I like to block you for a second.")
    time.Sleep(time.Second)
  }
}

func GetToWork() {
  for {
    fmt.Println("Please let me do my work")
    time.Sleep(time.Second)
  }
}

func main(){
  go BlockingFunc()
  GetToWork()
}
