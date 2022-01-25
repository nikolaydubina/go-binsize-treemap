package symtab

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// GoSymtabParser parser symtab files produced by `go tool nm`.
// https://pkg.go.dev/cmd/nm
type GoSymtabParser struct{}

func (s GoSymtabParser) ParseSymtab(lines []string) (*SymtabFile, error) {
	var f SymtabFile

	f.Entries = make([]SymtabEntry, 0, len(lines))

	for i, line := range lines {
		e, err := parseGoSymtabLine(line)
		if err != nil {
			err := fmt.Errorf("error parasing symtab file at line num(%d): %w: line: %s", i, err, line)
			log.Println(err.Error())
			continue
		}
		f.Entries = append(f.Entries, e)
	}

	return &f, nil
}

func parseGoSymtabLine(line string) (SymtabEntry, error) {
	var rawAddress, rawSize, rawType, rawSymbolName string

	fields := strings.Fields(line)
	numFields := len(fields)
	switch {
	case numFields > 4:
		// this CAN be Go symbols with type struct names that hard to parse: 10113fdc0        192 T type..eq.struct { github.com/gohugoio/hugo/source.FileWithoutOverlap; github.com/gohugoio/hugo/resources/page.DeprecatedWarningPageMethods1 }
		// assuming so
		rawAddress = fields[0]
		rawSize = fields[1]
		rawType = fields[2]
		rawSymbolName = strings.Join(fields[3:], " ")
	case numFields == 4:
		// normal: "101ae42a0          4 R $f32.3de978d5"
		rawAddress = fields[0]
		rawSize = fields[1]
		rawType = fields[2]
		rawSymbolName = fields[3]
	case numFields == 3:
		// undefined: "       4294971392 U _CFArrayGetCount"
		rawSize = fields[0]
		rawType = fields[1]
		rawSymbolName = fields[2]
	default:
		return SymtabEntry{}, errors.New("wrong number of elements in line")
	}

	var entry SymtabEntry

	entry.Address = rawAddress
	size, err := strconv.Atoi(rawSize)
	if err != nil {
		return SymtabEntry{}, fmt.Errorf("wrong size field: %w", err)
	}
	entry.Size = uint(size)
	entry.Type = SymbolType(rawType)
	entry.SymbolName = rawSymbolName

	if len(fields) == 3 {
		if entry.Type != Undefined {
			return entry, errors.New("got 3 fields but have non undefined type")
		}
	}

	return entry, nil
}
