# More Go Concurrency Patterns

## Parallel Search

Let's write a program that will make multiple Google searches in parallel and return all the results
once all of the responses have come back.

```go
type search func(query string) string

func getSearch(kind string) search {
  return func(query string) string {
    // Optional: implement actual HTTP requests.
    time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
    return fmt.Sprintf("%s result for %q\n", kind, query)
  }
}

var (
  web   = getSearch("web")
  image = getSearch("image")
  video = getSearch("video")
)
```

Now we can write our `google` function.

```go
func google(query string) (results []string) {
  ch := make(chan string)

  go func() {
    ch <- web(query)
  }()

  go func() {
    ch <- image(query)
  }()

  go func() {
    ch <- video(query)
  }()

  for i := 0; i < 3; i++ {
    r := <-ch
    results = append(results, r)
  }

  return results
}
```

We run it in main.

```go
func main() {
  rand.Seed(time.Now().UnixNano())

  start := time.Now()
  results := google("golang")
  elapsed := time.Since(start)

  fmt.Println(results)
  fmt.Println(elapsed)
}
```

## Search with Timeout

We can make our search a bit more robust by introducing timeout because there's a chance that some
results won't come back for a while and we don't want to wait for it.

```go
func google(query string) (results []string) {
  ch := make(chan string)

  go func() {
    ch <- web(query)
  }()

  go func() {
    ch <- image(query)
  }()

  go func() {
    ch <- video(query)
  }()

  timeout := time.After(80 * time.Millisecond)
  for i := 0; i < 3; i++ {
    select {
    case r := <-ch:
      results = append(results, r)
    case <-timeout:
      fmt.Println("timed out")
      return results
    }
  }

  return results
}
```

## Distributed Search

Suppose we can make requests to make multiple slave/replica servers and only return searches that
complete in under 80 milliseconds.

```go
func multisearch(query string, replicas ...search) string {
  ch := make(chan string)

  searchReplica := func(i int) {
    ch <- replicas[i](query)
  }

  for i := range replicas {
    go searchReplica(i)
  }

  // As soon as one of the replicas returns a result, return immediately.
  return <-ch
}
```

Now we can make requests to multiple slave servers.

```go
var (
  web1   = getSearch("web")
  web2   = getSearch("web")
  image1 = getSearch("image")
  image2 = getSearch("image")
  video1 = getSearch("video")
  video2 = getSearch("video")
)

func google(query string) (results []string) {
  ch := make(chan string)

  go func() {
    ch <- multisearch(query, web1, web2)
  }()

  go func() {
    ch <- multisearch(query, image1, image2)
  }()

  go func() {
    ch <- multisearch(query, video1, video2)
  }()

  timeout := time.After(80 * time.Millisecond)
  for i := 0; i < 3; i++ {
    select {
    case r := <-ch:
      results = append(results, r)
    case <-timeout:
      fmt.Println("timed out")
      return results
    }
  }

  return results
}
```