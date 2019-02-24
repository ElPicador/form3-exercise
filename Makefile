.PHONY: run
run: dep
	go run cmd/api/main.go

.PHONY: doc
doc:
	pandoc -t beamer --pdf-engine=/Library/TeX/texbin/pdflatex -o doc/design.pdf doc/design.md

.PHONY: dep
dep:
	dep ensure

.PHONY: test
test: dep
	go test ./...
