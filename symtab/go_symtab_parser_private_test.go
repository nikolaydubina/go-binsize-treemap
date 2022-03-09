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
		{
			name: "when empty, then ok",
			line: "      0          0 _",
			expEntry: SymtabEntry{
				Address:    "0",
				Size:       0,
				Type:       Underscore,
				SymbolName: "",
			},
		},
		{
			name: "when empty with name, then with name",
			line: "       0          0 _ asn.cpp",
			expEntry: SymtabEntry{
				Address:    "0",
				Size:       0,
				Type:       Underscore,
				SymbolName: "asn.cpp",
			},
		},
		{
			name: "when address with zero size and no symbol, then no symbol name",
			line: "   400338          0 r",
			expEntry: SymtabEntry{
				Address:    "400338",
				Size:       0,
				Type:       StaticReadOnly,
				SymbolName: "",
			},
		},
		{
			name: "when address hex with zero size and no symbol, then no symbol name",
			line: "   55e17b0          0 d",
			expEntry: SymtabEntry{
				Address:    "55e17b0",
				Size:       0,
				Type:       StaticData,
				SymbolName: "",
			},
		},
		{
			name: "when some long C++ string, then correctly address empty and size 0 and symbol name whole thing",
			line: "                    0 U std::basic_ostream<char, std::char_traits<char> >& std::operator<< <std::char_traits<char> >(std::basic_ostream<char, std::char_traits<char> >&, char const*)@@GLIBCXX_3.4",
			expEntry: SymtabEntry{
				Address:    "",
				Size:       0,
				Type:       Undefined,
				SymbolName: "std::basic_ostream<char, std::char_traits<char> >& std::operator<< <std::char_traits<char> >(std::basic_ostream<char, std::char_traits<char> >&, char const*)@@GLIBCXX_3.4",
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
