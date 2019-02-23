.PHONY: run
run: dep
	go run main.go

.PHONY: doc
doc:
	pandoc -t beamer -o doc/design.pdf doc/design.md

.PHONY: dep
dep:
	dep ensure

.PHONY: dep
test:
	go test ./...
