docs: 
	-rm docs/*.svg
	cat testdata/hugo.symtab | ./go-binsize-treemap  > docs/hugo.svg
	cat testdata/go-graphviz.symtab | ./go-binsize-treemap  > docs/go-graphviz.svg
	cat testdata/cockroach-v21.1-geb1aa69bc4.symtab | ./go-binsize-treemap  > docs/cockroach-v21.1-geb1aa69bc4.symtab.svg
	cat testdata/hugo.symtab | ./go-binsize-treemap  -w 1080 -h 360 > docs/hugo-1080x360.svg
	cat testdata/hugo.symtab | ./go-binsize-treemap  -w 1080 -h 180 > docs/hugo-1080x180.svg

.PHONY: docs

