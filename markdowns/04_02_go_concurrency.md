# Go Concurrency Patterns

Now we know what is a goroutine and let's try some examples of using them. Majority of the examples
come from Rob Pike's talk. I made couple enhancements here and there.

## Channels

Two goroutines can communicate and synchronize through channel.

```go
func ping(msg string, c chan string) {
  for i := 0; ; i++ {
    c <- fmt.Sprintf("%s %d", msg, i)
    time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
  }
}

func main() {
  ch := make(chan string)
  go ping("ping", ch)

  for i := 0; i < 5; i++ {
    fmt.Println(<-ch)
    // The receiver end is blocking because it's waiting for a message to come through the channel.
  }

  fmt.Println("Enough pings, done!")
}
```

## Generator

What if we want to listen to multiple pings? We can use a ping generator.

```go
func pingGen(msg string) <-chan string {
  ch := make(chan string)

  go func() {
    for i := 0; ; i++ {
      ch <- fmt.Sprintf("%s %d", msg, i)
      time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
    }
  }()

  return ch
}

func main() {
  ch1 := pingGen("ping")
  ch2 := pingGen("bing")

  for i := 0; i < 5; i++ {
    fmt.Println(<-ch1)
    fmt.Println(<-ch2)
  }

  fmt.Println("Enough pings, done!")
}
```

Here's a problem though, `ch1` and `ch2` are blocking each other. We are not really getting the live
updates of ping.

## Multiplexing

We can address the problem by multiplexing multiple channels into one.

```go
func fanIn(inputs ...<-chan string) <-chan string {
  out := make(chan string)

  for _, in := range inputs {
    go func(ch <-chan string) {
      for {
        out <- <-ch
      }
    }(in)
  }

  return out
}

func pingGen(msg string) <-chan string {
  ch := make(chan string)

  go func() {
    for i := 0; ; i++ {
      ch <- fmt.Sprintf("%s %d", msg, i)
      time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
    }
  }()

  return ch
}

func main() {
  out := fanIn(pingGen("ping"), pingGen("bing"), pingGen("sing"))

  for i := 0; i < 10; i++ {
    fmt.Println(<-out)
  }

  fmt.Println("Enough pings, done!")
}
```

However, we can easily spot that the messages are coming in out of order. What if I want the batch
of messages to come in order, e.g.

```text
sing 0
bing 0
ping 0
sing 1
ping 1
bing 1
sing 2
ping 2
bing 2
sing 3
ping 3
bing 3
```

## Fan-in Sequencing

We will use a signal channel to achieve the wait. We wait for three pings and then tell them ready
for the next batch of 3 pings. First, let's define a message struct.

```go
type message struct {
  content string
  ready   chan bool
}
```

Each message will carry a reference to a signal channel, called `ready`. When the channel receives a
signal, it indicates that it's ready to process next message.

```go
func fanIn(inputs ...<-chan message) <-chan message {
  out := make(chan message)

  for _, in := range inputs {
    go func(ch <-chan message) {
      for {
        out <- <-ch
      }
    }(in)
  }

  return out
}

func pingGen(msg string) <-chan message {
  ch := make(chan message)
  rdy := make(chan bool)

  go func(ready chan bool) {
    for i := 0; ; i++ {
      ch <- message{fmt.Sprintf("%s %d", msg, i), ready}
      time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
      <-ready // Wait for ready
    }
  }(rdy)

  return ch
}

func main() {
  out := fanIn(pingGen("ping"), pingGen("bing"), pingGen("sing"))

  for i := 0; i < 10; i++ {
    msgs := []message{}
    for j := 0; j < 3; j++ {
      msgs = append(msgs, <-out)
    }

    // Grab three messages and then tell each generator that it's ready for next message.
    for _, msg := range msgs {
      fmt.Println(msg.content)
      msg.ready <- true
    }
  }

  fmt.Println("Enough pings, done!")
}
```