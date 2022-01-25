# 🔍 Go binary size SVG treemap

> Make treemap breakdown of your Go binary

```
$ go install github.com/nikolaydubina/go-binsize-treemap@latest
$ go tool nm -size <binary finename> | go-binsize-treemap > binsize.svg
```

## Disclaimer

Should you be worried about your binary size?
In 2022, few seconds of cats videos on tiktok are larger than your binary.
So probably, you should not worry too much about it.
However, this tool can be useful if you are studying compiler (e.g. interfaces, types, linking)
This can be also useful if you want to study wich 3rd party dependencies are getting included and which take lots of size or have lots of code or embedded data.
Next, this can be useful if you are doing `cgo` and want to see how much of C vs Go is included.
Finally, I build this in spare time as another usecase for treemap tooling I built before.

## Examples

_<b><p align="center">github.com/gohugoio/hugo</p></b>_
_<p align="center">62MB, this famous example of large Go project</p>_
![](./docs/hugo.svg)

_<b><p align="center">github.com/cockroachdb/cockroach</p></b>_
_<p align="center">71MB, this famous db is building with C++</p>_
![](./docs/cockroach.svg)

_<b><p align="center">github.com/goccy/go-graphviz</p></b>_
_<p align="center">6.5MB, this project has CGO and builds with lots of graphviz code in C</p>_
![](./docs/go-graphviz.svg)

## Knowledge Base

> What is `go.itab`?

This is interface related code.
Refer to this [article](https://research.swtch.com/interfaces) by Russ Cox.

> What is `runtime.pclntab`? And why it is so big?

As investigated Cockroach team, it is Go runtime structure for traces ([reference](https://www.cockroachlabs.com/blog/go-file-size/)).
Past discussions in GitHub [thread](https://github.com/golang/go/issues/36313) on why it is big and what to do about it (well, nothing).

## Known Issues and TODOs

- [ ] Size slightly mismatches actual binary size. Including unknown does not help.
- [ ] C++
- [ ] identify go:embed
- [ ] color by type + increasing luminance (sys; user; c++; go:embed; etc.)
- [ ] color by symbol type
- [ ] heat by ????

## Related Work

- https://github.com/knz/go-binsize-viz — this was an inspiration for current tool. However, instead of Python and D3 and Javascript, this tool is using single stack purely in Go and has test coverage. Arguably, the downside it is not interactive.
- https://github.com/jondot/goweight — looks like it was working in the beginning, but as of 2022-01-22 it does not work anymore for me and there were reports dating back to 2020-01-23 for it to be not accurate.

## Reference

- https://github.com/knz/go-binsize-viz
- https://github.com/jondot/goweight
- https://github.com/nikolaydubina/treemap
- https://github.com/nikolaydubina/go-cover-treemap
- https://github.com/golang/go/blob/master/src/cmd/nm/doc.go
- https://linux.die.net/man/1/c++filt
- https://github.com/goccy/go-graphviz
- https://research.swtch.com/interfaces

## Appendix A: Strange Output / C++ / CGO

You many need to demungle symtab file first. Install `c++flit`. Then process symtab first.
Note, c++ support is work in progress.

```
$ go tool nm -size <binary finename> | c++filt | go-binsize-treemap > binsize.svg
```

### Appendix B: Large dimensions and lots of details

If you set dimensions very large you can see lots of details and navigate map.

4096x4096 is recommended

![](./docs/hugo-4096x4096.svg)

... but you can go much higher
![](./docs/hugo-16384x16384.svg)

### Appendix C: Small dimensions and informative preview

You can generate small preview of project that fits for embedding in README for example.

1024x256 is recommended

![](./docs/hugo-1024x256.svg)
