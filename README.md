# Running the server

To run this, you just do `go run chat_server.go` from the project directory.
The server should tell you that it's listening for connections on port 2222
(which you can change by modifying the part of `chat_server.go` that says "2222").

Quit the server using your favorite instrument of process extermination. I like Ctrl+C.

# Connecting to the server as a local client

On OSX and maybe most Unices, you can connect to the server with `netcat`:

```
nc localhost 2222
```

# Connecting to the server as a remote client

Same thing as far as I know. I was able to make it work from a second computer on the same LAN with:

```
nc 192.168.0.6 2222
```
