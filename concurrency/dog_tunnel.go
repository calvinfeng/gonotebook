package main

import (
  "fmt"
  "time"
)

type Dog struct {
  Name string
  Age int
}

func AgeTheDog(in chan *Dog, out chan *Dog) {
  for dog := range in {
    dog.Age += 1
    time.Sleep(time.Second)
    out <- dog
  }
}

func main() {
  inTunnel := make(chan *Dog)
  outTunnel := make(chan *Dog)

  go AgeTheDog(inTunnel, outTunnel)

  loki := &Dog{
    Name: "Loki",
    Age: 0,
  }

  for {
    inTunnel <- loki
    <- outTunnel

    fmt.Println(loki)
  }
}
