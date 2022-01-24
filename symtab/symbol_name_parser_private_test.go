package symtab

import "testing"

func Test_parseSymbolName(t *testing.T) {
	tests := []struct {
		rawSymbolName string
		expSymbolName SymbolName
	}{
		{
			rawSymbolName: "_CFArrayGetCount",
			expSymbolName: SymbolName{
				PackageParts: nil,
				SymbolParts:  []string{"_CFArrayGetCount"},
			},
		},
		{
			rawSymbolName: "$f32.3de978d5",
			expSymbolName: SymbolName{
				PackageParts: []string{"$f32"},
				SymbolParts:  []string{"3de978d5"},
			},
		},
		{
			rawSymbolName: "__rt0_arm64_darwin",
			expSymbolName: SymbolName{
				PackageParts: nil,
				SymbolParts:  []string{"__rt0_arm64_darwin"},
			},
		},
		{
			rawSymbolName: "github.com/clbanning/mxj/v2.elemList.Less",
			expSymbolName: SymbolName{
				PackageParts: []string{"github.com", "clbanning", "mxj", "v2"},
				SymbolParts:  []string{"elemList", "Less"},
			},
		},
		{
			rawSymbolName: "github.com/cpuguy83/go-md2man/v2/md2man.(*roffRenderer).RenderHeader",
			expSymbolName: SymbolName{
				PackageParts: []string{"github.com", "cpuguy83", "go-md2man", "v2", "md2man"},
				SymbolParts:  []string{"(*roffRenderer)", "RenderHeader"},
			},
		},
		{
			rawSymbolName: "github.com/cli/safeexec..inittask",
			expSymbolName: SymbolName{
				PackageParts: []string{"github.com", "cli", "safeexec"},
				SymbolParts:  []string{"", "inittask"},
			},
		},
		{
			rawSymbolName: "github.com/disintegration/gift.(*unsharpMaskFilter).Draw.func1",
			expSymbolName: SymbolName{
				PackageParts: []string{"github.com", "disintegration", "gift"},
				SymbolParts:  []string{"(*unsharpMaskFilter)", "Draw", "func1"},
			},
		},
		{
			rawSymbolName: "github.com/dlclark/regexp2.(*Regexp).Replace",
			expSymbolName: SymbolName{
				PackageParts: []string{"github.com", "dlclark", "regexp2"},
				SymbolParts:  []string{"(*Regexp)", "Replace"},
			},
		},
		{
			rawSymbolName: "github.com/dlclark/regexp2/syntax._category",
			expSymbolName: SymbolName{
				PackageParts: []string{"github.com", "dlclark", "regexp2", "syntax"},
				SymbolParts:  []string{"_category"},
			},
		},
		{
			rawSymbolName: "github.com/evanw/esbuild/internal/bundler.(*Bundle).computeDataForSourceMapsInParallel",
			expSymbolName: SymbolName{
				PackageParts: []string{"github.com", "evanw", "esbuild", "internal", "bundler"},
				SymbolParts:  []string{"(*Bundle)", "computeDataForSourceMapsInParallel"},
			},
		},
		{
			rawSymbolName: "github.com/evanw/esbuild/internal/bundler.(*Bundle).computeDataForSourceMapsInParallel.func1",
			expSymbolName: SymbolName{
				PackageParts: []string{"github.com", "evanw", "esbuild", "internal", "bundler"},
				SymbolParts:  []string{"(*Bundle)", "computeDataForSourceMapsInParallel", "func1"},
			},
		},
		{
			rawSymbolName: "bytes.Join",
			expSymbolName: SymbolName{
				PackageParts: []string{"bytes"},
				SymbolParts:  []string{"Join"},
			},
		},
		{
			rawSymbolName: "bytes.(*Reader).WriteTo",
			expSymbolName: SymbolName{
				PackageParts: []string{"bytes"},
				SymbolParts:  []string{"(*Reader)", "WriteTo"},
			},
		},
		{
			rawSymbolName: "bytes.makeCutsetFunc",
			expSymbolName: SymbolName{
				PackageParts: []string{"bytes"},
				SymbolParts:  []string{"makeCutsetFunc"},
			},
		},
		{
			rawSymbolName: "bytes.makeCutsetFunc.func1",
			expSymbolName: SymbolName{
				PackageParts: []string{"bytes"},
				SymbolParts:  []string{"makeCutsetFunc", "func1"},
			},
		},
		{
			rawSymbolName: "cloud.google.com/go/iam..inittask",
			expSymbolName: SymbolName{
				PackageParts: []string{"cloud.google.com", "go", "iam"},
				SymbolParts:  []string{"", "inittask"},
			},
		},
		{
			rawSymbolName: "cloud.google.com/go/iam.glob..func1",
			expSymbolName: SymbolName{
				PackageParts: []string{"cloud.google.com", "go", "iam"},
				SymbolParts:  []string{"glob", "", "func1"},
			},
		},
		{
			rawSymbolName: "context.init",
			expSymbolName: SymbolName{
				PackageParts: []string{"context"},
				SymbolParts:  []string{"init"},
			},
		},
		{
			rawSymbolName: "context.init.0",
			expSymbolName: SymbolName{
				PackageParts: []string{"context"},
				SymbolParts:  []string{"init", "0"},
			},
		},
		{
			rawSymbolName: "encoding/xml..gobytes.1",
			expSymbolName: SymbolName{
				PackageParts: []string{"encoding", "xml"},
				SymbolParts:  []string{"", "gobytes", "1"},
			},
		},
		{
			rawSymbolName: "encoding/xml..gobytes.2",
			expSymbolName: SymbolName{
				PackageParts: []string{"encoding", "xml"},
				SymbolParts:  []string{"", "gobytes", "2"},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.rawSymbolName, func(t *testing.T) {
			s := parseSymbolName(tc.rawSymbolName)
			if !EqSymbolName(tc.expSymbolName, s) {
				t.Errorf("got %#v != exp %#v", s, tc.expSymbolName)
			}
		})
	}
}
