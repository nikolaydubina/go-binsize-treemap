package main

import (
	"bufio"
	"flag"
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/nikolaydubina/go-binsize-treemap/symtab"
	"github.com/nikolaydubina/treemap"
	"github.com/nikolaydubina/treemap/render"
)

const doc string = `
Go binary size treemap.

Examples
$ go tool nm -size <binary finename> | go-binsize-treemap > binsize.svg
$ go tool nm -size <binary finename> | c++filt | go-binsize-treemap > binsize.svg

Command options:
`

var grey = color.RGBA{128, 128, 128, 255}

func main() {
	var (
		coverprofile   string
		w              float64
		h              float64
		marginBox      float64
		paddingBox     float64
		padding        float64
		maxDepth       uint
		includeSymbols bool
	)

	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), doc)
		flag.PrintDefaults()
	}
	flag.StringVar(&coverprofile, "coverprofile", "", "filename of input coverprofile (e.g. cover.out)")
	flag.Float64Var(&w, "w", 1028, "width of output")
	flag.Float64Var(&h, "h", 640, "height of output")
	flag.Float64Var(&marginBox, "margin-box", 4, "margin between boxes")
	flag.Float64Var(&paddingBox, "padding-box", 4, "padding between box border and content")
	flag.Float64Var(&padding, "padding", 16, "padding around root content")
	flag.UintVar(&maxDepth, "max-depth", 0, "if zero then no max depth is set, else will show only number of levels from root including")
	flag.BoolVar(&includeSymbols, "symbols", true, "include symbols or not")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	parser := symtab.GoSymtabParser{}
	symtabFile, err := parser.ParseSymtab(lines)
	if err != nil || symtabFile == nil {
		log.Fatal(err)
	}

	converter := symtab.BasicSymtabConverter{
		MaxDepth:       maxDepth,
		IncludeSymbols: includeSymbols,
	}
	tree := converter.SymtabFileToTreemap(*symtabFile)

	sizeImputer := treemap.SumSizeImputer{EmptyLeafSize: 1}
	sizeImputer.ImputeSize(tree)

	treemap.SetNamesFromPaths(&tree)
	treemap.CollapseRoot(&tree)

	uiBuilder := render.UITreeMapBuilder{
		Colorer:     render.NoneColorer{},
		BorderColor: grey,
	}
	spec := uiBuilder.NewUITreeMap(tree, w, h, marginBox, paddingBox, padding)
	renderer := render.SVGRenderer{}

	os.Stdout.Write(renderer.Render(spec, w, h))
}
