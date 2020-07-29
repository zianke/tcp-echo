# tcp-echo
tcp-echo is a tiny server that accepts TCP requests and responds with some provided text. This is inspired by [hashicorp/http-echo](https://github.com/hashicorp/http-echo) and can be useful for demos.

The default port is 5678, but this is configurable via the `-listen` flag:

```
tcp-echo -listen=:8080 -text="hello world"
```

Then run `nc localhost 8080` in your terminal.

## Use with Docker

[Docker image](https://hub.docker.com/r/zianke/tcp-echo)

```
docker pull zianke/tcp-echo
docker run -p 8080:8080 zianke/tcp-echo -listen=:8080 -text="hello world"
```
