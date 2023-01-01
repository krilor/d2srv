# d2srv

**Embed d2 graphs anywhere**

## Why?

I got excited when learning about the release of [d2](https://d2lang.com/). While playing around with it, I thought it would be nice to have something similar as [gravizo](https://www.gravizo.com/) so I can embed graphs anywhere and everywhere. First target is github READMEs.

It had been a while since I wrote any go code, so I thought, heck, lets give it a shot.

## How it works

I'm currently serving an online version of this at d2.reperto.no. To use it, include something like this in your markdown here on github.

```markdown
![d2.reperto.no is served with cloudflare as a cache and the origin server on OCI](https://d2.reperto.no/svg?direction:right;browser->cloudflare;cloudflare->oci.nginx;oci.nginx->oci.d2srv;)

```

It will look something like this:

![d2.reperto.no is served with cloudflare as a cache and the origin server on OCI](https://d2.reperto.no/svg?direction:right;browser->cloudflare;cloudflare->oci.nginx;oci.nginx->oci.d2srv;)

Make sure to:

1. put everything on one line
2. terminate everything with semicolon `;`


## How to make a graph

Use the [d2 playground](https://play.d2lang.com) to fiddle with graphs.

## How to use it on your own

Just compile and run the go binary or simply:

```go
go run main.go
```

## Things to think about going forward

* https://golang.org/issue/25192 (semicolon in query) might cause some issues
* Better ways to include data for graphs + customizations.

## Licence

See [LICENCE].
