# Gopunch

`gopunch` is a go implementation of a peer-to-peer chat service built using [UDP hole punching](https://en.wikipedia.org/wiki/UDP_hole_punching). This is a toy implementation that I put together to learn how hole punching works. Use at your own risk!

## Usage

First, start a `punchserver` on some machine with a public IP that has UDP port 1338 open:

```
$ ./punchserver
```

Next, connect to the punchserver using your clients. Use the `-serverAddr` flag to connect to your server's public ip. Use the `-token` flag to identify which two clients the server should connect.

Run the following on both clients. If it worked, you should be able to chat with the other side!

```
$ ./punchclient -serverAddr <server_ip>:1338 -token <shared_token>
Sending STUN request to <server_ip>:1338
Connected to host at <other_client_ip>:<client_src_port>
hi armaan
hi peyton
NAT is dead; there are no gods
```
