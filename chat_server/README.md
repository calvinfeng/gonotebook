# WebSocket Server in Go
## Prerequisite
Now it's time for you to finish Part 3 of [Go Tour](https://tour.golang.org/concurrency/1). The 
last four sections are not required for the websocket project. You only need to finish sections 
from *Goroutines* to *Default Selection*. Later we will re-visit `sync.mutex` when we dive deeper
into concurrency.

## Video 04a: Concurrency

[Concurrency in Go](https://youtu.be/uq9EocsraUQ)

## Project Requirement
We are going to learn about how to perform dependency management in Go and how to use Gorilla 
library to implement a websocket connection in Go. We are also going to learn about how to 
integrate React with Go.

## Dependency Management - `dep`
Dep is an awesome dependency management tool for Go, it's equivalent to `npm` for JavaScript. You can
learn more about [Dep](https://github.com/golang/dep).

### Local Installation
If you wish to install it locally to your project,

```
go get -u github.com/golang/dep/cmd/dep
```

However, I recommend that you do it globally like how you can run `npm` anywhere on your computer.

### Global Installation
If you are a Mac user, the installation step is very easy for you. First of all you should have 
Homebrew on your Mac, then perform the following commands in your terminal:
```
brew install dep
brew upgrade dep
```

And that's it.

## Frontend
You can copy and paste the front end code in my `user_auth/frontend` folder. However, please feel 
free to write your own frontend implementation.

We are going to use JavaScript's native `WebSocket` class.
```javascript
this.ws = new WebSocket("ws://localhost:8000/ws")
this.websocket.onopen = this.handleSocketOpen;
this.websocket.onmessage = this.handleSocketMessage;
this.websocket.onerror = this.handleSocketError;
this.websocket.onclose = this.handleSocketClose;
```

Client-side socket connection typically accepts 4 callbacks:

    1. Callback is invoked when socket is opened.
    2. Callback is invoked when socket receives a message.
    3. Callback is invoked when socket encounters an error.
    4. Callback is invoked when socket is closed.

## Node Modules
I am using babel and webpack for compiling the latest ES6/ES7 syntax into browser compatible version. 
I am also using `node-sass` for compiling `.scss` into `.css`. I make promise-based requests to 
server using `axios` instead of jQuery. I am not a big fan of jQuery anymore. For the complete list of 
dependency, please look at the `package.json`.

## Video 04b: WebSocket Server
Please excuse me that I repetitively said *so* in my speech; it's a result of having to think about what to type and what
to say simultaneously. I will make sure that in my next video I will suppress my impulse to say *so*.

[WebSocket Server in Go Part 1](https://youtu.be/KtdZinZIe3A)

[WebSocket Server in Go Part 2](https://youtu.be/Ue7z2BEavBU)