package symtab

import "testing"

func Test_ParseSymbolName(t *testing.T) {
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
		{
			rawSymbolName: "go.itab.vendor/golang.org/x/net/http/httpproxy.allMatch,vendor/golang.org/x/net/http/httpproxy.matcher",
			expSymbolName: SymbolName{
				PackageParts: []string{"go.itab.vendor", "golang.org", "x", "net", "http", "httpproxy.allMatch", "vendor", "golang.org", "x", "net", "http", "httpproxy"},
				SymbolParts:  []string{"matcher"},
			},
		},
		{
			rawSymbolName: "go.itab.net/http.persistConnWriter,io.Writer",
			expSymbolName: SymbolName{
				PackageParts: []string{"go.itab.net", "http.persistConnWriter", "io"},
				SymbolParts:  []string{"Writer"},
			},
		},
		{
			rawSymbolName: `go.itab.*cloud.google.com/go/iam/credentials/apiv1.IamCredentialsClient,interface { SignBlob(context.Context, *google.golang.org/genproto/googleapis/iam/credentials/v1.SignBlobRequest, ...github.com/googleapis/gax-go/v2.CallOption) (*google.golang.org/genproto/googleapis/iam/credentials/v1.SignBlobResponse, error) }`,
			expSymbolName: SymbolName{
				PackageParts: []string{"go.itab.*cloud.google.com", "go", "iam", "credentials", "apiv1"},
				SymbolParts:  []string{"IamCredentialsClient"},
			},
		},
		{
			rawSymbolName: `type..eq.struct { AccessToken string "json:\"access_token\""; ExpiresInSec int "json:\"expires_in\""; TokenType string "json:\"token_type\"" }`,
			expSymbolName: SymbolName{
				PackageParts: []string{"type"},
				SymbolParts:  []string{"", "eq", "struct"},
			},
		},
		{
			rawSymbolName: `type..eq.struct { AccessToken string "json:\"access_token\""; TokenType string "json:\"token_type\""; IDToken string "json:\"id_token\""; ExpiresIn int64 "json:\"expires_in\"" }`,
			expSymbolName: SymbolName{
				PackageParts: []string{"type"},
				SymbolParts:  []string{"", "eq", "struct"},
			},
		},
		{
			rawSymbolName: `type..eq.struct { ClientEmail string "json:\"client_email\""; PrivateKey string "json:\"private_key\"" }`,
			expSymbolName: SymbolName{
				PackageParts: []string{"type"},
				SymbolParts:  []string{"", "eq", "struct"},
			},
		},
		{
			rawSymbolName: `type..eq.struct { MediaType string "json:\"mediaType\""; github.com/gohugoio/hugo/output.Alias.2 }`,
			expSymbolName: SymbolName{
				PackageParts: []string{"type"},
				SymbolParts:  []string{"", "eq", "struct"},
			},
		},
		{
			rawSymbolName: `type..eq.struct { Name string "json:\"name\""; PrivateKeyData string "json:\"privateKeyData\"" }`,
			expSymbolName: SymbolName{
				PackageParts: []string{"type"},
				SymbolParts:  []string{"", "eq", "struct"},
			},
		},
		{
			rawSymbolName: `type..eq.struct { Name string "json:\"name\""; Type github.com/aws/aws-sdk-go/private/protocol/eventstream.valueType "json:\"type\""; Value interface {} "json:\"value\"" }`,
			expSymbolName: SymbolName{
				PackageParts: []string{"type"},
				SymbolParts:  []string{"", "eq", "struct"},
			},
		},
		{
			rawSymbolName: `type..eq.struct { Text string "json:\"text\""; Start struct { Offset int "json:\"offset\""; Column int "json:\"column\"" } "json:\"start\""; End struct { Offset int "json:\"offset\""; Column int "json:\"column\"" } "json:\"end\""; Url string "json:\"url\""; Context string "json:\"context\"" }`,
			expSymbolName: SymbolName{
				PackageParts: []string{"type"},
				SymbolParts:  []string{"", "eq", "struct"},
			},
		},
		{
			rawSymbolName: `type..eq.struct { github.com/aws/aws-sdk-go/aws/session.filename string; github.com/aws/aws-sdk-go/aws/session.field *io.Reader; github.com/aws/aws-sdk-go/aws/session.errCode string }`,
			expSymbolName: SymbolName{
				PackageParts: []string{"type"},
				SymbolParts:  []string{"", "eq", "struct"},
			},
		},
		{
			rawSymbolName: `type..eq.struct { github.com/gohugoio/hugo/resources/page.PaginatorProvider; github.com/gohugoio/hugo/resources/resource.ResourceLinksProvider; github.com/gohugoio/hugo/hugolib.targetPather }`,
			expSymbolName: SymbolName{
				PackageParts: []string{"type"},
				SymbolParts:  []string{"", "eq", "struct"},
			},
		},
		{
			rawSymbolName: `type..eq.struct { github.com/gohugoio/hugo/resources/resource.ResourceLinksProvider; github.com/gohugoio/hugo/resources/page.ContentProvider; github.com/gohugoio/hugo/resources/page.PageRenderProvider; github.com/gohugoio/hugo/resources/page.PaginatorProvider; github.com/gohugoio/hugo/resources/page.TableOfContentsProvider; github.com/gohugoio/hugo/resources/page.AlternativeOutputFormatsProvider; github.com/gohugoio/hugo/hugolib.targetPather }`,
			expSymbolName: SymbolName{
				PackageParts: []string{"type"},
				SymbolParts:  []string{"", "eq", "struct"},
			},
		},
		{
			rawSymbolName: `type..eq.struct { github.com/gohugoio/hugo/source.FileWithoutOverlap; github.com/gohugoio/hugo/resources/page.DeprecatedWarningPageMethods1 }`,
			expSymbolName: SymbolName{
				PackageParts: []string{"type"},
				SymbolParts:  []string{"", "eq", "struct"},
			},
		},
		{
			rawSymbolName: `type..eq.struct { github.com/spf13/cobra.name string; github.com/spf13/cobra.called bool }`,
			expSymbolName: SymbolName{
				PackageParts: []string{"type"},
				SymbolParts:  []string{"", "eq", "struct"},
			},
		},
		{
			rawSymbolName: `type..eq.struct { golang.org/x/image/font/sfnt.glyphIndex golang.org/x/image/font/sfnt.GlyphIndex; golang.org/x/image/font/sfnt.dx int16; golang.org/x/image/font/sfnt.dy int16; golang.org/x/image/font/sfnt.hasTransform bool; golang.org/x/image/font/sfnt.transformXX int16; golang.org/x/image/font/sfnt.transformXY int16; golang.org/x/image/font/sfnt.transformYX int16; golang.org/x/image/font/sfnt.transformYY int16 }`,
			expSymbolName: SymbolName{
				PackageParts: []string{"type"},
				SymbolParts:  []string{"", "eq", "struct"},
			},
		},
		{
			rawSymbolName: `type..eq.struct { google.golang.org/protobuf/internal/pragma.NoUnkeyedLiterals; Message google.golang.org/protobuf/reflect/protoreflect.Message }`,
			expSymbolName: SymbolName{
				PackageParts: []string{"type"},
				SymbolParts:  []string{"", "eq", "struct"},
			},
		},
		{
			rawSymbolName: `type..eq.struct { google.golang.org/protobuf/internal/pragma.NoUnkeyedLiterals; Message google.golang.org/protobuf/reflect/protoreflect.Message; Flags uint8 }`,
			expSymbolName: SymbolName{
				PackageParts: []string{"type"},
				SymbolParts:  []string{"", "eq", "struct"},
			},
		},
		{
			rawSymbolName: `type..eq.struct { google.golang.org/protobuf/internal/pragma.NoUnkeyedLiterals; Source google.golang.org/protobuf/reflect/protoreflect.Message; Destination google.golang.org/protobuf/reflect/protoreflect.Message }`,
			expSymbolName: SymbolName{
				PackageParts: []string{"type"},
				SymbolParts:  []string{"", "eq", "struct"},
			},
		},
		{
			rawSymbolName: `type..eq.struct { gopkg.in/yaml%2ev2.references int; gopkg.in/yaml%2ev2.anchor int; gopkg.in/yaml%2ev2.serialized bool }`,
			expSymbolName: SymbolName{
				PackageParts: []string{"type"},
				SymbolParts:  []string{"", "eq", "struct"},
			},
		},
		{
			rawSymbolName: `type..eq.struct { io.ReadCloser; io.Writer }`,
			expSymbolName: SymbolName{
				PackageParts: []string{"type"},
				SymbolParts:  []string{"", "eq", "struct"},
			},
		},
		{
			rawSymbolName: `type..eq.struct { io.Writer; io.Closer }`,
			expSymbolName: SymbolName{
				PackageParts: []string{"type"},
				SymbolParts:  []string{"", "eq", "struct"},
			},
		},
		{
			rawSymbolName: `type..eq.struct { runtime.gList; runtime.n int32 }`,
			expSymbolName: SymbolName{
				PackageParts: []string{"type"},
				SymbolParts:  []string{"", "eq", "struct"},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.rawSymbolName, func(t *testing.T) {
			s := ParseSymbolName(tc.rawSymbolName)
			if !EqSymbolName(tc.expSymbolName, s) {
				t.Errorf("got %#v != exp %#v", s, tc.expSymbolName)
			}
		})
	}
}
