package symtab

import (
	"github.com/nikolaydubina/treemap"
)

// no heat
// size is Bytes size
type BasicSymtabConverter struct {
	MaxDepth       uint // number of levels from root, including, if 0 then no limit
	IncludeSymbols bool
}

func (s BasicSymtabConverter) SymtabFileToTreemap(sf SymtabFile) treemap.Tree {
	if len(sf.Entries) == 0 {
		return treemap.Tree{}
	}

	tree := treemap.Tree{
		Nodes: map[string]treemap.Node{},
		To:    map[string][]string{},
	}

	for _, entry := range sf.Entries {
		if entry.Type == Undefined {
			continue
		}

		if _, ok := tree.Nodes[entry.SymbolName]; ok {
			continue
		}
	}

	return tree
}
