# Client-Server TCP/UDP Sample

## Rock, Paper & Scissors - Jokenpo

This is a simple sample, elaborated to demonstrate how to use Go's TCP and UDP implementation over a distributed ecosystem.

- Execute each file separately over terminal tabs

Start running server two:
```bash
$ go run server-two-<tcp/udo>.go
```

Next, run server one:
```bash
$ go run server-one-<tcp/udo>.go
```

and finally, the client:
```bash
$ go run client-<tcp/udp>.go
```

Type 'R' for Rock, 'P' for Paper or 'S' for Scissor:
```
> (Choose R, P or S): R
> Result: It's a Draw! User > Rock vs Rock < Brain

> (Choose R, P or S): P
> Result: It's a Draw! User > Paper vs Paper < Brain

> (Choose R, P or S): S
> Result: You Lose! User > Scissor vs Rock < Brain
```

## Stopping Sample:
Just `CTRL + C` over the terminal instances.