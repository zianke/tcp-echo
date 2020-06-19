# tcp-echo
tcp-echo is a tiny server that accepts TCP requests and responds with some provided text. This is inspired by [hashicorp/http-echo](https://github.com/hashicorp/http-echo) and can be useful for demos.

The default port is 5678, but this is configurable via the `-listen` flag:

```
tcp-echo -listen=:8080 -text="hello world"
```

Then run `nc localhost 5678` in your terminal.
