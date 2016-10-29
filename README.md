# Learning Go - Writing a simple Websocket Chat

A simple implementation of a websocket chat server written in [Go] and using the [Gorilla Websocket] library.

## Features
* rendering a chat interface in for the browser
* handling communication with chatters in different chat rooms
* handling user login and logout
 
## Running

Using [Go] and go get:

```
go get bitbucket.org/hascode/go-websocket-chat/chat
go run chat.go
```

Afterwards you may access your chat at http://localhost:8080/

## Using the Library

Simply import _bitbucket.org/hascode/go-websocket-chat/chat_ in your application e.g.:
```
package main

import "bitbucket.org/hascode/go-websocket-chat/chat"

func main() {
	chat.Run()
}

```

## Disclaimer

The chat is not intended for production use, it's just an experiment with Go an websockets..

Please feel free to have a look at my blog at [www.hascode.com] for the full tutorial.

---------

**2016 Micha Kops / hasCode.com**

   [www.hascode.com]:http://www.hascode.com/
   [Gorilla Websocket]:https://github.com/gorilla/websocket
   [Go]:https://golang.org/
