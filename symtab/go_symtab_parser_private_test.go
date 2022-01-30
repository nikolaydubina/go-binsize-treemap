package symtab

import "testing"

func Test_parseGoSymtabLine(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expEntry SymtabEntry
	}{
		{
			name: "when 3 elemnets and U, then ok",
			line: "4294971392 U _CFArrayGetCount",
			expEntry: SymtabEntry{
				Size:       4294971392,
				Type:       Undefined,
				SymbolName: "_CFArrayGetCount",
			},
		},
		{
			name: "when 4 elemnets, then ok",
			line: "101ae42a0          4 R $f32.3de978d5",
			expEntry: SymtabEntry{
				Address:    "101ae42a0",
				Size:       4,
				Type:       ReadOnly,
				SymbolName: "$f32.3de978d5",
			},
		},
		{
			name: "when more then 4 elemnets, then ok-ish with some symbol name",
			line: "10113fdc0        192 T type..eq.struct { github.com/gohugoio/hugo/source.FileWithoutOverlap; github.com/gohugoio/hugo/resources/page.DeprecatedWarningPageMethods1 }",
			expEntry: SymtabEntry{
				Address:    "10113fdc0",
				Size:       192,
				Type:       Text,
				SymbolName: "type..eq.struct { github.com/gohugoio/hugo/source.FileWithoutOverlap; github.com/gohugoio/hugo/resources/page.DeprecatedWarningPageMethods1 }",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			entry, err := parseGoSymtabLine(tc.line)
			if err != nil {
				t.Errorf("got error(%s), expected none", err)
			}
			if tc.expEntry != entry {
				t.Errorf("exp %#v != got %#v", tc.expEntry, entry)
			}
		})
	}
}
