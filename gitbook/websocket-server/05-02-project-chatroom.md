# Build a Chatroom

We are going to build a full stack application using React/TypeScript frontend and Go backend. The Go backend is a server that can receive websocket messages from clients and broadcast them to other clients.

## Prerequisite

Part 3 of [Go Tour](https://tour.golang.org/concurrency/1).

## Project Requirement

## Dependency Management - `dep`

Dep is an awesome dependency management tool for Go, it's equivalent to `npm` for JavaScript. You can learn more about [Dep](https://github.com/golang/dep).

### Installation

#### `go-get`

The following command will install `dep` to your Go workspace. When you run `dep` in terminal, the binary is discoverable because you have your GOPATH configured.

```text
go get -u github.com/golang/dep/cmd/dep
```

#### `brew`

If you are a Mac user, there is one more convenient option for you. You can use Homebrew! Simply run the following.

```text
brew install dep
brew upgrade dep
```

### Run `dep`

Once `dep` is installed, you can run `dep init` to initialize a dep environment for your project.

```text
cd $GOPATH/src/go-academy/your_project/
dep init
```

You will see `Gopkg.lock` and `Gopkg.toml` files. Now run `dep ensure`.

```text
dep ensure -v
```

## Frontend

You can copy and paste the front end code in my `chatroom/frontend` folder. However, please feel free to write your own frontend implementation.

### WebSocket

We are going to use JavaScript's native `WebSocket` class to establish websocket connection to the server.

```javascript
this.ws = new WebSocket("ws://localhost:8000/streams/messages")
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

### Node Modules

I am using babel and webpack for compiling the latest ES6/ES7 syntax into browser compatible version. I am also using `node-sass` for compiling `.scss` into `.css`. I make promise-based requests to server using `axios`. For the complete list of node modules, please look at the `package.json`.

### Build `index.js`

Navigate to the `frontend/` folder and run the following npm commands.

```text
cd frontend/
npm install
npm run build
```

**Remember** to run `npm run build`, otherwise you won't see an `index.js` in your `public` folder.

## Project Chatroom

* Video - [Lesson 4 Chatroom Naive Implementation](https://youtu.be/6vS6wYLbyjg)
* Video - [Lesson 4 Chatroom Smart Implementation](https://youtu.be/Q29wM5sYKiw)

## Source

[GitHub](https://github.com/calvinfeng/go-academy/tree/master/chatroom)

