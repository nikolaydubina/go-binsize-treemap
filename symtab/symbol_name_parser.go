package symtab

import "strings"

type SymbolName struct {
	PackageParts []string
	SymbolParts  []string
}

func EqSymbolName(a, b SymbolName) bool {
	if len(a.PackageParts) != len(b.PackageParts) {
		return false
	}
	if len(a.SymbolParts) != len(b.SymbolParts) {
		return false
	}
	for i := 0; i < len(a.PackageParts); i++ {
		if a.PackageParts[i] != b.PackageParts[i] {
			return false
		}
	}
	for i := 0; i < len(a.SymbolParts); i++ {
		if a.SymbolParts[i] != b.SymbolParts[i] {
			return false
		}
	}
	return true
}

func parseSymbolName(r string) SymbolName {
	// pure symbol
	if !strings.ContainsAny(r, "./") {
		return SymbolName{SymbolParts: []string{r}}
	}

	// single-part package just symbol
	if !strings.Contains(r, "/") {
		parts := strings.Split(r, ".")
		return SymbolName{
			PackageParts: parts[:1],
			SymbolParts:  parts[1:],
		}
	}

	// has multi-parts package
	lastSlashIdx := strings.LastIndex(r, "/")

	partsPacakge := strings.Split(r[:lastSlashIdx], "/")
	partsSymbol := strings.Split(r[lastSlashIdx:], ".")

	return SymbolName{
		PackageParts: append(partsPacakge, partsSymbol[0][1:]),
		SymbolParts:  partsSymbol[1:],
	}
}
