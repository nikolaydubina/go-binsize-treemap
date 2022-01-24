docs: 
	-rm docs/*.svg
	cat testdata/hugo.symtab | ./go-binsize-treemap  > docs/hugo.svg
	cat testdata/go-graphviz.symtab | ./go-binsize-treemap  > docs/go-graphviz.svg
	cat testdata/cockroach-v21.1-geb1aa69bc4.symtab | ./go-binsize-treemap  > docs/cockroach-v21.1-geb1aa69bc4.symtab.svg
	cat testdata/hugo.symtab | ./go-binsize-treemap  -w 1080 -h 360 > docs/hugo-1080x360.svg
	cat testdata/hugo.symtab | ./go-binsize-treemap  -w 1080 -h 180 > docs/hugo-1080x180.svg
	cat testdata/hugo.symtab | ./go-binsize-treemap  -w 1080 -h 1080 -symbols > docs/hugo-1080x1080-symbols.svg
	cat testdata/hugo.symtab | ./go-binsize-treemap  -w 4096 -h 4096 -symbols > docs/hugo-4096x4096-symbols.svg

.PHONY: docs

