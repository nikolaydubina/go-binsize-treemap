# 🗺 Go binary size treemap

Install
```
$ go install github.com/nikolaydubina/go-binsize-treemap@latest
# get c++flit
```

Then run
```
go tool nm -size <binary finename> | c++filt | go-binsize-treemap > binsize.svg
```

## Related Work

- https://github.com/knz/go-binsize-viz — this was an inspiration for current tool. However, instead of Python and D3 and Javascript, this tool is using single stack purely in Go and has test coverage. Arguably, the downside it is not interactive.
- https://github.com/jondot/goweight — looks like it was working in the beginning, but as of 2022-01-22 it does not work anymore for me and there were reports dating back to 2020-01-23 for it to be not accurate.

## Reference

- https://github.com/knz/go-binsize-viz
- https://github.com/jondot/goweight
- https://github.com/nikolaydubina/treemap
- https://github.com/nikolaydubina/go-cover-treemap