# Slack Clone

Now it is time to put everything you've learned together to create one final capstone project to test your understanding. Exercise your creativity and combine Chatroom and User Authentication projects together into one; create a Slack clone!

## Server-side Requirements

* User can create account and authenticate to server.
* User can send message to each other and the change should be broadcast via websocket.
* User can join a **channel** just like the real Slack.
* Messages must be sent to a channel or directly from one user to another user.
* Messages should be stored for future retrieval.
* Server must be **highly concurrent**, it can handle high volume of messages in real time.

It is your job to architect the following items.

* Data model and database schema \(either SQL or NoSQL.\)
* HTTP & WebSocket endpoints
* Message demultiplexing/multiplexing
* Authentication logic

## **Client-side Requirements**

* User can enter email and password to login or sign up.
* User can see and join any channel.
* User can read historical messages of a channel and historical messages of directed message.
* User can send messages to each other via direct message or channel, and receive live message updates from other users.
* Browser should **not** need to refresh to receive any updates, all updates should be shown immediately without the user intervention. 
* Channel should indicate unread messages if there is any.

It is your job to architect the following items.

* React presentation and container components, i.e. which one should be connected with Redux store
* WebSocket reconnect logic, if the server ever goes down, the socket should reconnect itself once server comes back online
* Handle live-updates from the server and render them to user

## Testing Requirement

Each endpoint and React component should be thoroughly tested. If you want to write good test, the key is **dependency injection.**

* Enzyme/Jest for React
  * Rendering logic on shallow & full DOM
  * Business logic on each helper and utility functions
* Testify for Golang
  * HTTP handler logic
  * WS handler logic

## Collaboration

One of you should create a GitHub repository to host this project. Use GitHub's pull request and code review feature to perform collaborative coding. 



