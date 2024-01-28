# Manduco

yeah you know who this is  
it's your one and only manduco in da house!

![manduco the ripper.](manduco_the_ripper.png)

## Usage

Server:

```sh
$ cd server
$ go build
$ ./server
```

Client:

```sh
$ cd client
$ go build
$ ./client <path-to-png-image>
```

## Usage between two different computers

Use a ssh tunnel from client to server:

```sh
ssh -L 1234:localhost:1234 user@server_ip
```

## Add tracing on the client/server

Use `go build -tags trace` instead of `go build` when building the client/server.
