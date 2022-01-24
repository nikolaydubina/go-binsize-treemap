package symtab

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/nikolaydubina/treemap"
)

// no heat
// size is Bytes size
type BasicSymtabConverter struct {
	MaxDepth           uint // number of levels from root, including, if 0 then no limit
	IncludeUnknown     bool
	IncludeSymbols     bool
	IncludePureSymbols bool
	ShowSizeBytes      bool
	Verbosity          uint
}

func (s BasicSymtabConverter) SymtabFileToTreemap(sf SymtabFile) treemap.Tree {
	if len(sf.Entries) == 0 {
		return treemap.Tree{}
	}

	tree := treemap.Tree{
		Nodes: map[string]treemap.Node{},
		To:    map[string][]string{},
	}

	hasParent := map[string]bool{}

	for _, entry := range sf.Entries {
		if !s.IncludeUnknown && entry.Type == Undefined {
			continue
		}

		symbolName := parseSymbolName(entry.SymbolName)
		if !s.IncludePureSymbols && len(symbolName.PackageParts) == 0 {
			continue
		}

		parts := symbolName.PackageParts
		if s.IncludeSymbols {
			parts = append(parts, symbolName.SymbolParts...)
		}

		if s.MaxDepth > 0 && len(parts) > int(s.MaxDepth) {
			parts = parts[:s.MaxDepth]
		}

		nodeName := strings.Join(parts, "/")

		if _, ok := tree.Nodes[nodeName]; ok {
			if s.Verbosity > 0 {
				log.Printf("got duplicate node(%s)", nodeName)
			}
			continue
		}

		nodeNameDisplay := nodeName
		if s.ShowSizeBytes {
			count, suffix := byteCountIEC(entry.Size)
			nodeNameDisplay = fmt.Sprintf("%s %f.1%s", nodeNameDisplay, count, suffix)
		}

		tree.Nodes[nodeName] = treemap.Node{
			Path: nodeName,
			Size: float64(entry.Size),
			Name: nodeNameDisplay,
		}

		hasParent[parts[0]] = false

		for parent, i := parts[0], 1; i < len(parts); i++ {
			child := parent + "/" + parts[i]

			tree.Nodes[parent] = treemap.Node{
				Path: parent,
			}

			tree.To[parent] = append(tree.To[parent], child)
			hasParent[child] = true

			parent = child
		}
	}

	for node, v := range tree.To {
		tree.To[node] = unique(v)
	}

	var roots []string
	for node, has := range hasParent {
		if !has {
			roots = append(roots, node)
		}
	}

	switch {
	case len(roots) == 0:
		log.Fatalf(errors.New("no roots, possible cycle in graph").Error())
	case len(roots) > 1:
		tree.Root = "some-secret-string"
		tree.To[tree.Root] = roots
	default:
		tree.Root = roots[0]
	}

	return tree
}

func byteCountIEC(b uint) (float64, string) {
	const unit = 1024
	if b < unit {
		return float64(b), "B"
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return float64(b) / float64(div), string("KMGTPE"[exp])
}

func unique(a []string) []string {
	u := map[string]bool{}
	var b []string
	for _, q := range a {
		if _, ok := u[q]; !ok {
			u[q] = true
			b = append(b, q)
		}
	}
	return b
}
