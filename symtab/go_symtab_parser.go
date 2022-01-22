package symtab

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// GoSymtabParser parser symtab files produced by `go tool nm`. https://pkg.go.dev/cmd/nm
type GoSymtabParser struct{}

func (s GoSymtabParser) ParseSymtab(lines []string) (*SymtabFile, error) {
	var f SymtabFile

	f.Entries = make([]SymtabEntry, 0, len(lines))

	for i, line := range lines {
		e, err := parseGoSymtabLine(line)
		if err != nil {
			return nil, fmt.Errorf("error parasing symtab file at line(%d): %w", i, err)
		}
		f.Entries = append(f.Entries, e)
	}

	return &f, nil
}

// https://pkg.go.dev/cmd/nm
func parseGoSymtabLine(line string) (SymtabEntry, error) {
	fields := strings.Fields(line)
	if len(fields) < 3 {
		return SymtabEntry{}, errors.New("wrong number of fields, expected 3+")
	}

	var entry SymtabEntry

	entry.Address = fields[0]

	size, err := strconv.Atoi(fields[1])
	if err == nil {
		return SymtabEntry{}, fmt.Errorf("wrong size: %w", err)
	}
	entry.Size = uint(size)

	entry.Type = SymbolType(fields[2])

	if len(fields) > 3 {
		entry.SymbolName = fields[3]
	}

	return entry, nil
}
