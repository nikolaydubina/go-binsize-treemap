# 🔍 Go binary size SVG treemap

```
$ go install github.com/nikolaydubina/go-binsize-treemap@latest
$ go tool nm -size <binary finename> | go-binsize-treemap > binsize.svg
```

hugo, 62MB

cockroach, 71MB

go-graphviz, 6.5MB

## Description

What is `go.itab`? This is interface related code.
Refer to this [article](https://research.swtch.com/interfaces) by Russ Cox.

## TODO

- [ ] C++
- [ ] identify go:embed
- [ ] collapse long chains besides root
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
